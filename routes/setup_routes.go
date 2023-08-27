package routes

import (
	"FP-Sanbercode-Go-48-Tubagus_Saifulloh/controller"
	"FP-Sanbercode-Go-48-Tubagus_Saifulloh/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/user/signup", controller.SignUp(db))
	r.POST("/user/signin", controller.SignIn(db))
	r.GET("/user/:id", controller.FindUserById(db))
	r.GET("/users", controller.FindUsers(db))
	r.PUT("/user", middleware.Auth(), controller.UpdateUser(db))
	r.DELETE("/user", middleware.Auth(), controller.DeleteUser(db))

	r.POST("/category", middleware.Auth(), controller.AddCategory(db))
	r.GET("/category/:id", controller.FindCategoryById(db))
	r.GET("/categories", controller.FindCategories(db))
	r.PUT("/category/:id", controller.UpdateCategory(db))
	r.DELETE("/category/:id", controller.DeleteCategory(db))

	r.POST("/post", middleware.Auth(), controller.AddPost(db))
	r.GET("/post/:id", controller.FindPostById(db))
	r.GET("/posts", controller.FindPosts(db))
	r.PUT("/post/:id", middleware.Auth(), controller.UpdatePost(db))
	r.DELETE("/post/:id", middleware.Auth(), controller.DeletePost(db))

	r.POST("/comment", middleware.Auth(), controller.AddComment(db))
	r.GET("/comment/:id", controller.FindCommentById(db))
	r.GET("/post/:id/comments", controller.FindCommentsByPostId(db))
	r.PUT("/comment/:id", middleware.Auth(), controller.UpdateComment(db))
	r.DELETE("/comment/:id", middleware.Auth(), controller.DeleteComment(db))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
