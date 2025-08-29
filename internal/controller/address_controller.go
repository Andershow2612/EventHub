package controller

import (
	"net/http"

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