package server

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"mail/config"
	"mail/controllers"
	"mail/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	authController := new(controllers.AuthController)
	parcelsController := new(controllers.ParcelsController)
	itemsController := new(controllers.ItemsController)
	statusesController := new(controllers.StatusesController)

	authGroup := r.Group("api/auth")
	{
		authGroup.POST("/token", authController.Token)
	}

	parcelsGroup := r.Group("api/parcels", middlewares.AuthMiddleware())
	{
		parcelsGroup.GET("/load", parcelsController.Load)
		parcelsGroup.GET("/get/:id", parcelsController.Get)
		parcelsGroup.POST("/create", parcelsController.Create)
	}

	itemsGroup := r.Group("api/items", middlewares.AuthMiddleware())
	{
		itemsGroup.GET("/get/all/:id", itemsController.GetItemsByParcelID)
		itemsGroup.POST("/create/:id", itemsController.Create)
		itemsGroup.POST("/update/:id", itemsController.Update)
	}

	statusesGroup := r.Group("api/status", middlewares.AuthMiddleware())
	{
		statusesGroup.POST("/add/:id", statusesController.Add)
	}

	cfg := config.GetConfig()

	r.Use(static.Serve("/", static.LocalFile(cfg.GetString("frontend.path"), false)))

	return r
}
