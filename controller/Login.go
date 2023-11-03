package controller

import (
	"net/http"

	"github.com/GuiSchuffner/chatApi/database"
	"github.com/GuiSchuffner/chatApi/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	IsSuccessfully bool   `json:"isSuccessfully"`
	Message        string `json:"message"`
}

func Login(c *gin.Context) {
	var input loginInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, loginResponse{false, "Invalid data"})
		return
	}

	inputUser := models.User{}

	inputUser.Email = input.Email
	inputUser.Password = input.Password

}

func checkUserLogin(email string, password string) (string, error) {
	user := models.User{}
	var err error
	err = database.DB.Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", err
	}

	return "", nil
}
