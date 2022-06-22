package business

import (
	"fmt"
	"restcleanarch/middlewares"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockLoginData struct{}

func (mock mockLoginData) LoginUser(email string, password string) (int, error) {
	return 1, nil
}

type mockLoginDataFailed struct{}

func (mock mockLoginDataFailed) LoginUser(email string, password string) (int, error) {
	return 0, fmt.Errorf("Invalid Email or Password")
}

func TestLoginUser(t *testing.T) {
	t.Run("Test Get All Data Success", func(t *testing.T) {
		loginBusiness := NewLoginBusiness(mockLoginData{})
		result, _ := loginBusiness.LoginUser("alta@gmail.com", "123")
		token, errToken := middlewares.CreateToken(1)
		assert.Nil(t, errToken)
		assert.Equal(t, token, result)
	})
	t.Run("Test Get All Data Failed when Email is empty", func(t *testing.T) {
		loginBusiness := NewLoginBusiness(mockLoginData{})
		result, err := loginBusiness.LoginUser("", "123")
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
	t.Run("Test Get All Data Failed when Password is empty", func(t *testing.T) {
		loginBusiness := NewLoginBusiness(mockLoginData{})
		result, err := loginBusiness.LoginUser("alta@gmail.com", "")
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
	t.Run("Test Get All Data Failed when Email and Password is empty", func(t *testing.T) {
		loginBusiness := NewLoginBusiness(mockLoginData{})
		result, err := loginBusiness.LoginUser("", "")
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
	t.Run("Test Get All Data Failed when Token is ", func(t *testing.T) {
		loginBusiness := NewLoginBusiness(mockLoginDataFailed{})
		result, _ := loginBusiness.LoginUser("alta@gmail.com", "123")
		token, errToken := middlewares.CreateToken(0)
		assert.Nil(t, errToken)
		assert.Equal(t, token, result)
	})
}
