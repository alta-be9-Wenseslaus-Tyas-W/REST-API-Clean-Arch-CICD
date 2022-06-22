package data

import (
	_logins "restcleanarch/features/logins"

	"gorm.io/gorm"
)

type mysqlLoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(conn *gorm.DB) _logins.Data {
	return &mysqlLoginRepository{
		db: conn,
	}
}

func (repo *mysqlLoginRepository) LoginUser(email string, password string) (int, error) {
	var user = _logins.User{}
	resultEmail := repo.db.Where("email = ?", email).First(&user)
	if resultEmail.Error != nil {
		return 0, resultEmail.Error
	} else if resultEmail.RowsAffected != 1 {
		return 0, resultEmail.Error
	} else {
		id := user.ID
		resultPassword := repo.db.Where("password = ? AND id = ?", password, id).First(&user)
		if resultPassword.Error != nil {
			return 0, resultPassword.Error
		}
		if resultPassword.RowsAffected != 1 {
			return 0, resultPassword.Error
		}
	}
	return user.ID, nil
}
