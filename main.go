package main

import (
	"cant_forget/model"
	"cant_forget/routes"
)

func main() {
	model.InitAirtable()
	//model.InitDb()
	routes.InitRouter()
}
