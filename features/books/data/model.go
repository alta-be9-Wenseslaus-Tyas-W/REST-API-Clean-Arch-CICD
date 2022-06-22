package data

import (
	_books "restcleanarch/features/books"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	ISBN      string `json:"isbn" form:"isbn"`
	UserID    uint   `json:"user_id" form:"user_id"`
	User      User
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Books    []Book
}

func (data *Book) toCore() _books.Core {
	return _books.Core{
		ID:        int(data.ID),
		Title:     data.Title,
		Publisher: data.Publisher,
		ISBN:      data.ISBN,
		Author:    data.Author,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		User: _books.User{
			ID:   int(data.User.ID),
			Name: data.User.Name,
		},
	}
}

func toCoreList(data []Book) []_books.Core {
	result := []_books.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core _books.Core) Book {
	return Book{
		Title:     core.Title,
		Author:    core.Author,
		Publisher: core.Publisher,
		ISBN:      core.ISBN,
		UserID:    uint(core.User.ID),
	}
}
