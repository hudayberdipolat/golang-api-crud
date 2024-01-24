package services

import (
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/golang-api-crud/internal/domain/post/dto"
	"github.com/hudayberdipolat/golang-api-crud/internal/domain/post/repositories"
	"github.com/hudayberdipolat/golang-api-crud/internal/models"
	"time"
)

type postServiceImp struct {
	postRepo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return postServiceImp{
		postRepo: repo,
	}
}

func (p postServiceImp) GetAllPost() ([]dto.PostResponse, error) {
	posts, err := p.postRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var postResponses []dto.PostResponse
	for _, post := range posts {
		postResponse := dto.NewPostResponse(post)
		postResponses = append(postResponses, postResponse)
	}
	return postResponses, nil
}

func (p postServiceImp) GetOnePost(postID int) (*dto.PostResponse, error) {
	post, err := p.postRepo.GetOne(postID)
	if err != nil {
		return nil, err
	}
	postResponse := dto.NewPostResponse(*post)
	return &postResponse, nil
}

func (p postServiceImp) CreatePost(createRequest *dto.CreateRequest) error {
	// image upload yerine yetirmeli
	post := models.Post{
		PostTitle: createRequest.PostTitle,
		PostSlug:  slug.Make(createRequest.PostTitle),
		PostDesc:  createRequest.PostDesc,
		PostImage: *createRequest.PostImage,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := p.postRepo.Create(post); err != nil {
		return err
	}
	return nil
}

func (p postServiceImp) UpdatePost(postID int, updateRequest *dto.UpdateRequest) error {
	updatePost := models.Post{
		PostTitle: updateRequest.PostTitle,
		PostSlug:  slug.Make(updateRequest.PostTitle),
		PostDesc:  updateRequest.PostDesc,
		PostImage: *updateRequest.PostImage,
		UpdatedAt: time.Now(),
	}
	if err := p.postRepo.Update(postID, updatePost); err != nil {
		return err
	}
	return nil
}

func (p postServiceImp) DeletePost(postID int) error {
	if err := p.postRepo.Delete(postID); err != nil {
		return err
	}
	return nil
}
