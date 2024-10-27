package routes

import (
	"github.com/labstack/echo/v4"
	"starter-template/configs"
	"starter-template/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func NewRoute(c *configs.ProgramConfig, uh users.UserHandlerInterface) *echo.Echo {
	e := echo.New()

	jwtAuth := echojwt.JWT([]byte(c.Secret))

	group := e.Group("/api/v1")

	//	Route Authentication
	group.POST("/register", uh.Register())
	group.POST("/login", uh.Login())
	group.GET("/profile", uh.Profile(), jwtAuth)

	return e
}
