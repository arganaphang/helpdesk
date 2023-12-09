package dto

type CreateUser struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type Login struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
