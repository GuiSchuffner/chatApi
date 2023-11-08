package routes

import (
	"github.com/GuiSchuffner/chatApi/controller"
	"github.com/GuiSchuffner/chatApi/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(userRoute *gin.RouterGroup) {
	userRoute.Use(middlewares.JwtMiddleware())
	userRoute.GET("/personalData", controller.UserData)
}
