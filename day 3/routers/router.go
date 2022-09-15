package routers

import (
	"go-restful-demo/controllers"
	m "go-restful-demo/middlewares"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New() *echo.Echo {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	m.LogMiddlewares(e)

	publicRoutes := e.Group("/v1")
	{
		publicRoutes.POST("/login", controllers.LoginUserControllers)

		publicRoutes.POST("/users", controllers.CreateUserControllers)

		publicRoutes.GET("/books", controllers.GetBookControllers)
		publicRoutes.GET("/books/:id", controllers.GetBookByIdControllers)
	}

	authenticatedRoutes := e.Group("/v1")
	{
		secret := os.Getenv("JWT_SECRET")
		authenticatedRoutes.Use(middleware.JWT([]byte(secret)))

		authenticatedRoutes.GET("/users", controllers.GetUserControllers)
		authenticatedRoutes.GET("/users/:id", controllers.GetUserByIdControllers)
		authenticatedRoutes.PUT("/users/:id", controllers.UpdateUserControllers)
		authenticatedRoutes.DELETE("/users/:id", controllers.DeleteUserControllers)

		authenticatedRoutes.POST("/books", controllers.CreateBookControllers)
		authenticatedRoutes.PUT("/books/:id", controllers.UpdateBookControllers)
		authenticatedRoutes.DELETE("/books/:id", controllers.DeleteBookControllers)
	}

	return e
}
