package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"framework/common"
	"framework/data"
)

// CreatePost insert a new Post document
// Handler for HTTP Post - "/posts
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var dataResource PostResource

	// Decode the incoming Post json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Post data",
			500,
		)
		return
	}
	post := &dataResource.Data

	val := common.GetContextAuth(r)
	if val != "" {
		post.CreatedBy = val
	}else{
		post.CreatedBy = "noone"
	}
	repo := &data.PostRepository{C: "posts"}
	// Insert a post document
	repo.Create(post)
	j, err := json.Marshal(PostResource{Data: *post})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

// GetPosts returns all Post document
// Handler for HTTP Get - "/posts"
func GetPosts(w http.ResponseWriter, r *http.Request) {

	repo := &data.PostRepository{C: "posts"}
	posts := repo.GetAll()
	j, err := json.Marshal(PostsResource{Data: posts})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// GetPostByID returns a single Post document by id
// Handler for HTTP Get - "/posts/{id}"
func GetPostByID(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	repo := &data.PostRepository{C: "posts"}
	post, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)

		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)

		}
		return
	}
	j, err := json.Marshal(post)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// GetPostsByUser returns all Posts created by a User
// Handler for HTTP Get - "/posts/users/{id}"
func GetPostsByUser(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	user := vars["id"]

	repo := &data.PostRepository{C: "posts"}
	posts := repo.GetByUser(user)
	j, err := json.Marshal(PostsResource{Data: posts})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// UpdatePost update an existing Post document
// Handler for HTTP Put - "/posts/{id}"
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource PostResource
	// Decode the incoming Post json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Post data",
			500,
		)
		return
	}
	post := &dataResource.Data
	post.Id = id

	repo := &data.PostRepository{C: "posts"}
	// Update an existing Post document
	if err := repo.Update(post); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// DeletePost deelete an existing Post document
// Handler for HTTP Delete - "/posts/{id}"
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	repo := &data.PostRepository{C: "posts"}
	// Delete an existing Post document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
