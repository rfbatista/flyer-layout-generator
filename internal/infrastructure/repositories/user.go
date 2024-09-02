package repositories

func NewUserRepository() (UserRepository, error) {
	return UserRepository{}, nil
}

type UserRepository struct{}
