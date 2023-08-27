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

// AddComment godoc
// @Summary create a new comment.
// @Description Creating a new comment.
// @Tags Comment
// @Param Body body models.AddCommentRequest true "the body for creating comment"
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 201 {object} models.Comment
// @Router /comment [post]
func AddComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.AddCommentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		authUserId, err := ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		comment := models.Comment{
			PostID:    req.PostID,
			AuthorId:  authUserId,
			Body:      req.Body,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, comment)
	}
}

// FindCommentById godoc
// @Summary get a comment by id.
// @Description finding a comment by id.
// @Tags Comment
// @Produce json
// @Param commentId path int true "Comment id"
// @Success 200 {object} models.Comment
// @Router /comment/{commentId} [get]
func FindCommentById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var comment models.Comment

		if err := db.Where("id = ?", commentId).First(&comment).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, comment)
	}
}

// FindCommentsByPostId godoc
// @Summary get comments by post id.
// @Description finding  comment by post id.
// @Tags Comment
// @Produce json
// @Param postId path int true "Post id"
// @Success 200 {object} []models.Comment
// @Router /post/{postId}/comments [get]
func FindCommentsByPostId(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		comments := make([]models.Comment, 0)
		db.Where("post_id", postID).Find(&comments)

		c.JSON(200, comments)
	}
}

// UpdateComment godoc
// @Summary update single comment.
// @Description updating a comment.
// @Tags Comment
// @Param Body body models.UpdateCommentRequest true "the body for updating comment"
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param commentId path int true "Comment id"
// @Success 200 {object} models.Comment
// @Router /comment/{commentId} [put]
func UpdateComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req models.UpdateCommentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var comment models.Comment

		if err := db.Where("id = ?", commentId).First(&comment).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		comment.Body = req.Body
		comment.UpdatedAt = time.Now()

		db.Save(&comment)

		c.JSON(http.StatusOK, comment)
	}
}

// DeleteComment godoc
// @Summary delete single comment.
// @Description deleting a comment by id.
// @Tags Comment
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param commentId path int true "Comment id"
// @Success 204
// @Router /comment/{commentId} [delete]
func DeleteComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var comment models.Comment
		if err := db.Where("id = ?", commentId).First(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", commentId).Delete(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
