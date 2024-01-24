package dto

type CreateRequest struct {
	PostTitle string  `json:"post_title" form:"post_title" validate:"required,min=5"`
	PostDesc  string  `json:"post_desc" form:"post_desc"  validate:"required,min=5"`
	PostImage *string `json:"post_image,omitempty" form:"post_image,omitempty" `
}

type UpdateRequest struct {
	PostTitle string  `json:"post_title" form:"post_title" validate:"required,min=5"`
	PostDesc  string  `json:"post_desc" form:"post_desc"  validate:"required,min=5"`
	PostImage *string `json:"post_image,omitempty" form:"post_image,omitempty"`
	OldImage  *string `json:"old_image,omitempty" form:"old_image,omitempty" validate:"required"`
}
