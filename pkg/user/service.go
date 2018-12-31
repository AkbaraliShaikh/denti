package user

type Service interface {
	Delete(id string) error
	GetAll() ([]*User, error)
	GetByID(id string) (*User, error)
	Store(u *User) error
	Update(u *User) error
}

type userService struct {
	repo Repository
}

func NewUserService(repo Repository) Service {
	return &userService{
		repo: repo,
	}
}

func (svc *userService) Delete(id string) error { return svc.repo.Delete(id) }

func (svc *userService) GetAll() ([]*User, error) { return svc.repo.GetAll() }

func (svc *userService) GetByID(id string) (*User, error) { return svc.repo.GetByID(id) }

func (svc *userService) Store(u *User) error { return svc.repo.Store(u) }

func (svc *userService) Update(u *User) error { return svc.repo.Update(u) }
