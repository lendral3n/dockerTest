package routers

import (
	"kupon/features/coupon"
	"kupon/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc users.Handler, cc coupon.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uc)
	routeCoupon(e, cc)
}

func routeUser(e *echo.Echo, uc users.Handler) {
	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())
	e.PUT("/users", uc.Update(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/users", uc.GetUser())
}

func routeCoupon(e *echo.Echo, cc coupon.Handler) {
	e.POST("/coupons", cc.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/coupons", cc.GetCoupons(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
