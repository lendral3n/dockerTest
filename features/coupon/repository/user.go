package repository

import (
	"kupon/features/coupon"
	"kupon/helper/cloudinery"

	"github.com/cloudinary/cloudinary-go"
	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	Title      string
	LinkCoupon string
	CodeCoupon string
	Images     string
	UserID     uint
}

type coupounQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) coupon.Repository {
	return &coupounQuery{
		db: db,
	}
}

// AddCoupon implements coupon.Repository.
func (r *coupounQuery) AddCoupon(userID uint, newCoupon coupon.Coupon) (coupon.Coupon, error) {
	// panic("unimplemented")
	inputCoupon := new(Coupon)
	inputCoupon.Title = newCoupon.Title
	inputCoupon.LinkCoupon = newCoupon.LinkCoupon
	inputCoupon.CodeCoupon = newCoupon.CodeCoupon
	inputCoupon.UserID = userID
	// inputCoupon.Images = newCoupon.Images

	cld, err := cloudinary.NewFromURL("cloudinary://533421842888945:Oish5XyXkCiiV6oTW2sEo0lEkGg@dlxvvuhph")
	if err != nil {
		return coupon.Coupon{}, err
	}

	imageURL, err := cloudinery.UploadToCloudinary(cld, inputCoupon.Images)
	if err != nil {
		return coupon.Coupon{}, err
	}

	inputCoupon.Images = imageURL

	tx := r.db.Create(&inputCoupon)
	if tx.Error != nil {
		return coupon.Coupon{}, tx.Error
	}
	newCoupon.ID = inputCoupon.ID
	return newCoupon, nil
}

// GetCoupons implements coupon.Repository.

// GetCoupons implements coupon.Repository.
func (r *coupounQuery) GetCoupons() ([]coupon.Coupon, error) {
	// panic("unimplemented")
	var couponData []Coupon
	tx := r.db.Find(&couponData)
	if tx.Error != nil {
		return []coupon.Coupon{}, tx.Error
	}
	var listCoupons []coupon.Coupon
	for _, data := range couponData {
		couponData = append(couponData, Coupon{
			Model: gorm.Model{
				ID:        data.ID,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
				DeletedAt: data.DeletedAt,
			},
			Title:      data.Title,
			LinkCoupon: data.LinkCoupon,
			CodeCoupon: data.CodeCoupon,
			Images:     data.Images,
			UserID:     data.UserID,
		})
		listCoupons = append(listCoupons, coupon.Coupon{
			ID:         data.ID,
			Title:      data.Title,
			LinkCoupon: data.LinkCoupon,
			CodeCoupon: data.CodeCoupon,
			Images:     data.Images,
			UserID:     data.UserID,
		})
	}
	return listCoupons, nil
}
