package services

import (
	"errors"
	"project-blog/models"
	"project-blog/repository"
)

type PostService interface {
	Create(post *models.Post) error
	FindAll() ([]models.Post, error)
	FindByID(id string) (*models.Post, error)
	Update(post *models.Post) error
	Delete(id string, userID string) error
}

type postService struct {
	Repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{Repo: repo}
}

func (s *postService) Create(post *models.Post) error {
	return s.Repo.Create(post)
}

func (s *postService) FindAll() ([]models.Post, error) {
	return s.Repo.FindAll()
}

func (s *postService) FindByID(id string) (*models.Post, error) {
	return s.Repo.FindByID(id)
}

func (s *postService) Update(post *models.Post) error {
	return s.Repo.Update(post)
}

func (s *postService) Delete(id string, userID string) error {
	post, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}
	if post.AuthorID.String() != userID {
		return errors.New("unauthorized")
	}
	return s.Repo.Delete(id)
}
