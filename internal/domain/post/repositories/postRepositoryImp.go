package repositories

import (
	"github.com/hudayberdipolat/golang-api-crud/internal/models"
	"gorm.io/gorm"
	"log"
)

type postRepositoryImp struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return postRepositoryImp{
		db: db,
	}
}

func (p postRepositoryImp) GetAll() ([]models.Post, error) {
	var posts []models.Post
	if err := p.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p postRepositoryImp) GetOne(postID int) (*models.Post, error) {
	var post models.Post
	if err := p.db.Where("id=?", postID).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (p postRepositoryImp) Create(post models.Post) error {
	if err := p.db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func (p postRepositoryImp) Update(postID int, post models.Post) error {
	var postModel models.Post
	if err := p.db.Model(&postModel).Where("id=?", postID).Updates(&post).Error; err != nil {
		return err
	}
	return nil
}

func (p postRepositoryImp) Delete(postID int) error {
	var post models.Post
	err := p.db.Model(&post).Where("id=?", postID).Unscoped().Delete(&post).Error
	if err != nil {
		log.Println("error ", err.Error())
		return err
	}
	return nil
}
