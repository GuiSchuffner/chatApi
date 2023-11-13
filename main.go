package main

import (
	"github.com/GuiSchuffner/chatApi/database"
	"github.com/GuiSchuffner/chatApi/routes"
	"github.com/GuiSchuffner/chatApi/utils"
)

func main() {
	utils.LoadEnv() // Public .env for this study app
	database.ConnectToDB()
	database.InitChatRooms()
	routes.Routes()
}
