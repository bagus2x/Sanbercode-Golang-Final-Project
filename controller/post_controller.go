package controller

import (
	"FP-Sanbercode-Go-48-Tubagus_Saifulloh/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// AddPost godoc
// @Summary create a new post.
// @Description Creating a new post.
// @Tags Post
// @Param Body body models.CreatePostRequest true "the body for creating post"
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 201 {object} models.PostResponse
// @Router /post [post]
func AddPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreatePostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		authUserId, err := ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		post := models.Post{
			AuthorID:   authUserId,
			Title:      req.Title,
			Body:       req.Body,
			Thumbnail:  req.Thumbnail,
			Categories: make([]models.Category, 0),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		for _, categoryID := range req.CategoryIDs {
			post.Categories = append(post.Categories, models.Category{ID: categoryID})
		}

		if err := db.Create(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ? ", post.ID).Preload("Categories").First(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := models.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Body:      post.Body,
			Thumbnail: post.Thumbnail,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		for _, category := range post.Categories {
			res.Categories = append(res.Categories, models.CategoryResponse{
				ID:        category.ID,
				Name:      category.Name,
				CreatedAt: category.CreatedAt,
				UpdatedAt: category.UpdatedAt,
			})
		}

		c.JSON(201, res)
	}
}

// FindPostById godoc
// @Summary get a category by id.
// @Description finding a post by id.
// @Tags Post
// @Produce json
// @Param postId path int true "Post id"
// @Success 200 {object} models.PostResponse
// @Router /post/{postId} [get]
func FindPostById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var post models.Post

		if err := db.Where("id = ?", categoryId).Preload("Categories").First(&post).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := models.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Body:      post.Body,
			Thumbnail: post.Thumbnail,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		for _, category := range post.Categories {
			res.Categories = append(res.Categories, models.CategoryResponse{
				ID:        category.ID,
				Name:      category.Name,
				CreatedAt: category.CreatedAt,
				UpdatedAt: category.UpdatedAt,
			})
		}

		c.JSON(200, res)
	}
}

// FindPosts godoc
// @Summary get all posts.
// @Description list of all posts.
// @Tags Post
// @Produce json
// @Success 200 {object} []models.PostResponse
// @Router /posts [get]
func FindPosts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts := make([]models.Post, 0)

		if err := db.Preload("Categories").Find(&posts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := make([]models.PostResponse, 0, len(posts))
		for _, post := range posts {
			postRes := models.PostResponse{
				ID:        post.ID,
				Title:     post.Title,
				Body:      post.Body,
				Thumbnail: post.Thumbnail,
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
			}
			for _, category := range post.Categories {
				postRes.Categories = append(postRes.Categories, models.CategoryResponse{
					ID:        category.ID,
					Name:      category.Name,
					CreatedAt: category.CreatedAt,
					UpdatedAt: category.UpdatedAt,
				})
			}
			res = append(res, postRes)
		}

		c.JSON(200, res)
	}
}

// UpdatePost godoc
// @Summary update single post.
// @Description updating a post.
// @Tags Post
// @Param Body body models.UpdatePostRequest true "the body for updating category"
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param postId path int true "Post id"
// @Success 200 {object} models.PostResponse
// @Router /post/{postId} [put]
func UpdatePost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req models.UpdatePostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var post models.Post

		if err := db.Where("id = ?", postID).First(&post).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		post.Title = req.Title
		post.Body = req.Body
		post.Thumbnail = req.Thumbnail
		post.UpdatedAt = time.Now()

		post.Categories = make([]models.Category, 0, len(req.CategoryIds))
		for _, categoryID := range req.CategoryIds {
			post.Categories = append(post.Categories, models.Category{ID: categoryID})
		}

		if err := db.Model(&post).Association("Categories").Replace(post.Categories); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Save(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", postID).Preload("Categories").First(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := models.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Body:      post.Body,
			Thumbnail: post.Thumbnail,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		for _, category := range post.Categories {
			res.Categories = append(res.Categories, models.CategoryResponse{
				ID:        category.ID,
				Name:      category.Name,
				CreatedAt: category.CreatedAt,
				UpdatedAt: category.UpdatedAt,
			})
		}

		c.JSON(http.StatusOK, res)
	}
}

// DeletePost godoc
// @Summary delete single post.
// @Description deleting a post.
// @Tags Post
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param postId path int true "Post id"
// @Success 204
// @Router /post/{postId} [delete]
func DeletePost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var post models.Post
		if err := db.Where("id = ?", postID).First(&post).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "post not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", postID).Delete(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
