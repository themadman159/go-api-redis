package productsvc

import (
	"go-api/pkg/model"
	"go-api/pkg/repository/productrepo"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type IProductService interface {
	GetAll() ([]model.Product, error)
}

type ProductService struct {
	ProductRepository IProductService
}

func NewProductService(dbconn *gorm.DB, redis *redis.Client) IProductService {
	return &ProductService{
		ProductRepository: productrepo.NewProductRepository(dbconn, redis),
	}
}
