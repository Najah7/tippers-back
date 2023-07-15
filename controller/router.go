package controller

import (
	"tippers-back/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.AuthMiddleware())
	h := handler{}
	h.Init()
	r.POST("/user", h.RegisterUser)
	r.POST("/login", h.Login)
	return r
}
