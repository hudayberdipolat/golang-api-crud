package constructor

import (
	"github.com/hudayberdipolat/golang-api-crud/internal/domain/post/handler"
	"github.com/hudayberdipolat/golang-api-crud/internal/domain/post/repositories"
	"github.com/hudayberdipolat/golang-api-crud/internal/domain/post/services"
	"github.com/hudayberdipolat/golang-api-crud/pkg/config"
	"gorm.io/gorm"
)

var postRepository repositories.PostRepository
var posService services.PostService
var PostHandler handler.PostHandler

func PostRequirementsCreator(db *gorm.DB, conf config.Config) {
	postRepository = repositories.NewPostRepository(db)
	posService = services.NewPostService(postRepository)
	PostHandler = handler.NewPostHandler(posService, conf)
}
