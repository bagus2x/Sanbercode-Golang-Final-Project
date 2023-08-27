package models

import (
	"time"
)

type (
	Category struct {
		ID        int       `json:"id" gorm:"primaryKey"`
		Name      string    `json:"name" gorm:"not null;type:varchar(256);unique"`
		CreatedAt time.Time `json:"createdAt" gorm:"not null"`
		UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`

		Posts []Post `json:"posts" gorm:"many2many:post_categories;constraint:OnDelete:CASCADE"`
	}

	AddCategoryRequest struct {
		Name string `json:"name" binding:"required,lte=256"`
	}

	CategoryResponse struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	UpdateCategoryRequest struct {
		Name string `json:"name" binding:"required,lte=256"`
	}
)

func (Category) TableName() string {
	return "category"
}
