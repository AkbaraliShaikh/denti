package patient

type Service interface {
	Delete(id int) error
	GetAll() ([]*Patient, error)
	GetByID(id int) (*Patient, error)
	Store(u *Patient) error
	Update(u *Patient) error
}

type patientService struct {
	repo Repository
}

func NewPatientService(repo Repository) Service {
	return &patientService{
		repo: repo,
	}
}

func (svc *patientService) Delete(id int) error { return svc.repo.Delete(id) }

func (svc *patientService) GetAll() ([]*Patient, error) { return svc.repo.GetAll() }

func (svc *patientService) GetByID(id int) (*Patient, error) { return svc.repo.GetByID(id) }

func (svc *patientService) Store(u *Patient) error { return svc.repo.Store(u) }

func (svc *patientService) Update(u *Patient) error { return svc.repo.Update(u) }
