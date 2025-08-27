package routes

import (
	"eventHub.com/internal/controller"
	"eventHub.com/internal/middleware"
	"eventHub.com/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB){
	UserController := controller.NewUserController(db)

	r.GET("/users", middleware.AuthMiddleware(), middleware.RoleMiddleware(utils.AdmRole), UserController.Getusers)
	r.GET("/user/:id", UserController.UserById)
	r.POST("/create", UserController.CreateUser)
	r.POST("/login", UserController.Login)
	r.PUT("/updateUser/:id", UserController.UpdateUser)
	r.PUT("/updatePassord/:id", UserController.UpdatePassword)
}