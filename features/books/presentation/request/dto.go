package request

import _books "restcleanarch/features/books"

type Book struct {
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	ISBN      string `json:"isbn" form:"isbn"`
	UserId    int    `json:"user_id" form:"user_id"`
}

func ToCore(req Book) _books.Core {
	return _books.Core{
		Title:     req.Title,
		Author:    req.Author,
		Publisher: req.Publisher,
		ISBN:      req.ISBN,
		User: _books.User{
			ID: req.UserId,
		},
	}
}
