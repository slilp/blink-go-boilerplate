package internal

import (
	order "blink-go-gin-boilerplate/app/order/api"

	"gorm.io/gorm"
)


func NewRepository(db *gorm.DB) order.Repository {
	return &repository{db: db}
}

type repository struct {
	db *gorm.DB
}


func (r *repository) Create(order *order.OrderEntity) (*order.OrderEntity, error) {
	tx := r.db.Create(&order)
	return order, tx.Error
}

func (r *repository) GetById(id string) (*order.OrderEntity, error) {
	var order *order.OrderEntity
	tx := r.db.First(&order, "id = ?", id)
	return order, tx.Error
}

func (r *repository) GetByIdWithProduct(id uint) (*order.OrderEntity, error) {
	var order *order.OrderEntity
	tx := r.db.Preload("Products").First(&order, "id = ?", id)
	return order, tx.Error
}


func (r *repository) Update(order *order.OrderEntity) error {
	tx := r.db.Save(&order)
	return tx.Error
}

func (r *repository) UpdateProduct(order *order.OrderEntity) error {

	err := r.db.Model(&order).Association("Products").Replace(order.Products)
	return err
}


func (r *repository) Delete(order *order.OrderEntity) error {
	tx := r.db.Delete(&order)
	return tx.Error
}