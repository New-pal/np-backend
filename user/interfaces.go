package user

type userInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() []byte
}
