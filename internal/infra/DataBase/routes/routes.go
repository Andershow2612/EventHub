package routes

import (
	"eventHub.com/internal/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB){
	UserController := controller.NewUserController(db)

	r.GET("/users", UserController.Getusers)
}