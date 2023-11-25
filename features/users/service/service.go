package service

import (
	"errors"
	"kupon/features/users"
	haspass "kupon/helper/enkrip"
	"kupon/helper/jwt"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type userService struct {
	repo users.Repository
	h    haspass.HashInterface
}

func New(r users.Repository, h haspass.HashInterface) users.Service {
	return &userService{
		repo: r,
		h:    h,
	}
}

// Register implements users.Service.
func (s *userService) Register(newUser users.User) (users.User, error) {
	// panic("unimplemented")
	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		return users.User{}, errors.New("validation error. name/email/password required")
	}

	ePassword, err := s.h.HashPassword(newUser.Password)

	if err != nil {
		return users.User{}, errors.New("terdapat masalah saat memproses data")
	}

	newUser.Password = ePassword
	result, err := s.repo.Register(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return users.User{}, errors.New("data telah terdaftar pada sistem")
		}
		return users.User{}, errors.New("terjadi kesalahan pada sistem")
	}

	return result, nil

}

// Login implements users.Service.
func (s *userService) Login(email string, password string) (users.User, error) {
	// panic("unimplemented")
	result, err := s.repo.Login(email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return users.User{}, errors.New("data tidak ditemukan")
		}
		return users.User{}, errors.New("terjadi kesalahan pada sistem")
	}
	err = s.h.Compare(result.Password, password)
	if err != nil {
		return users.User{}, errors.New("password salah")
	}

	return result, nil

}

// Update implements users.Service.
func (s *userService) Update(token *golangjwt.Token, updateUser users.User) (users.User, error) {
	// panic("unimplemented")
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return users.User{}, err
	}
	// if id == 0 {
	// 	return users.User{}, errors.New("validation error. invalid ID")
	// }
	// ePassword, err := s.h.HashPassword(updateUser.Password)

	// if err != nil {
	// 	return users.User{}, errors.New("terdapat masalah saat memproses data")
	// }

	// updateUser.Password = ePassword

	result, err := s.repo.Update(userID, updateUser)
	if err != nil {
		return users.User{}, err
	}
	return result, nil
}

// GetUser implements users.Service.
func (s *userService) GetUser() ([]users.User, error) {
	// panic("unimplemented")
	result, err := s.repo.GetUser()
	if err != nil {
		return nil, err
	}
	return result, nil
}
