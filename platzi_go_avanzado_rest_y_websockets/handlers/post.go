package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"platzi/go/rest_websockets/models"
	"platzi/go/rest_websockets/repository"
	"platzi/go/rest_websockets/server"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

type UpsertPostRequest struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}

type PostUpdateResponse struct {
	Message string `json:"message"`
}

func GetPostByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		post, err := repository.GetPostById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PostResponse{
			Id:          post.Id,
			PostContent: post.PostContent,
		})
	}
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, postRequest, err := ParseAndDecode(r, s)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		post := models.Post{
			Id:          id.String(),
			UserId:      claims.UserId,
			PostContent: postRequest.PostContent,
		}

		err = repository.InsertPost(r.Context(), &post)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PostResponse{
			Id:          post.Id,
			PostContent: post.PostContent,
		})
	}
}

func UpdatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, postRequest, err := ParseAndDecode(r, s)
		params := mux.Vars(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		post := models.Post{
			Id:          params["id"],
			UserId:      claims.UserId,
			PostContent: postRequest.PostContent,
		}

		err = repository.UpdatePost(r.Context(), &post)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PostUpdateResponse{
			Message: "Post updated correctly",
		})
	}
}

func DeletePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("DeletePostHandler")

		claims, err := ParseJWT(r, s)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		params := mux.Vars(r)

		err = repository.DeletePost(r.Context(), params["id"], claims.UserId)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PostUpdateResponse{
			Message: "Post deleted correctly",
		})
	}
}

func ListPostsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("ListPostsHandler")

		var err error
		pageStr := r.URL.Query().Get("page")

		var page uint64

		if pageStr == "" {
			page = 0
		} else {
			page, err = strconv.ParseUint(pageStr, 10, 64)
		}

		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		posts, err := repository.ListPosts(r.Context(), page)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}
