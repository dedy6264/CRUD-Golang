package main

import (
	"net/http"
	"sejutacita/config"
	sejutacita "sejutacita/config"
	"sejutacita/model"
	"sejutacita/usecase"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"

	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {

	config.SetEnv()
	config.SetConnectionDB()
	defer sejutacita.CloseConnectionDB()
	sejutacita.SetConnectionsMongo()
	defer sejutacita.CloseConnectionsMongo()

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	//tanpa grup
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	//grup tanpa middleware
	a := e.Group("/testcoding/v1")
	a.GET("/daftar", func(c echo.Context) error {

		return c.JSON(http.StatusOK, "Hello, TestCoding V1!")
	})
	a.POST("/register", func(c echo.Context) error {
		u := new(model.DataRegister)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		if err := c.Validate(u); err != nil {
			var result model.ResponseGlobal
			result.StatusDateTime = time.Now()
			result.Status = "31"
			result.StatusDesc = err.Error()
			return c.JSON(http.StatusOK, result)
		} else {
			result := usecase.InsertData(*u)
			return c.JSON(http.StatusOK, result)
		}
	})
	a.POST("/login", func(c echo.Context) error {
		u := new(model.DataUser)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		if err := c.Validate(u); err != nil {
			var result model.ResponseGlobal
			result.StatusDateTime = time.Now()
			result.Status = "31"
			result.StatusDesc = err.Error()
			return c.JSON(http.StatusOK, result)
		} else {
			result := usecase.Login(*u)
			return c.JSON(http.StatusOK, result)
		}
	})
	//grup dengan middleware
	b := e.Group("/testcoding/v2")
	b.Use(middleware.JWT([]byte(config.AppKey)))
	b.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
	b.POST("/update_data", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		username := claims["username"].(string)
		role := claims["role"].(string)
		u := new(model.DataUpdate)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		if err := c.Validate(u); err != nil {
			var result model.ResponseGlobal
			result.StatusDateTime = time.Now()
			result.Status = "31"
			result.StatusDesc = err.Error()
			return c.JSON(http.StatusOK, result)
		} else {
			// fmt.Println("kene iki :::")
			result := usecase.UpdateDataUser(*u, username, role)
			return c.JSON(http.StatusOK, result)
		}
	})
	b.POST("/delete_user", func(c echo.Context) error {
		var result model.ResponseGlobal
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		username := claims["username"].(string)
		role := claims["role"].(string)
		if role != "admin" {
			result.StatusDateTime = time.Now()
			result.Status = "41"
			result.StatusDesc = "Sorry, you are not Admin"
			return c.JSON(http.StatusOK, result)
		}
		u := new(model.DeleteUser)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		if err := c.Validate(u); err != nil {
			result.StatusDateTime = time.Now()
			result.Status = "31"
			result.StatusDesc = err.Error()
			return c.JSON(http.StatusOK, result)
		} else {
			// fmt.Println("kene iki :::")
			result := usecase.DeleteDataUser(*u, username)
			return c.JSON(http.StatusOK, result)
		}
	})

	e.Logger.Fatal(e.Start(":9000"))
}
func Login(u model.DataUser) {

}
