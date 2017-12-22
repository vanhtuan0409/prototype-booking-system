package main

import (
	"apiServer/app"
	"apiServer/entities"
	"fmt"
)

const (
	N_RESOURCE = 36
)

func main() {
	env := app.NewEnv()
	for i := 0; i < N_RESOURCE; i++ {
		p := &entities.Resource{
			Availability: true,
		}
		if err := env.DB.Save(p).Error; err != nil {
			fmt.Println(err)
			continue
		}
	}
}
