package business

import (
	"errors"
	_books "restcleanarch/features/books"
)

type bookUsecase struct {
	bookData _books.Data
}

func NewBookBusiness(bookD _books.Data) _books.Business {
	return &bookUsecase{
		bookData: bookD,
	}
}

func (uc *bookUsecase) GetAllBooks() (response []_books.Core, err error) {
	response, err = uc.bookData.SelectAllBook()
	return response, err
}

func (uc *bookUsecase) GetBookById(id int) (response _books.Core, err error) {
	response, err = uc.bookData.SelectBookById(id)
	return response, err
}

func (uc *bookUsecase) CreateNewBook(data _books.Core) (row int, err error) {
	if data.Title == "" || data.Author == "" || data.Publisher == "" || data.ISBN == "" || data.User.ID == 0 {
		return -1, errors.New("all input must be filled")
	}
	row, err = uc.bookData.InsertNewBook(data)
	return row, err
}

func (uc *bookUsecase) UpdatedBook(id int, data _books.Core) (row int, err error) {
	row, err = uc.bookData.UpdatedBook(id, data)
	return row, err
}

func (uc *bookUsecase) SoftDeleteBook(id int) (row int, err error) {
	row, err = uc.bookData.SoftDeleteBook(id)
	return row, err
}
