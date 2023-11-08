package controller

import (
	"net/http"

	"github.com/GuiSchuffner/chatApi/database"
	"github.com/GuiSchuffner/chatApi/models"
	"github.com/GuiSchuffner/chatApi/utils/token"
	"github.com/gin-gonic/gin"
)

type response struct {
	IsSuccessfully bool             `json:"isSuccessfully"`
	Message        string           `json:"message"`
	UserData       responseUserData `json:"userData"`
}

type responseUserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserData(c *gin.Context) {
	userId, err := token.GetUserIdFromToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response{false, "Invalid token", responseUserData{}})
		return
	}

	var userData responseUserData

	err = database.DB.Model(&models.User{}).Where("id = ?", userId).Take(&userData).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, response{false, "Invalid Data", responseUserData{}})
		return
	}

	c.JSON(http.StatusOK, response{true, "Success", userData})
}
