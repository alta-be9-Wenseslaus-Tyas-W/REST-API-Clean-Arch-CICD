package logins

type Core struct {
	Email    string
	Password string
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Business interface {
	LoginUser(email string, password string) (interface{}, error)
}

type Data interface {
	LoginUser(email string, password string) (int, error)
}
