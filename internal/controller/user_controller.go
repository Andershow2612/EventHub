package controller

import (
	"net/http"
	"strconv"

	dto "eventHub.com/internal/DTO"
	"eventHub.com/internal/mapper"
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

	ctx.JSON(http.StatusOK, mapper.ToUserResponseList(users))
}

func (c *UserController) UserById(ctx *gin.Context){

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.Service.ListAUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToUserResponse(user))
	
}

func (c *UserController) CreateUser(ctx *gin.Context){
	var req dto.UserReq

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	user := mapper.ToUserEntity(req)

	createdUser, err := c.Service.CreateUser(user)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, mapper.ToUserCreated(createdUser))

}

func (c *UserController) Login(ctx *gin.Context){
	var req dto.LoginReq

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	token, err := c.Service.Login(req.Email, req.Password)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Message": "access denied",
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "access valid",
		"Token": token,
	})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var req dto.UserUpdateReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedEntity := mapper.ToUserUpdateEntity(req)
    updatedUser, err := c.Service.UpdateUserFields(id, updatedEntity)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    res := mapper.ToUserResponse(updatedUser)
    ctx.JSON(http.StatusOK, res)
}

func (c *UserController) UpdatePassword(ctx *gin.Context){
	
	var req dto.UserUpdatePassword

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.Service.UpdatePassword(id, req.OldPassword, req.NewPassword)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user updated sucessfully",
	})
}
