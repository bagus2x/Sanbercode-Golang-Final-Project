package models

import (
	"time"
)

type (
	User struct {
		ID        int       `json:"id" gorm:"primaryKey;"`
		Name      string    `json:"name" gorm:"not null;type:varchar(256)"`
		Email     string    `json:"email" gorm:"not null;unique;type:varchar(512)"`
		Password  string    `json:"password" gorm:"not null;type:varchar(512)"`
		Photo     string    `json:"photo" gorm:"null;type:varchar(512)"`
		CreatedAt time.Time `json:"createdAt" gorm:"not null"`
		UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`

		Posts    []Post    `json:"-" gorm:"foreignKey:author_id;association_foreignkey:id;constraint:OnDelete:CASCADE"`
		Comments []Comment `json:"-" gorm:"foreignKey:author_id;association_foreignkey:id;constraint:OnDelete:CASCADE"`
	}

	SignUpRequest struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"gte=5,lte=64"`
	}

	SignUpResponse struct {
		Token string `json:"token"`
	}

	SignInRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"gte=5,lte=64"`
	}

	SignInResponse struct {
		Token string `json:"token"`
	}

	UpdateUserRequest struct {
		Name     string `json:"name" binding:"required"`
		Photo    string `json:"photo" binding:"url"`
		Password string `json:"password" binding:"required"`
	}

	UserResponse struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Photo     string    `json:"photo"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)

func (User) TableName() string {
	return "user"
}
