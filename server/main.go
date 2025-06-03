package main

import "github.com/matsumo-and/dr-score-gen/infrastructure/router"

func main() {
	router := router.Router()
	router.Run(":8080")
}
