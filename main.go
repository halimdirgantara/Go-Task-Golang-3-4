package main

import (
	"task-list/routes"
	"github.com/joho/godotenv"
	
)

func main() {
	godotenv.Load()
	routes.Web()
}
