package models

import (
	"time"
)

type (
	Post struct {
		ID        int       `json:"id" gorm:"primaryKey"`
		AuthorID  int       `json:"authorId" gorm:"not null"`
		Title     string    `json:"title" gorm:"not null;type:varchar(256)"`
		Body      string    `json:"body" gorm:"not null"`
		Thumbnail string    `json:"thumbnail" gorm:"not null;type:varchar(512)"`
		CreatedAt time.Time `json:"createdAt" gorm:"not null"`
		UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`

		Categories []Category `json:"categories" gorm:"many2many:post_categories;constraint:OnDelete:CASCADE"`
	}

	CreatePostRequest struct {
		Title       string `json:"title" binding:"required,lte=256"`
		Body        string `json:"body" binding:"required,lte=8142"`
		Thumbnail   string `json:"thumbnail" binding:"url"`
		CategoryIDs []int  `json:"categoryIds"`
	}

	UpdatePostRequest struct {
		Title       string `json:"title" binding:"required,lte=256"`
		Body        string `json:"body" binding:"required,lte=8142"`
		Thumbnail   string `json:"thumbnail" binding:"url"`
		CategoryIds []int  `json:"categoryIds"`
	}

	PostResponse struct {
		ID         int                `json:"id"`
		Title      string             `json:"title" binding:"required,lte=256"`
		Body       string             `json:"body" binding:"required,lte=8142"`
		Categories []CategoryResponse `json:"categories"`
		Thumbnail  string             `json:"thumbnail" binding:"url"`
		CreatedAt  time.Time          `json:"createdAt" gorm:"not null"`
		UpdatedAt  time.Time          `json:"updatedAt" gorm:"not null"`
	}
)

func (Post) TableName() string {
	return "post"
}
