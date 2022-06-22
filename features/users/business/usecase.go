package business

import (
	"fmt"
	_users "restcleanarch/features/users"
)

type userUsecase struct {
	userData _users.Data
}

func NewUserBusiness(usrData _users.Data) _users.Business {
	return &userUsecase{
		userData: usrData,
	}
}

func (uc *userUsecase) GetAllData(param string) (resp []_users.Core, err error) {
	resp, err = uc.userData.SelectData(param)
	return resp, err
}

func (uc *userUsecase) CreateNewData(data _users.Core) (row int, err error) {
	if data.Name == "" || data.Email == "" || data.Password == "" {
		return 0, fmt.Errorf("all input must be filled")
	}
	row, err = uc.userData.InsertData(data)
	return row, err
}

func (uc *userUsecase) PutData(id int, data _users.Core) (row int, err error) {
	//var update = _users.Core{}
	row, err = uc.userData.UpdateData(id, data)
	return row, err
}

func (uc *userUsecase) DeleteData(id int) (row int, err error) {
	row, err = uc.userData.DeleteData(id)
	return row, err
}

func (uc *userUsecase) GetProfileData(id int) (resp _users.Core, err error) {
	resp, err = uc.userData.SelectDataById(id)
	return resp, err
}
