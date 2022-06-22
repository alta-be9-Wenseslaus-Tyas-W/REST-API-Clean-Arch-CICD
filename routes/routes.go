package routes

import (
	_factory "restcleanarch/factory"
	_middleware "restcleanarch/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter _factory.Presenter) *echo.Echo {
	e := echo.New()

	e.POST("/login", presenter.LoginPresenter.LoginUser)

	e.GET("/users", presenter.UserPresenter.GetAll, _middleware.JWTMiddleware())
	e.GET("/users/:id", presenter.UserPresenter.GetProfileUser, _middleware.JWTMiddleware())
	e.POST("/users", presenter.UserPresenter.CreateNewData)
	e.PUT("/users/:id", presenter.UserPresenter.PutData, _middleware.JWTMiddleware())
	e.DELETE("/users/:id", presenter.UserPresenter.DeleteData, _middleware.JWTMiddleware())

	e.GET("/books", presenter.BookPresenter.GetAllBook)
	e.GET("/books/:id", presenter.BookPresenter.GetBookById)
	e.POST("/books", presenter.BookPresenter.PostNewBook, _middleware.JWTMiddleware())
	e.PUT("/books/:id", presenter.BookPresenter.PutBookById, _middleware.JWTMiddleware())
	e.DELETE("/books/:id", presenter.BookPresenter.SoftDeleteBookById, _middleware.JWTMiddleware())

	return e
}
