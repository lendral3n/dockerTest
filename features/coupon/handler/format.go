package handler

import "mime/multipart"

type CouponRequest struct {
	Title      string                `json:"title" form:"title"`
	LinkCoupon string                `json:"link_coupon" form:"link_coupon"`
	CodeCoupon string                `json:"code_coupon" form:"code_coupon"`
	Images     *multipart.FileHeader `json:"images" form:"images"`
}

type CouponResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	LinkCoupon string `json:"link_coupon"`
	CodeCoupon string `json:"code_coupon"`
	Images     string `json:"images"`
}
