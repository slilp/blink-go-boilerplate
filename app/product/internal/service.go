package internal

import product "github.com/slilp/blink-go-boilerplate/app/product/api"


func NewService(repo product.Repository) product.Service {
	return &service{repo: repo}
}

type service struct {
	repo product.Repository
}

func (s *service) Create(product *product.ProductEntity) (*product.ProductEntity, error) {
	return s.repo.Create(product)
}

func (s *service) GetById(id string) (*product.ProductEntity, error) {
	return s.repo.GetById(id)
}

func (s *service) Update(product *product.ProductEntity) error {
	return s.repo.Update(product)
}

func (s *service) Delete(product *product.ProductEntity) error {
	return s.repo.Delete(product)
}