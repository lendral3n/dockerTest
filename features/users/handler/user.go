package handler

import (
	"net/http"
	"kupon/features/users"
	"kupon/helper/jwt"
	"kupon/helper/response"
	"strings"

	gojwt "github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

type userController struct {
	srv users.Service
}

func New(s users.Service) users.Handler {
	return &userController{
		srv: s,
	}
}

// Register implements users.Handler.
func (uc *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		userInput := new(UserRequest)
		err := c.Bind(&userInput)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.WebResponese(http.StatusBadRequest, "error bind data , data not valid", nil))
		}
		inputProses := new(users.User)
		inputProses.Name = userInput.Name
		inputProses.Email = userInput.Email
		inputProses.Password = userInput.Password

		result, err := uc.srv.Register(*inputProses)
		if err != nil {
			if strings.Contains(err.Error(), "validation") {
				c.Logger().Error("ERROR Register, explain:", err.Error())
				return c.JSON(http.StatusBadRequest, response.WebResponese(http.StatusBadRequest, err.Error(), nil))
			}
			return c.JSON(http.StatusInternalServerError, response.WebResponese(http.StatusInternalServerError, "terjadi permasalahan ketika memproses data", err.Error()))
		}

		responseInput := new(UserResponse)
		responseInput.ID = result.ID
		responseInput.Name = result.Name
		responseInput.Email = result.Email

		return c.JSON(http.StatusCreated, response.WebResponese(http.StatusCreated, "success create data", responseInput))
	}
}

// Login implements users.Handler.
func (uc *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		inputLogin := new(LoginRequest)
		err := c.Bind(inputLogin)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.WebResponese(http.StatusBadRequest, "input tidak sesuai", nil))
		}
		result, err := uc.srv.Login(inputLogin.Email, inputLogin.Password)
		if err != nil {
			c.Logger().Error("ERROR Login, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, response.WebResponese(http.StatusBadRequest, err.Error(), nil))
			}
			return c.JSON(http.StatusInternalServerError, response.WebResponese(http.StatusInternalServerError, "terjadi permasalahan ketika memproses data", err.Error()))
		}
		tokenStr, err := jwt.GenerateJWT(result.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.WebResponese(http.StatusInternalServerError, "terjadi permasalahan ketika mengenkripsi data", nil))
		}
		responseLogin := new(LoginResponse)
		responseLogin.ID = result.ID
		responseLogin.Name = result.Name
		responseLogin.Email = result.Email
		responseLogin.Token = tokenStr

		return c.JSON(http.StatusOK, response.WebResponese(http.StatusOK, "success login", responseLogin))
	}

}

// Update implements users.Handler.
func (uc *userController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// idStr := c.Param("user_id")
		// id, err := strconv.Atoi(idStr)
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, response.WebResponese(http.StatusBadRequest, "error. id should be a number", nil))
		// }
		updateInput := new(UserRequest)
		err := c.Bind(&updateInput)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.WebResponese(http.StatusBadRequest, "error binding data", nil))
		}
		inputUpdate := new(users.User)
		inputUpdate.Name = updateInput.Name
		inputUpdate.Email = updateInput.Email
		inputUpdate.Password = updateInput.Password

		// Assuming you have the token available in the context
		token := c.Get("user").(*gojwt.Token)

		result, err := uc.srv.Update(token, *inputUpdate)
		if err != nil {
			if strings.Contains(err.Error(), "validation") {
				c.Logger().Error("ERROR Update, explain:", err.Error())
				return c.JSON(http.StatusBadRequest, response.WebResponese(http.StatusBadRequest, err.Error(), nil))
			}
			return c.JSON(http.StatusInternalServerError, response.WebResponese(http.StatusInternalServerError, "terjadi permasalahan ketika memproses data", err.Error()))
		}
		responseUpdate := new(UserResponse)
		responseUpdate.ID = result.ID
		responseUpdate.Name = result.Name
		responseUpdate.Email = result.Email

		return c.JSON(http.StatusOK, response.WebResponese(http.StatusOK, "User updated successfully", responseUpdate))

	}

}

// GetUser implements users.Handler.
func (uc *userController) GetUser() echo.HandlerFunc {
	// panic("unimplemented")
	return func(c echo.Context) error {
		result, err := uc.srv.GetUser()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.WebResponese(http.StatusInternalServerError, "error", nil))
		}
		var userResponse []UserResponse

		for _, v := range result {
			userResponse = append(userResponse, UserResponse{
				ID:    v.ID,
				Name:  v.Name,
				Email: v.Email,
			})
		}
		return c.JSON(http.StatusOK, response.WebResponese(http.StatusOK, "Seccess read data", userResponse))
	}

}
