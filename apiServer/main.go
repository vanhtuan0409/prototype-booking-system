package main

import (
	"apiServer/app"
	"apiServer/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	applyMdw(e)

	env := app.NewEnv()
	h := handlers.NewHandlers(env)
	e.GET("/", h.Echo)
	e.Logger.Fatal(e.Start(":8080"))
}

func applyMdw(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}, lantency=${latency_human}\n",
	}))
}
