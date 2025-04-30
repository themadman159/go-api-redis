package productroute

import (
	"go-api/handler/producthandler"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func ProductRoute(api fiber.Router, db *gorm.DB, redis *redis.Client) {

	productRoute := producthandler.NewProductHandler(db, redis)
	api.Group("/products").
		Get("/", productRoute.GetAll)
}
