package presentation

import (
	"net/http"
	_users "restcleanarch/features/users"
	_requestUser "restcleanarch/features/users/presentation/request"
	_response_user "restcleanarch/features/users/presentation/response"
	"restcleanarch/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness _users.Business
}

func NewUserHandler(business _users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler) GetAll(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	result, err := h.userBusiness.GetAllData("")
	ticket := 0
	for _, v := range result {
		if idToken == v.ID {
			ticket = 1
			break
		} else {
			ticket = 0
		}
	}
	if ticket != 1 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _response_user.FromCoreList(result),
	})
}

func (h *UserHandler) CreateNewData(c echo.Context) error {
	var newUser _requestUser.User
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data",
		})
	}
	result, err := h.userBusiness.CreateNewData(_requestUser.ToCore(newUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success insert data",
	})
}

func (h *UserHandler) PutData(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idUser, errId := strconv.Atoi(id)
	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	var user _requestUser.User
	errBind := c.Bind(&user)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data",
		})
	}
	result, err := h.userBusiness.PutData(idUser, _requestUser.ToCore(user))
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

func (h *UserHandler) DeleteData(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idUser, errId := strconv.Atoi(id)
	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	result, err := h.userBusiness.DeleteData(idUser)
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

func (h *UserHandler) GetProfileUser(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idUser, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to recognized ID",
		})
	}
	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	result, err := h.userBusiness.GetProfileData(idUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get profile",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _response_user.FromCore(result),
	})
}
