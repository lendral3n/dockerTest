package service

import (
	"errors"
	"kupon/features/coupon"
	"kupon/helper/jwt"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type couponService struct {
	repo coupon.Repository
}

func New(r coupon.Repository) coupon.Service {
	return &couponService{
		repo: r,
	}
}

// AddCoupon implements coupon.Service.
func (s *couponService) AddCoupon(token *golangjwt.Token, newCoupon coupon.Coupon) (coupon.Coupon, error) {
	// panic("unimplemented")

	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return coupon.Coupon{}, err
	}
	result, err := s.repo.AddCoupon(userID, newCoupon)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return coupon.Coupon{}, errors.New("barang sudah pernah diinputkan")
		}
		return coupon.Coupon{}, errors.New("terjadi kesalahan pada server")
	}
	return result, nil
}

// GetCoupons implements coupon.Service.
func (s *couponService) GetCoupons(token *golangjwt.Token) ([]coupon.Coupon, error) {
	// panic("unimplemented")
	result, err := s.repo.GetCoupons()
	if err != nil {
		return nil, err
	}
	return result, nil
}
