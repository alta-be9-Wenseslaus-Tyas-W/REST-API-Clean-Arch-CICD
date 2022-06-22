package factory

import (
	_bookBusiness "restcleanarch/features/books/business"
	_bookData "restcleanarch/features/books/data"
	_bookPresentation "restcleanarch/features/books/presentation"
	_loginBusiness "restcleanarch/features/logins/business"
	_loginData "restcleanarch/features/logins/data"
	_loginPresentation "restcleanarch/features/logins/presentation"
	_userBusiness "restcleanarch/features/users/business"
	_userData "restcleanarch/features/users/data"
	_userPresentation "restcleanarch/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter  *_userPresentation.UserHandler
	BookPresenter  *_bookPresentation.BookHandler
	LoginPresenter *_loginPresentation.LoginHandler
}

func InitFactory(db *gorm.DB) Presenter {
	//Init For User
	userData := _userData.NewUserReporsitory(db)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	//Init For Book
	bookData := _bookData.NewBookRepository(db)
	bookBusiness := _bookBusiness.NewBookBusiness(bookData)
	bookPresentation := _bookPresentation.NewBookHandler(bookBusiness)

	loginData := _loginData.NewLoginRepository(db)
	loginBusiness := _loginBusiness.NewLoginBusiness(loginData)
	loginPresentation := _loginPresentation.NewLoginHandler(loginBusiness)

	return Presenter{
		UserPresenter:  userPresentation,
		BookPresenter:  bookPresentation,
		LoginPresenter: loginPresentation,
	}
}
