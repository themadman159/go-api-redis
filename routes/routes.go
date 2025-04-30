package routes

import (
	"go-api/routes/example"
	"go-api/routes/productroute"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func InitRoutes(app *fiber.App, db *gorm.DB, redis *redis.Client) {
	api := app.Group("/api")

	example.ExampleRoute(api, db)
	productroute.ProductRoute(api, db, redis) // Pass nil for Redis client if not used

}
