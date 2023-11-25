package main

import (
	"kupon/config"
	ch "kupon/features/coupon/handler"
	couponmodel "kupon/features/coupon/repository"
	cr "kupon/features/coupon/repository"
	cs "kupon/features/coupon/service"
	uh "kupon/features/users/handler"
	ur "kupon/features/users/repository"
	usermodel "kupon/features/users/repository"
	us "kupon/features/users/service"
	helper "kupon/helper/enkrip"

	"kupon/routers"
	"kupon/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.InitConfig()

	if cfg == nil {
		e.Logger.Fatal("tidak bisa start karena ENV error")
		return
	}

	db, err := database.InitMySQL(*cfg)

	if err != nil {
		e.Logger.Fatal("tidak bisa start karena DB error:", err.Error())
		return
	}
	err = db.AutoMigrate(&usermodel.User{}, &couponmodel.Coupon{})
	if err != nil {
		e.Logger.Fatal("gagal migrasi db:", err.Error())
		panic(err)
	}
	userRepoitory := ur.New(db)
	userService := us.New(userRepoitory, helper.New())
	userHandler := uh.New(userService)

	couponRepository := cr.New(db)
	couponService := cs.New(couponRepository)
	couponHandler := ch.New(couponService)

	routers.InitRoute(e, userHandler, couponHandler)
	e.Logger.Fatal(e.Start(":8000"))
}
