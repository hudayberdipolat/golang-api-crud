package dto

import (
	"github.com/hudayberdipolat/golang-api-crud/internal/models"
)

type PostResponse struct {
	ID        uint   `json:"id"`
	PostTitle string `json:"post_title"`
	PostSlug  string `json:"post_slug"`
	PostDesc  string `json:"post_desc"`
	PostImage string `json:"post_image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewPostResponse(post models.Post) PostResponse {
	return PostResponse{
		ID:        post.ID,
		PostTitle: post.PostTitle,
		PostSlug:  post.PostSlug,
		PostDesc:  post.PostDesc,
		PostImage: post.PostImage,
		CreatedAt: post.CreatedAt.Format("01-02-2006"),
		UpdatedAt: post.UpdatedAt.Format("01-02-2006"),
	}
}
