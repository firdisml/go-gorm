package main

import (
	"gorm/model"
	"gorm/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
