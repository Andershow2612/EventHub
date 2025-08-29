package routes

import (
	"eventHub.com/internal/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB){
	UserController := controller.NewUserController(db)
	AddressController := controller.NewAddressController(db)
	EventController := controller.NewEventController(db)

	//rota exemplo com os middlewares
	//r.GET("/users", middleware.AuthMiddleware(), middleware.RoleMiddleware(utils.AdmRole), UserController.Getusers)
	r.GET("/users", UserController.Getusers)
	r.GET("/user/:id", UserController.UserById)
	r.POST("/create", UserController.CreateUser)
	r.POST("/login", UserController.Login)
	r.PUT("/updateUser/:id", UserController.UpdateUser)
	r.PUT("/updatePassord/:id", UserController.UpdatePassword)


	//address
	r.GET("/address", AddressController.GetAddress)
	r.GET("/address/:id", AddressController.GetAddresID)
	r.POST("/createAddress", AddressController.CreateAddress)
	r.DELETE("/DeleteAddress/:id", AddressController.DeleteAddress)


	//event
	r.GET("/events", EventController.GetEvents)
	r.GET("/event/:id", EventController.GetEventID)
}