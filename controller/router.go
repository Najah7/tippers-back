package controller

import (
	"tippers-back/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	h := handler{}
	h.Init()
	{
		r.GET("/user", h.GetUsers)
		r.GET("/user/:id", h.GetUser)
		r.POST("/user", h.RegisterUser)
		r.POST("/login", h.Login)
	}
	{
		r.GET("/restaurant", h.GetRestaurants)
		r.GET("/restaurant/:id", h.GetRestaurant)
	}

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.PATCH("/user/:id", h.UpdateUser)
		authorized.DELETE("/user/:id", h.DeleteUser)
		authorized.POST("/restaurant", h.RegisterRestaurant)
	}

	return r
}
