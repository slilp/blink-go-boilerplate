package internal

import (
	product "github.com/slilp/blink-go-boilerplate/app/product/api"

	"gorm.io/gorm"
)


func NewRepository(db *gorm.DB) product.Repository {
	return &repository{db: db}
}

type repository struct {
	db *gorm.DB
}


func (r *repository) Create(product *product.ProductEntity) (*product.ProductEntity, error) {
	tx := r.db.Create(&product)
	return product, tx.Error
}

func (r *repository) GetById(id string) (*product.ProductEntity, error) {
	var profile *product.ProductEntity
	tx := r.db.First(&profile, "id = ?", id)
	return profile, tx.Error
}

func (r *repository) Update(product *product.ProductEntity) error {
	tx := r.db.Save(&product)
	return tx.Error
}

func (r *repository) Delete(product *product.ProductEntity) error {
	tx := r.db.Delete(&product)
	return tx.Error
}