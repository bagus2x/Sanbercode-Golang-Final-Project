package models

import (
	"time"
)

type (
	Comment struct {
		ID        int       `json:"id" gorm:"primaryKey"`
		AuthorId  int       `json:"authorId" gorm:"not null"`
		PostID    int       `json:"postId" gorm:"not null"`
		Body      string    `json:"body" gorm:"not null"`
		CreatedAt time.Time `json:"createdAt" gorm:"not null"`
		UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
	}

	AddCommentRequest struct {
		PostID int    `json:"postId" binding:"required"`
		Body   string `json:"body" binding:"required"`
	}

	UpdateCommentRequest struct {
		Body string `json:"body" binding:"required"`
	}
)

func (Comment) TableName() string {
	return "comment"
}
