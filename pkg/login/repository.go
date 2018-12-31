package login

type Repository interface {
	Signin(*Login) bool
}
