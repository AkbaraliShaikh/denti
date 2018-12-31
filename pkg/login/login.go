package login

type Login struct {
	Email    string `json:"email" db:"email" gorm:"not null"`
	Password string `json:"password" db:"password" gorm:"not null"`
}
