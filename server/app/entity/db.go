package entity

// DB is interface
type DB interface {
	CreateUser(name, email, pass string) (err error)
	FindUserByID(id int) (User, error)
	UpdateUser(foundUser User) (err error)
	DeleteUser(id int, pass string) (err error)
}
