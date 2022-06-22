package presentation

import (
	"net/http"
	_books "restcleanarch/features/books"
	_requestBook "restcleanarch/features/books/presentation/request"
	_responseBook "restcleanarch/features/books/presentation/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookBusiness _books.Business
}

func NewBookHandler(bookB _books.Business) *BookHandler {
	return &BookHandler{
		bookBusiness: bookB,
	}
}

func (h *BookHandler) GetAllBook(c echo.Context) error {
	result, err := h.bookBusiness.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseBook.FromCoreList(result),
	})
}

func (h *BookHandler) GetBookById(c echo.Context) error {
	id := c.Param("id")
	idBook, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	result, err := h.bookBusiness.GetBookById(idBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseBook.FromCore(result),
	})
}

func (h *BookHandler) PostNewBook(c echo.Context) error {
	var newBook _requestBook.Book
	errBind := c.Bind(&newBook)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data",
		})
	}
	row, err := h.bookBusiness.CreateNewBook(_requestBook.ToCore(newBook))
	if row == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}
	if row == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success insert data",
	})
}

func (h *BookHandler) PutBookById(c echo.Context) error {
	id := c.Param("id")
	idBook, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	var book _requestBook.Book
	errBind := c.Bind(&book)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data",
		})
	}
	result, err := h.bookBusiness.UpdatedBook(idBook, _requestBook.ToCore(book))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to update data",
		})
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data",
	})
}

func (h *BookHandler) SoftDeleteBookById(c echo.Context) error {
	id := c.Param("id")
	idBook, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	result, err := h.bookBusiness.SoftDeleteBook(idBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to delete data",
		})
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data",
	})
}
