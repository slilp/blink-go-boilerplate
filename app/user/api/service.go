package user

type Service interface {
	Create(user *UserEntity) (*UserEntity, error)
	GetByUsername(username string) (*UserEntity, error)
}
