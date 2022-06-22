package presentation

import (
	"net/http"
	_logins "restcleanarch/features/logins"
	_requestLogin "restcleanarch/features/logins/presentation/request"

	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	loginBusiness _logins.Business
}

func NewLoginHandler(login _logins.Business) *LoginHandler {
	return &LoginHandler{
		loginBusiness: login,
	}
}

func (h *LoginHandler) LoginUser(c echo.Context) error {
	var login = _requestLogin.Login{}
	errBind := c.Bind(&login)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data",
		})
	}
	new := _requestLogin.ToCore(login)
	token, err := h.loginBusiness.LoginUser(new.Email, new.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "email or password incorrect",
		})
	}
	data := map[string]interface{}{
		"token": token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"data":    data,
	})
}
