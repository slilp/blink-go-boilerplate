package internal

import (
	user "github.com/slilp/blink-go-boilerplate/app/user/api"

	"gorm.io/gorm"
)


func NewRepository(db *gorm.DB) user.Repository {
	return &repository{db: db}
}

type repository struct {
	db *gorm.DB
}


func (r *repository) Create(user *user.UserEntity) (*user.UserEntity, error) {
	tx := r.db.Create(&user)
	return user, tx.Error
}

func (r *repository) GetById(id string) (*user.UserEntity, error) {
	var user *user.UserEntity
	tx := r.db.First(&user, "id = ?", id)
	return user, tx.Error
}

func (r *repository) GetByUsername(username string) (*user.UserEntity, error) {
	var user *user.UserEntity
	tx := r.db.First(&user, "username = ?", username)

	return user, tx.Error
}

func (r *repository) Update(product *user.UserEntity) error {
	tx := r.db.Save(&product)
	return tx.Error
}

func (r *repository) Delete(product *user.UserEntity) error {
	tx := r.db.Delete(&product)
	return tx.Error
}