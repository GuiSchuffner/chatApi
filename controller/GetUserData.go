package controller

import (
	"net/http"

	"github.com/GuiSchuffner/chatApi/database"
	"github.com/GuiSchuffner/chatApi/models"
	"github.com/GuiSchuffner/chatApi/utils/token"
	"github.com/gin-gonic/gin"
)

type apiUserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUserData(c *gin.Context) (apiUserData, error) {
	userId, err := token.GetUserIdFromToken(c.GetHeader("Authorization"))
	if err != nil {
		return apiUserData{}, err
	}

	var userData apiUserData

	err = database.DB.Model(&models.User{}).Where("id = ?", userId).Take(&userData).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, response{false, "Invalid Data", apiUserData{}})
		return apiUserData{}, err
	}

	return userData, nil
}
