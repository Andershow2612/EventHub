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

type TicketController struct {
	Service *service.TicketService
}

func NewTicketController(db *gorm.DB) *TicketController{
	repo := repository.NewTicketRepository(db)
	eventRepo := repository.NewEventRepository(db)
	service := service.NewTicketService(*repo, *eventRepo)
	return &TicketController{
		Service: service,
	}
}

func (c *TicketController) GetTickets(ctx *gin.Context){
	tickets, err := c.Service.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToTicketResponseList(tickets))
}

func (c *TicketController) CreateTicket(ctx *gin.Context){
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req dto.TicketRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket := mapper.ToTicketEntity(req)

	if err := c.Service.CreateTicket(eventID, ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusCreated, gin.H{"Message": "ticket created"})
}


func (c *TicketController) GetTicketsByEvent(ctx *gin.Context) {
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	tickets, err := c.Service.GetTicketsByEvent(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToTicketResponseList(tickets))
}

// tem que arrumar o retorno desses dois metodos.