package entity

// DB is interface
type DB interface {
	// user
	CreateUser(name, email, pass string) (err error)
	FindUserByID(id int) (User, error)
	UpdateUser(foundUser User) (err error)
	DeleteUser(id int, pass string) (err error)
	// like
	CreateLike(user User, symbol string) (err error)
	CheckLikesSymbol(userID int, symbol string) (Like, error)
	FetchSymbol(symbol string) (err error)
	FindUsersLikes(user User) ([]Like, error)
	FindLikeByID(likeID int) (Like, error)
	DeleteLike(like Like) (err error)
}
