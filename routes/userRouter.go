package routes

import (
	"github.com/GuiSchuffner/chatApi/controller"
	"github.com/GuiSchuffner/chatApi/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userRoute := r.Group("/user")
	userRoute.Use(middlewares.JwtMiddleware())
	userRoute.GET("/personalData", controller.UserData)
}
