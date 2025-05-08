package main

import (
	"go-api/cmd/database"
	serverconfig "go-api/config/server_config"
	"go-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	redis, rerr := database.InitRedisDatabase()
	if rerr != nil {
		log.Fatalf("Error initializing Redis: %v", err)
	}

	if err := database.Migration(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	app := fiber.New()

	routes.InitRoutes(app, db, redis)

	app.Listen(":" + serverconfig.ServerConfig().PORT)
}
