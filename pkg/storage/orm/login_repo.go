package orm

import (
	"denti/pkg/logger"
	"denti/pkg/login"

	"github.com/jinzhu/gorm"
)

type loginRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewLoginRepo(db *gorm.DB, log logger.LogInfoFormat) login.Repository {
	return &loginRepo{db, log}
}

func (l *loginRepo) Signin(ln *login.Login) bool {
	var login login.Login

	err := l.db.Raw("SELECT email, password FROM users WHERE email = ? AND password = ?", ln.Email, ln.Password).Scan(&login).Error
	if err != nil {
		l.log.Errorf("signin failed with reason : %v", err)
		return false
	}

	if login.Email == ln.Email && login.Password == ln.Password {
		return true
	}

	return false
}
