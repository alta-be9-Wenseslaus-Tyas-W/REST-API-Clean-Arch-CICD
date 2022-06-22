package request

import _users "restcleanarch/features/users"

type User struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

//Mengubah Format Data dari luar Ke dalam
func ToCore(req User) _users.Core {
	return _users.Core{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
