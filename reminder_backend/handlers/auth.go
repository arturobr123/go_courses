package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reminder_backend/models"
	"reminder_backend/services"

	"firebase.google.com/go/auth"
	"golang.org/x/crypto/bcrypt"
)

type SignUpLogInResponse struct {
	Token string `json:"token"`
	// Id       string `json:"id"`
	// Email    string `json:"email"`
	// Password string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := (&auth.UserToCreate{}).
		Email(creds.Email).
		Password(string(hash)).
		DisplayName("Test user")

	_, err = services.FirebaseAuth.CreateUser(context.Background(), params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := authenticateUser(services.FirebaseAuth, creds.Email, creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SignUpLogInResponse{
		Token: token,
	})
}

func authenticateUser(client *auth.Client, email, password string) (string, error) {
	// Sign in with email and password using Firebase REST API
	type SignInWithPasswordResponse struct {
		IDToken      string `json:"idToken"`
		Email        string `json:"email"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
		LocalID      string `json:"localId"`
		Registered   bool   `json:"registered"`
	}

	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=[API_KEY]"

	payload := map[string]string{
		"email":             email,
		"password":          password,
		"returnSecureToken": "true",
	}

	var result SignInWithPasswordResponse
	err := postRequest(url, payload, &result)
	if err != nil {
		return "", err
	}

	return result.IDToken, nil
}

func postRequest(url string, payload interface{}, result interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response status: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}
