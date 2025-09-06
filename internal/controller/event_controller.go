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

func (c *EventController) CreateEvent(ctx *gin.Context){
	var req dto.EventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	event := mapper.ToEventEntity(req)

	createdEvent, err := c.Service.CreateEvent(event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToEventResponse(createdEvent))
}

func (c *EventController) UpdateEvent(ctx *gin.Context){
	var req dto.EventRequestUpdate

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
        return
    }

	event := mapper.ToEventUpdate(req)

	updatedEvent, err := c.Service.Update(id, &event)
	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	 ctx.JSON(http.StatusOK, updatedEvent)
}

func (c *EventController) DeleteEvent(ctx *gin.Context){
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	userID, exits := ctx.Get("user_id")
	if !exits {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	message, err := c.Service.DeleteEvent(eventID, userID.(int))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "not found"})
			return
		}

		if err.Error() == "forbidden"{
			ctx.JSON(http.StatusForbidden, gin.H{"error": "you cannot delete this event"})
        	return
		}


		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}