package request

import _logins "restcleanarch/features/logins"

type Login struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(req Login) _logins.Core {
	return _logins.Core{
		Email:    req.Email,
		Password: req.Password,
	}
}
