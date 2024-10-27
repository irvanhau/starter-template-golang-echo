package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"starter-template/configs"
	"starter-template/utils/database"
	"starter-template/utils/database/seeds"
)

type Server struct {
	e *echo.Echo
	c *configs.ProgramConfig
}

func (s *Server) RunServer() {
	s.e.Pre(middleware.RemoveTrailingSlash())

	s.e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))
	s.e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"*"}, // example: localhost:3000
			AllowHeaders:     []string{"*"},
			AllowMethods:     []string{"GET", "HEAD", "PUT", "PATCH", "OPTIONS", "DELETE", "POST"},
			AllowCredentials: true,
		}))

	s.e.Use(middleware.Recover())
	s.e.Logger.Debug()
}
func (s *Server) MigrateDB() {
	db := database.InitDB(s.c)
	database.Migrate(db)
}

func (s *Server) SeederDB() {
	db := database.InitDB(s.c)
	for _, seed := range seeds.All() {
		if err := seed.Run(db); err != nil {
			logrus.Errorf("Running seed '%s', failed with error: %s", seed.Name, err.Error())
		}
	}
}

func InitServer(c *configs.ProgramConfig, e *echo.Echo) *Server {
	return &Server{
		e: e,
		c: c,
	}
}
