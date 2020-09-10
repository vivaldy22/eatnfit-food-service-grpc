package main

import "github.com/vivaldy22/eatnfit-food-service/config"

func main() {
	db, _ := config.InitDB()
	config.RunServer(db)
}
