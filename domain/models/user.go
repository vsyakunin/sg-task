package models

type User struct {
	ID    int64
	Login string
	Role  string
}

const (
	RoleUser     = "user"
	RoleOperator = "operator"
)

var Users = []User{UserOne, UserTwo, Manager}

var UserOne = User{
	ID:    1,
	Login: "userone",
	Role:  RoleUser,
}

var UserTwo = User{
	ID:    2,
	Login: "usertwo",
	Role:  RoleUser,
}

var Manager = User{
	ID:    3,
	Login: "operator",
	Role:  RoleOperator,
}
