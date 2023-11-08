package routes

import (
	"github.com/GuiSchuffner/chatApi/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(authRoute *gin.RouterGroup) {
	authRoute.POST("/register", controller.Register)
	authRoute.POST("/login", controller.Login)
}
