package controller

import (
	"tippers-back/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())
	h := handler{}
	h.Init()
	r.POST("/user",h.RegisterUser)
	return r
}
