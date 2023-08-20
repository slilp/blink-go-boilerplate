package internal

import (
	order "blink-go-gin-boilerplate/app/order/api"
)


func NewService(repo order.Repository) order.Service {
	return &service{repo: repo}
}

type service struct {
	repo order.Repository
}

func (s *service) Create(order *order.OrderEntity) (*order.OrderEntity, error) {
	return s.repo.Create(order)
}


func (s *service) GetById(id string) (*order.OrderEntity, error) {
	return s.repo.GetById(id)
}

func (s *service) GetByIdWithProduct(id uint) (*order.OrderEntity, error) {
	return s.repo.GetByIdWithProduct(id)
}


func (s *service) Update(order *order.OrderEntity) error {
	return s.repo.Update(order)
}

func (s *service) UpdateProduct(order *order.OrderEntity) error {
	return s.repo.UpdateProduct(order)
}


func (s *service) Delete(order *order.OrderEntity) error {
	return s.repo.Delete(order)
}