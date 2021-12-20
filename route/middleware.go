package route

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

var IsLoggedInUser = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("user"),
})

var IsLoggedInAdmin = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("admin"),
})

func Init() *echo.Echo {
	routes := echo.New()

	// set logger
	routes.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, time:${time_unix}, uri=${uri}, status=${status}, error=${error}, latency_human=${latency}, bytes_in=${bytes_in}, bytes_out=${bytes_out} \n",
	}))

	// Gzip Compression
	routes.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	routes.Use(middleware.CORS())
	return routes
}