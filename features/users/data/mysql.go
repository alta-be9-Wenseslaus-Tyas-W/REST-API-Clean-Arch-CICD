package data

import (
	_users "restcleanarch/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserReporsitory(conn *gorm.DB) _users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

//SQL using GORM to GET all data User
func (repo *mysqlUserRepository) SelectData(data string) (response []_users.Core, err error) {
	var dataUsers []User
	result := repo.db.Find(&dataUsers)
	if result.Error != nil {
		return []_users.Core{}, result.Error
	}
	return toCoreList(dataUsers), nil
}

//SQL using GORM to INSERT new User
func (repo *mysqlUserRepository) InsertData(data _users.Core) (int, error) {
	var new = fromCore(data)
	result := repo.db.Create(&new)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) UpdateData(id int, data _users.Core) (int, error) {
	var update = fromCore(data)
	result := repo.db.Model(&User{}).Where("id = ?", id).Updates(&update)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) DeleteData(id int) (int, error) {
	result := repo.db.Where("id = ?", id).Delete(&User{})
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) SelectDataById(id int) (_users.Core, error) {
	var profile User
	result := repo.db.Model(&User{}).First(&profile, "id = ?", id)
	if result.Error != nil {
		return _users.Core{}, result.Error
	}
	return profile.toCore(), nil
}
