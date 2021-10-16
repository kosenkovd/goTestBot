package product

import (
	"errors"
	"fmt"
	"log"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(index int) (*Product, error) {
	log.Printf("Getting file with id %v", index)

	if index < 1 {
		return nil, errors.New("Product id can not be less that 1")
	}

	if index > len(allProducts) {
		return nil, fmt.Errorf("There is no such product %v", index)
	}

	return &allProducts[index-1], nil
}
