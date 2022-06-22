package response

import (
	_books "restcleanarch/features/books"
	"time"
)

type Book struct {
	ID        int       `json:"id" form:"id"`
	Title     string    `json:"title" form:"title"`
	Author    string    `json:"author" form:"author"`
	Publisher string    `json:"publisher" form:"publisher"`
	ISBN      string    `json:"isbn" form:"isbn"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	User      User      `json:"user" form:"user"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCore(core _books.Core) Book {
	return Book{
		ID:        core.ID,
		Title:     core.Title,
		Author:    core.Author,
		Publisher: core.Publisher,
		ISBN:      core.ISBN,
		CreatedAt: core.CreatedAt,
		User: User{
			ID:   core.User.ID,
			Name: core.User.Name,
		},
	}
}

func FromCoreList(data []_books.Core) []Book {
	result := []Book{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
