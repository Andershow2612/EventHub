package controller

import (
	"net/http"

	"eventHub.com/internal/repository"
	"eventHub.com/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController(db *gorm.DB) *UserController{
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(*repo)
	return &UserController{
		Service: service,
	}
}


func(c *UserController) Getusers(ctx *gin.Context){

	users, err := c.Service.ListUser()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, users)
}