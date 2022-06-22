package business

import (
	"fmt"
	_logins "restcleanarch/features/logins"
	_middlewares "restcleanarch/middlewares"
)

type loginUsecase struct {
	loginData _logins.Data
}

func NewLoginBusiness(login _logins.Data) _logins.Business {
	return &loginUsecase{
		loginData: login,
	}
}

func (uc *loginUsecase) LoginUser(email string, password string) (interface{}, error) {
	if email == "" || password == "" {
		return nil, fmt.Errorf("email and password must be filled")
	}
	result, err := uc.loginData.LoginUser(email, password)
	if result == 0 {
		return nil, err
	}
	token, errToken := _middlewares.CreateToken(result)
	if errToken != nil {
		return nil, errToken
	}
	return token, errToken
}
