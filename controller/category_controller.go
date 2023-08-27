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

// AddCategory godoc
// @Summary create a new category.
// @Description Creating a new category.
// @Tags Category
// @Param Body body models.AddCategoryRequest true "the body for creating category"
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 201 {object} models.CategoryResponse
// @Router /category [post]
func AddCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.AddCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		category := models.Category{
			Name:      req.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&category).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := models.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}

		c.JSON(201, res)
	}
}

// FindCategoryById godoc
// @Summary get a category by id.
// @Description finding a category by id.
// @Tags Category
// @Produce json
// @Param categoryId path int true "Category id"
// @Success 200 {object} models.CategoryResponse
// @Router /category/{categoryId} [get]
func FindCategoryById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var category models.Category

		if err := db.Where("id = ?", categoryId).First(&category).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := models.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}

		c.JSON(200, res)
	}
}

// FindCategories godoc
// @Summary get all categories.
// @Description list of all categories.
// @Tags Category
// @Produce json
// @Success 200 {object} []models.CategoryResponse
// @Router /categories [get]
func FindCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories := make([]models.Category, 0)

		if err := db.Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := make([]models.CategoryResponse, 0, len(categories))
		for _, category := range categories {
			res = append(res, models.CategoryResponse{
				ID:        category.ID,
				Name:      category.Name,
				CreatedAt: category.CreatedAt,
				UpdatedAt: category.UpdatedAt,
			})
		}

		c.JSON(200, res)
	}
}

// UpdateCategory godoc
// @Summary update single category.
// @Description updating a category.
// @Tags Category
// @Param Body body models.UpdateCategoryRequest true "the body for updating category"
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param categoryId path int true "Category id"
// @Success 200 {object} models.CategoryResponse
// @Router /category/{categoryId} [put]
func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req models.UpdateCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var category models.Category

		if err := db.Where("id = ?", categoryId).First(&category).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		category.Name = req.Name
		category.UpdatedAt = time.Now()

		db.Save(&category)

		res := models.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		c.JSON(http.StatusOK, res)
	}
}

// DeleteCategory godoc
// @Summary delete a category.
// @Description deleting a new category.
// @Tags Category
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param categoryId path int true "Category id"
// @Router /category/{categoryId} [delete]
func DeleteCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var category models.Category
		if err := db.Where("id = ?", categoryId).First(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", categoryId).Delete(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
