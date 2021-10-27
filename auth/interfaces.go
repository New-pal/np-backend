package auth

type userInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() []byte
}

type userRepositoryInterface interface {
	GetUserByEmail(email string) (interface{ userInterface }, error)
}

type userServiceInterface interface {
	CreateUser(email string, name string, password string) (interface{ userInterface }, error)
}
