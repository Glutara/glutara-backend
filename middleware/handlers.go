package middleware

import (
	"encoding/json"
	"net/http"
	"fmt"
	"math/rand"

	"glutara/repository"
	"glutara/models"
)

// Create a global instance of Repo
var (
	repo repository.PostRepository = repository.NewPostRepository() 
)

// MainHandler for dummy test
func MainHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, Welcome to Glutara Web Service!")
}

// Get all the posts from the firestore
func GetPosts(response http.ResponseWriter , request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error" : "Error Getting the Posts"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

// Add a new post to the firestore
func AddPost(response http.ResponseWriter , request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post models.Post
	err:= json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error" : "Error Unmarshalling Data"}`))
		return
	}

	post.ID = rand.Int63()
	repo.Save(&post)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}