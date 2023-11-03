package main

import (
	"github.com/GuiSchuffner/chatApi/database"
	"github.com/GuiSchuffner/chatApi/routes"
)

func main() {
	database.ConnectToDB()
	routes.Routes()
}
