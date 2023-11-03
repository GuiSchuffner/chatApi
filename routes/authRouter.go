package routes

import (
	"github.com/GuiSchuffner/chatApi/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controller.Register)
}
