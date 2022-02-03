package config

type Config struct {
	Auth Auth
}

type Auth struct {
	UserOne  User
	UserTwo  User
	Operator User
}

type User struct {
	Login    string
	Password string
}
