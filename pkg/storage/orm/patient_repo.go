package orm

import (
	"errors"

	"denti/pkg/patient"

	"github.com/jinzhu/gorm"
)

type patientRepo struct {
	db *gorm.DB
}

func NewPatientRepo(db *gorm.DB) patient.Repository {
	return &patientRepo{db}
}

func (p *patientRepo) Delete(id int) error {
	if p.db.Delete(&patient.Patient{}, "patient_id = ?", id).Error != nil {
		return errors.New("error while deleting the patient")
	}
	return nil
}

func (p *patientRepo) GetAll() ([]*patient.Patient, error) {
	patients := make([]*patient.Patient, 0)
	err := p.db.Find(&patients).Error
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (p *patientRepo) GetByID(id int) (*patient.Patient, error) {
	patient := &patient.Patient{}
	if p.db.Where("patient_id = ?", id).First(&patient).Error != nil {
		return nil, errors.New("patient not found")
	}
	return patient, nil
}

func (p *patientRepo) Store(pt *patient.Patient) error {
	if p.db.Create(&pt).Error != nil {
		return errors.New("error while creating the patient")
	}
	return nil
}

func (p *patientRepo) Update(pt *patient.Patient) error {
	err := p.db.Model(&pt).Updates(patient.Patient{FirstName: pt.FirstName, LastName: pt.LastName, Email: pt.Email, Age: pt.Age, Address: pt.Address, Cheifcomplaint: pt.Cheifcomplaint, PhoneNumber: pt.PhoneNumber, BloodGroup: pt.BloodGroup, MedicalHistory: pt.BloodGroup}).Error
	if err != nil {
		return errors.New("error while updating the patient")
	}
	return nil
}
