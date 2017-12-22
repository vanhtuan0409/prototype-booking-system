package handlers

import (
	"apiServer/app"
	"net/http"

	"github.com/labstack/echo"
)

type Handlers struct {
	Env *app.Env
}

func NewHandlers(env *app.Env) *Handlers {
	return &Handlers{
		Env: env,
	}
}

func (h *Handlers) Echo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}
