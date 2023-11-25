package users

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}
type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetUser()echo.HandlerFunc
}

type Service interface {
	Register(newUser User) (User, error)
	Login(email string, password string) (User, error)
	// Update(token *jwt.Token, id uint, updateUser User) (User, error)
	Update(token *jwt.Token, updateUser User) (User, error)
	GetUser()([]User,error)
}

type Repository interface {
	Register(newUser User) (User, error)
	Login(email string) (User, error)
	Update(id uint, updateUser User) (User, error)
	GetUser()([]User,error)
}
