package models

import "time"

type Post struct {
	ID        uint      `json:"id"`
	PostTitle string    `json:"post_title"`
	PostSlug  string    `json:"post_slug"`
	PostDesc  string    `json:"post_desc"`
	PostImage string    `json:"post_image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (*Post) TableName() string {
	return "posts"
}
