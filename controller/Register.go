package controller

import (
	"net/http"

	"github.com/GuiSchuffner/chatApi/database"
	"github.com/GuiSchuffner/chatApi/models"
	helper "github.com/GuiSchuffner/chatApi/utils/helpers"
	"github.com/gin-gonic/gin"
)

type registerInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerResponse struct {
	IsSuccessfully bool   `json:"isSuccessfully"`
	Message        string `json:"message"`
}

func Register(c *gin.Context) {
	var input registerInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, registerResponse{false, "Invalid data"})
		return
	}

	newUser := models.User{}
	newUser.Name = input.Name
	newUser.Email = input.Email
	newUser.Password = input.Password

	err = database.DB.Create(&newUser).Error

	if err != nil {
		if helper.IsDuplicatedKeyError(err) {
			c.JSON(http.StatusOK, registerResponse{false, "Email already in use"})
		} else {
			c.JSON(http.StatusInternalServerError, registerResponse{false, "Internal server error"})
		}
	} else {
		c.JSON(http.StatusOK, registerResponse{true, "User created successfully"})
	}
}
