package handler

import (
	"net/http"
	"kupon/features/coupon"
	"kupon/helper/response"
	"strings"

	gojwt "github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

type couponController struct {
	srv coupon.Service
}

func New(s coupon.Service) coupon.Handler {
	return &couponController{
		srv: s,
	}
}

// Add implements coupon.Handler.
func (cc *couponController) Add() echo.HandlerFunc {
	// panic("unimplemented")
	return func(c echo.Context) error {
		input := new(CouponRequest)
		err := c.Bind(input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.WebResponese(http.StatusBadRequest, "input yang diberikan tidak sesuai", nil))
		}
		inputProses := new(coupon.Coupon)
		inputProses.Title = input.Title
		inputProses.LinkCoupon = input.LinkCoupon
		inputProses.CodeCoupon = input.CodeCoupon
		// inputProses.Images = im

		token := c.Get("user").(*gojwt.Token)
		result, err := cc.srv.AddCoupon(token, *inputProses)
		if err != nil {
			if strings.Contains(err.Error(), "validation") {
				c.Logger().Error("ERROR Uppload kupon, explain:", err.Error())
				return c.JSON(http.StatusBadRequest, response.WebResponese(http.StatusBadRequest, err.Error(), nil))
			}
			return c.JSON(http.StatusInternalServerError, response.WebResponese(http.StatusInternalServerError, "terjadi permasalahan ketika memproses data", err.Error()))
		}
		responseInput := new(CouponResponse)
		responseInput.ID = result.ID
		responseInput.Title = result.Title
		responseInput.LinkCoupon = result.LinkCoupon
		responseInput.CodeCoupon = result.CodeCoupon
		responseInput.Images = result.Images

		return c.JSON(http.StatusCreated, response.WebResponese(http.StatusCreated, "success create data kupon", responseInput))

	}
}

// GetCoupons implements coupon.Handler.
func (cc *couponController) GetCoupons() echo.HandlerFunc {
	// panic("unimplemented")
	return func(c echo.Context) error {
		token := c.Get("user").(*gojwt.Token)

		result, err := cc.srv.GetCoupons(token)
		if err != nil {
			if strings.Contains(err.Error(), "validation") {
				c.Logger().Error("ERROR Update, explain:", err.Error())
				return c.JSON(http.StatusBadRequest, response.WebResponese(http.StatusBadRequest, err.Error(), nil))
			}
			return c.JSON(http.StatusInternalServerError, response.WebResponese(http.StatusInternalServerError, "terjadi permasalahan ketika memproses data", err.Error()))
		}
		var responseCoupon []CouponResponse
		for _, v := range result {
			responseCoupon = append(responseCoupon, CouponResponse{
				ID:         v.ID,
				Title:      v.Title,
				LinkCoupon: v.LinkCoupon,
				CodeCoupon: v.CodeCoupon,
				Images:     v.Images,
			})
		}
		return c.JSON(http.StatusOK, response.WebResponese(http.StatusOK, "Seccess read data", responseCoupon))
	}

}
