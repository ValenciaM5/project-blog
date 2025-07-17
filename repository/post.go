package repository

import (
	"project-blog/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) error
	FindAll() ([]models.Post, error)
	FindByID(id string) (*models.Post, error)
	Update(post *models.Post) error
	Delete(id string) error
}

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{DB: db}
}

func (r *postRepository) Create(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *postRepository) FindAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.DB.Preload("Author").Find(&posts).Error
	return posts, err
}

func (r *postRepository) FindByID(id string) (*models.Post, error) {
	var post models.Post
	err := r.DB.Preload("Author").First(&post, "id = ?", id).Error
	return &post, err
}

func (r *postRepository) Update(post *models.Post) error {
	return r.DB.Save(post).Error
}

func (r *postRepository) Delete(id string) error {
	return r.DB.Delete(&models.Post{}, "id = ?", id).Error
}
