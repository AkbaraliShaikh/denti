package user

type Repository interface {
	Delete(id string) error
	GetAll() ([]*User, error)
	GetByID(id string) (*User, error)
	Store(u *User) error
	Update(u *User) error
}
