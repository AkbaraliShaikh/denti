package patient

type Repository interface {
	Delete(id int) error
	GetAll() ([]*Patient, error)
	GetByID(id int) (*Patient, error)
	Store(u *Patient) error
	Update(u *Patient) error
}
