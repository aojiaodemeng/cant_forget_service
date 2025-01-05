package main

import (
	"cant_forget/model"
	"cant_forget/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
