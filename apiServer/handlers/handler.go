package handlers

import (
	"apiServer/app"
)

type Handlers struct {
	Env *app.Env
}

func NewHandlers(env *app.Env) *Handlers {
	return &Handlers{
		Env: env,
	}
}
