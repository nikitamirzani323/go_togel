package main

import (
	"gotogel_apibackend/db"
	"gotogel_apibackend/routes"
	"log"
)

func main() {
	db.Init()
	app := routes.Init()
	log.Fatal(app.Listen(":7071"))
}
