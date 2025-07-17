package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"project-blog/models"
	"project-blog/services"
)

type PostHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	Service services.PostService
}

func NewPostHandler(service services.PostService) PostHandler {
	return &postHandler{Service: service}
}

func (h *postHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(string)
	post.AuthorID = uuid.MustParse(userID)
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	if err := h.Service.Create(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdPost, err := h.Service.FindByID(post.ID.String())
	if err != nil {
		http.Error(w, "Failed to fetch created post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}

func (h *postHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func (h *postHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	post, err := h.Service.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func (h *postHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(string)
	post.ID = uuid.MustParse(id)
	post.AuthorID = uuid.MustParse(userID)
	post.UpdatedAt = time.Now()

	if err := h.Service.Update(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	updatedPost, err := h.Service.FindByID(id)
	if err != nil {
		http.Error(w, "Failed to fetch updated post", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedPost)
}
func (h *postHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	
	userID := r.Context().Value("user_id").(string)

	err := h.Service.Delete(id, userID)
	if err != nil {
		
		if err.Error() == "unauthorized" {
			http.Error(w, "Unauthorized to delete this post", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Post deleted successfully",
	})
}
