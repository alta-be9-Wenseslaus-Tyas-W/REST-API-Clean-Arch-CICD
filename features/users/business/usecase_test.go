package business

import (
	"fmt"
	"restcleanarch/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

//mock data success case
type mockUserData struct{}

func (mock mockUserData) SelectData(param string) (data []users.Core, err error) {
	return []users.Core{
		{ID: 1, Name: "alta", Email: "alta@gmail.com", Password: "123"},
	}, nil
}
func (mock mockUserData) InsertData(data users.Core) (row int, err error) {
	return 1, nil
}
func (mock mockUserData) UpdateData(id int, data users.Core) (row int, err error) {
	return 1, nil
}
func (mock mockUserData) DeleteData(id int) (row int, err error) {
	return 1, nil
}
func (mock mockUserData) SelectDataById(id int) (data users.Core, err error) {
	return users.Core{
		ID: 1, Name: "alta", Email: "alta@gmail.com", Password: "123",
	}, nil
}

type mockUserDataFailed struct{}

func (mock mockUserDataFailed) SelectData(param string) (data []users.Core, err error) {
	return nil, fmt.Errorf("failed to get all data")
}
func (mock mockUserDataFailed) InsertData(data users.Core) (row int, err error) {
	return 0, fmt.Errorf("failed to insert data")
}
func (mock mockUserDataFailed) UpdateData(id int, data users.Core) (row int, err error) {
	return 0, fmt.Errorf("failed to update data")
}
func (mock mockUserDataFailed) DeleteData(id int) (row int, err error) {
	return 0, fmt.Errorf("failed to delete data")
}
func (mock mockUserDataFailed) SelectDataById(id int) (data users.Core, err error) {
	return users.Core{Name: ""}, fmt.Errorf("failed to get data")
}

func TestGetAllData(t *testing.T) {
	t.Run("Test Get All Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.GetAllData("")
		assert.Nil(t, err)
		assert.Equal(t, "alta", result[0].Name)
	})
	t.Run("Test Get All Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.GetAllData("")
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
func TestCreateNewData(t *testing.T) {
	t.Run("Test Insert Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		newUser := users.Core{
			Name:     "alta",
			Email:    "alta@gmail.com",
			Password: "123",
		}
		result, err := userBusiness.CreateNewData(newUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Insert Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta",
			Email:    "alta@gmail.com",
			Password: "123",
		}
		result, err := userBusiness.CreateNewData(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
	t.Run("Test Insert Data Failed when Email Empty", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta",
			Password: "123",
		}
		result, err := userBusiness.CreateNewData(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestPutData(t *testing.T) {
	t.Run("Test Updated Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		updateUser := users.Core{
			ID:       1,
			Name:     "alta",
			Email:    "alta@gmail.com",
			Password: "123",
		}
		result, err := userBusiness.PutData(updateUser.ID, updateUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Updated Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		updateUser := users.Core{
			ID:       1,
			Name:     "alta",
			Email:    "alta@gmail.com",
			Password: "123",
		}
		result, err := userBusiness.PutData(updateUser.ID, updateUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Test Delete Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.DeleteData(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Delete Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.DeleteData(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestGetProfileData(t *testing.T) {
	t.Run("Test Get Profile Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.GetProfileData(1)
		assert.Nil(t, err)
		assert.Equal(t, "alta", result.Name)
	})
	t.Run("Test Get Profile Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.GetProfileData(1)
		assert.NotNil(t, err)
		assert.Equal(t, "", result.Name)
	})
}
