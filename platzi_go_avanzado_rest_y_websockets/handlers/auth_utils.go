package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"platzi/go/rest_websockets/models"
	"platzi/go/rest_websockets/server"

	"github.com/golang-jwt/jwt"
)

// func startDebugger() {
// 	// Check if Delve is installed and the GOBIN path is set correctly
// 	if _, err := exec.LookPath("dlv"); err != nil {
// 		fmt.Println("Delve is not installed or not in PATH")
// 		return
// 	}

// 	// Start Delve attached to this process
// 	cmd := exec.Command("dlv", "attach", fmt.Sprint(os.Getpid()))
// 	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
// 	if err := cmd.Start(); err != nil {
// 		fmt.Println("Failed to start Delve:", err)
// 		return
// 	}
// 	fmt.Println("Delve started successfully")
// }

// ParseJWT parses the JWT from the request, validates it, and returns the claims if valid.
func ParseJWT(r *http.Request, s server.Server) (*models.AppClaims, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// Helper function to parse JWT and decode request
func ParseAndDecode(r *http.Request, s server.Server) (*models.AppClaims, *UpsertPostRequest, error) {
	claims, err := ParseJWT(r, s)

	if err != nil {
		return nil, nil, err
	}

	postRequest := UpsertPostRequest{}
	if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
		return nil, nil, errors.New("bad request, missing params")
	}

	return claims, &postRequest, nil
}
