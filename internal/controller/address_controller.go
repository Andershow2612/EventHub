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

type AddressController struct {
	Service *service.AddressService
}

func NewAddressController(db *gorm.DB) *AddressController{
	repo := repository.NewAddressRepository(db)
	service := service.NewAddressService(*repo)
	return &AddressController{
		Service: service,
	}
}

func (c *AddressController) GetAddress(ctx *gin.Context){
	address, err := c.Service.ListAddrres()
	
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, mapper.ToAddressResponseList(address))
}

func (c *AddressController) GetAddresID(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	address, err := c.Service.ListAddrresID(id)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, mapper.ToAddressResponse(address))
}

func (c *AddressController) CreateAddress(ctx *gin.Context){
	var req dto.AddressRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	address := mapper.ToAddressEntity(req)

	createdAddress, err := c.Service.CreateAddress(address)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, createdAddress)
}

func (c *AddressController) DeleteAddress(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	address, err := c.Service.DeleteAddress(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": address})
}