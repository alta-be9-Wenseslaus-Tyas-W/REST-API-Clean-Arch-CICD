package data

import (
	_books "restcleanarch/features/books"

	"gorm.io/gorm"
)

type mysqlBookRepository struct {
	db *gorm.DB
}

func NewBookRepository(conn *gorm.DB) _books.Data {
	return &mysqlBookRepository{
		db: conn,
	}
}

func (repo *mysqlBookRepository) SelectAllBook() ([]_books.Core, error) {
	var dataBooks []Book
	result := repo.db.Preload("User").Find(&dataBooks)
	if result.Error != nil {
		return []_books.Core{}, result.Error
	}
	return toCoreList(dataBooks), nil
}

func (repo *mysqlBookRepository) SelectBookById(id int) (_books.Core, error) {
	var profile Book
	result := repo.db.Model(&Book{}).First(&profile, "id = ?", id)
	if result.Error != nil {
		return _books.Core{}, result.Error
	}
	return profile.toCore(), nil
}

func (repo *mysqlBookRepository) InsertNewBook(data _books.Core) (int, error) {
	var newBook = fromCore(data)
	result := repo.db.Create(&newBook)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlBookRepository) UpdatedBook(id int, data _books.Core) (int, error) {
	var update = fromCore(data)
	result := repo.db.Model(&Book{}).Where("id = ?", id).Updates(&update)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlBookRepository) SoftDeleteBook(id int) (int, error) {
	result := repo.db.Where("id = ?", id).Delete(&Book{})
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}
