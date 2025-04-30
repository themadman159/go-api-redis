package productsvc

import "go-api/pkg/model"

func (s *ProductService) GetAll() ([]model.Product, error) {
	products, err := s.ProductRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}
