package patient

import "time"

type Patient struct {
	Address        string    `json:"address" db:"address" gorm:"varchar(250)"`
	Age            int8      `json:"age" db:"age"`
	BloodGroup     string    `json:"blood_group" db:"blood_group" gorm:"varchar(15);column:blood_group"`
	Cheifcomplaint string    `json:"cheif_complaint" db:"cheif_complaint" gorm:"not null;column:cheif_complaint;varchar(250)"`
	CreatedAt      time.Time `json:"created_at" db:"created_at" gorm:"column:created_at"`
	Email          string    `json:"email" db:"email" gorm:"varchar(200)"`
	FirstName      string    `json:"first_name" db:"first_name" gorm:"varchar(50);column:first_name;not null"`
	Gender         int8      `json:"gender" db:"gender" gorm:"not null"`
	ID             int       `json:"patient_id" db:"patient_id" gorm:"column:patient_id;primary_key;AUTO_INCREMENT"`
	LastName       string    `json:"last_name" db:"last_name" gorm:"varchar(50);column:last_name"`
	MedicalHistory string    `json:"medical_history" db:"medical_history" gorm:"column:medical_history;varchar(500)"`
	PhoneNumber    string    `json:"phone_number" db:"phone_number" gorm:"varchar(15);not null;column:phone_number"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at" gorm:"column:updated_at"`
}

func (Patient) TableName() string {
	return "patients"
}
