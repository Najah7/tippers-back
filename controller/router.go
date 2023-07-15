package controller

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default()
	h := handler{}
	h.Init()
	r.POST("/user",h.RegisterUser)
	return r
}
