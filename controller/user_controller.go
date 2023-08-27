package controller

import (
	"FP-Sanbercode-Go-48-Tubagus_Saifulloh/models"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// SignUp godoc
// @Summary create a new account.
// @Description Creating a new account to authenticate a user.
// @Tags User
// @Param Body body models.SignUpRequest true "the body for sign up"
// @Produce json
// @Success 200 {object} models.SignUpResponse
// @Router /user/signup [post]
func SignUp(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.SignUpRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if user.ID != 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user = models.User{
			Name:      req.Name,
			Email:     req.Email,
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		db.Create(&user)

		token, err := GenerateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := models.SignUpResponse{Token: token}

		c.JSON(201, gin.H{"data": res})
	}
}

// SignIn godoc
// @Summary log in with registered email.
// @Description Gaining access to the protected API. Put the token from response in the Authorization header to access protected API.
// @Tags User
// @Param Body body models.SignInRequest true "the body to sign in"
// @Produce json
// @Success 200 {object} models.SignInResponse
// @Router /user/signin [post]
func SignIn(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.SignInRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		err := db.Where("email = ?", req.Email).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if user.ID == 0 || errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password is not correct"})
			return
		}

		token, err := GenerateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := models.SignInResponse{Token: token}

		c.JSON(201, gin.H{"data": res})
	}
}

// FindUserById godoc
// @Summary get a user by id.
// @Description Finding user by id.
// @Tags User
// @Produce json
// @Param userId path int true "Project id"
// @Success 200 {object} models.UserResponse
// @Router /user/{userId} [get]
func FindUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User

		if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		res := models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Photo:     user.Photo,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		c.JSON(200, gin.H{"data": res})
	}
}

// FindUsers godoc
// @Summary get all users.
// @Description list of all users.
// @Tags User
// @Produce json
// @Success 200 {object} []models.UserResponse
// @Router /users [get]
func FindUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users := make([]models.User, 0)

		db.Find(&users)

		res := make([]models.UserResponse, 0, len(users))
		for _, user := range users {
			res = append(res, models.UserResponse{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Photo:     user.Photo,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			})
		}

		c.JSON(200, res)
	}
}

// UpdateUser godoc
// @Summary Update a user.
// @Description Update authenticated user.
// @Tags User
// @Param Body body models.UpdateUserRequest true "the body for updating user"
// @Param Authorization header string true "Bearer token"
// @Produce json
// @Success 200 {object} models.SignUpResponse
// @Router /user [put]
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authUserId, err := ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var req models.UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		if err := db.Where("id = ?", authUserId).First(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user.Name = req.Name
		user.Photo = req.Photo
		user.UpdatedAt = time.Now()

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user.Password = string(hashedPassword)

		db.Save(&user)

		res := models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Photo:     user.Photo,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		c.JSON(200, gin.H{"data": res})
	}
}

// DeleteUser godoc
// @Summary Delete a user.
// @Description Delete authenticated user.
// @Tags User
// @Param Authorization header string true "Bearer token"
// @Produce json
// @Router /user [delete]
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authUserId, err := ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		if err := db.Where("id = ?", authUserId).First(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", authUserId).Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
