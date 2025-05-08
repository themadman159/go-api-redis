package productrepo

import (
	"context"
	"encoding/json"
	"go-api/pkg/model"
	"time"
)

func (r *ProductRepository) GetAll() ([]model.Product, error) {

	var products []model.Product

	key := "PRODUCT::GET_ALL"

	productJson, err := r.Redis.Get(context.Background(), key).Result()
	if err == nil {
		err := json.Unmarshal([]byte(productJson), &products)
		if err != nil {
			return nil, err
		}
		return products, nil
	}

	if err := r.Database.Find(&products).Error; err != nil {
		return nil, err
	}

	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	if err := r.Redis.Set(context.Background(), key, data, 20*time.Second).Err(); err != nil {
		return nil, err
	}

	return products, nil
}
