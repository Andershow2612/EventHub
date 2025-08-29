package controller

import (
	"net/http"
	"strconv"

	"eventHub.com/internal/mapper"
	"eventHub.com/internal/repository"
	"eventHub.com/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EventController struct {
	Service *service.EventService
}

func NewEventController(db *gorm.DB) *EventController{
	repo := repository.NewEventRepository(db)
	service:= service.NewEventService(*repo)
	return &EventController{
		Service: service,
	}
}

func (c *EventController) GetEvents(ctx *gin.Context){
	events, err := c.Service.ListEvents()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToEventResponseList(events))
}

func (c *EventController) GetEventID(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event, err := c.Service.ListEventID(id)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToEventResponse(event))
}