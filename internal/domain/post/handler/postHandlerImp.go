package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/golang-api-crud/internal/domain/post/dto"
	"github.com/hudayberdipolat/golang-api-crud/internal/domain/post/services"
	"github.com/hudayberdipolat/golang-api-crud/internal/utils"
	"github.com/hudayberdipolat/golang-api-crud/internal/utils/response"
	"github.com/hudayberdipolat/golang-api-crud/internal/utils/validate"
	"github.com/hudayberdipolat/golang-api-crud/pkg/config"
	"net/http"
	"strconv"
)

type postHandlerImp struct {
	config      config.Config
	postService services.PostService
}

func NewPostHandler(service services.PostService, conf config.Config) PostHandler {
	return postHandlerImp{
		config:      conf,
		postService: service,
	}
}

func (p postHandlerImp) GetAll(ctx *fiber.Ctx) error {
	posts, err := p.postService.GetAllPost()
	if err != nil {
		errResponse := response.Error("get all posts error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(true, "get All post", posts)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (p postHandlerImp) GetOne(ctx *fiber.Ctx) error {
	postID, _ := strconv.Atoi(ctx.Params("postID"))
	post, err := p.postService.GetOnePost(postID)
	if err != nil {
		errResponse := response.Error("get post error", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(true, "get All post", post)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (p postHandlerImp) Create(ctx *fiber.Ctx) error {
	var createPostRequest dto.CreateRequest
	// body parser
	if err := ctx.BodyParser(&createPostRequest); err != nil {
		errResponse := response.Error("body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if err := validate.ValidateStruct(createPostRequest); err != nil {
		errResponse := response.Error("validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	//HANDLE FILE STARTED
	path, err := utils.UploadFile(ctx, "post_image", p.config.PublicPath, "postImages")
	if err != nil {
		errResponse := response.Error("Error file upload", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	//HANDLE FILE END

	createPostRequest.PostImage = path
	if err := p.postService.CreatePost(&createPostRequest); err != nil {
		errResponse := response.Error("post can't created", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(true, "post created successfully", nil)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (p postHandlerImp) Update(ctx *fiber.Ctx) error {
	var updatePostRequest dto.UpdateRequest
	postID, _ := strconv.Atoi(ctx.Params("postID"))

	// body parser
	if err := ctx.BodyParser(&updatePostRequest); err != nil {
		errResponse := response.Error("body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if err := validate.ValidateStruct(updatePostRequest); err != nil {
		errResponse := response.Error("validate error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// post update
	// eger input-da post-image bar bolsa onda post-yn onki bar bolan suratyny ocurmeli
	// we taze post-image upload etmeli
	file, _ := ctx.FormFile("post_image")
	if file != nil {
		//old_image delete
		if errOldImageDelete := utils.DeleteFile(*updatePostRequest.OldImage); errOldImageDelete != nil {
			errResponse := response.Error("Error not deleted old image", errOldImageDelete.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
		}
		//new image upload

		path, errFileUpload := utils.UploadFile(ctx, "post_image", p.config.PublicPath, "postImages")
		if errFileUpload != nil {
			errResponse := response.Error("Error file upload", errFileUpload.Error(), nil)
			return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
		}
		updatePostRequest.PostImage = path
	} else {
		updatePostRequest.PostImage = updatePostRequest.OldImage
	}

	if errUpdatePost := p.postService.UpdatePost(postID, &updatePostRequest); errUpdatePost != nil {
		errResponse := response.Error("update post error", errUpdatePost.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	successResponse := response.Success(true, "post updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (p postHandlerImp) Delete(ctx *fiber.Ctx) error {
	postID, _ := strconv.Atoi(ctx.Params("postID"))
	post, err := p.postService.GetOnePost(postID)
	if err != nil {
		errResponse := response.Error("post not deleted", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	errImageDelete := utils.DeleteFile(post.PostImage)
	if errImageDelete != nil {
		errResponse := response.Error("post not deleted", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	errPostDelete := p.postService.DeletePost(postID)
	if errPostDelete != nil {
		errResponse := response.Error("post not deleted", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(true, "post deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
