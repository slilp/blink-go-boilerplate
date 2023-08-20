package user


type Repository interface {
	Create(profile *UserEntity) (*UserEntity, error)
	GetByUsername(username string) (*UserEntity, error)
}