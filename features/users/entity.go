package users

import "time"

type Core struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	GetAllData(param string) (data []Core, err error)
	CreateNewData(data Core) (row int, err error)
	PutData(id int, data Core) (row int, err error)
	DeleteData(id int) (row int, err error)
	GetProfileData(id int) (data Core, err error)
}

type Data interface {
	SelectData(param string) (data []Core, err error)
	InsertData(data Core) (row int, err error)
	UpdateData(id int, data Core) (row int, err error)
	DeleteData(id int) (row int, err error)
	SelectDataById(id int) (data Core, err error)
}
