package repositories

import "github.com/hudayberdipolat/golang-api-crud/internal/models"

type PostRepository interface {
	GetAll() ([]models.Post, error)
	GetOne(postID int) (*models.Post, error)
	Create(post models.Post) error
	Update(postID int, post models.Post) error
	Delete(postID int) error
}
