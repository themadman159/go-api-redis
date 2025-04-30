package producthandler

import (
	"go-api/pkg/service/productsvc"
	"go-api/utils/responseutil"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type IProductHandler interface {
	GetAll(c *fiber.Ctx) error
}

type ProductHandler struct {
	ProductService productsvc.IProductService
	Response       responseutil.IResponseUtil
}

func NewProductHandler(dbconn *gorm.DB, redis *redis.Client) IProductHandler {
	return &ProductHandler{
		ProductService: productsvc.NewProductService(dbconn, redis),
		Response:       responseutil.NewResponseUtil(),
	}
}
