package order

type Service interface {
	Create(order *OrderEntity) (*OrderEntity, error)
	GetById(id string) (*OrderEntity, error)
	GetByIdWithProduct(id uint) (*OrderEntity, error)
	Update(order *OrderEntity) error
	UpdateProduct(order *OrderEntity) error
	Delete(order *OrderEntity) error
}
