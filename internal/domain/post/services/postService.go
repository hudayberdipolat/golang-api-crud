package services

import "github.com/hudayberdipolat/golang-api-crud/internal/domain/post/dto"

type PostService interface {
	GetAllPost() ([]dto.PostResponse, error)
	GetOnePost(postID int) (*dto.PostResponse, error)
	CreatePost(createRequest *dto.CreateRequest) error
	UpdatePost(postID int, updateRequest *dto.UpdateRequest) error
	DeletePost(postID int) error
}
