package books

import "time"

type Core struct {
	ID        int
	Title     string
	Author    string
	Publisher string
	ISBN      string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}

type User struct {
	ID   int
	Name string
}

type Business interface {
	GetAllBooks() (data []Core, err error)
	GetBookById(id int) (data Core, err error)
	CreateNewBook(data Core) (row int, err error)
	UpdatedBook(id int, data Core) (row int, err error)
	SoftDeleteBook(id int) (row int, err error)
}

type Data interface {
	SelectAllBook() (data []Core, err error)
	SelectBookById(id int) (data Core, err error)
	InsertNewBook(data Core) (row int, err error)
	UpdatedBook(id int, data Core) (row int, err error)
	SoftDeleteBook(id int) (row int, err error)
}
