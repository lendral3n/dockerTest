package coupon

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Coupon struct {
	ID         uint
	Title      string
	LinkCoupon string
	CodeCoupon string
	Images     string
	UserID     uint
}


type Handler interface {
	Add() echo.HandlerFunc
	GetCoupons() echo.HandlerFunc
}

type Service interface {
	AddCoupon(token *jwt.Token, newCoupon Coupon) (Coupon, error)
	GetCoupons(token *jwt.Token) ([]Coupon, error)
}

type Repository interface {
	AddCoupon(userID uint, newCoupon Coupon) (Coupon, error)
	GetCoupons() ([]Coupon, error)
}
