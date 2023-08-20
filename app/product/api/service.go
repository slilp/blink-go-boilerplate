package product

type Service interface {
	Create(product *ProductEntity) (*ProductEntity, error)
	GetById(id string) (*ProductEntity, error)
	Update(product *ProductEntity) error
	Delete(product *ProductEntity) error
}
