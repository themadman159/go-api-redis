package productrepo

import (
	"go-api/pkg/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type IProductRepository interface {
	GetAll() ([]model.Product, error)
}

type ProductRepository struct {
	Database *gorm.DB
	Redis    *redis.Client
}

func NewProductRepository(dbconn *gorm.DB, redis *redis.Client) IProductRepository {
	return &ProductRepository{
		Database: dbconn,
		Redis:    redis,
	}
}
