package main

import (
	"Go/ContainerDIWithDingo/services"
	"fmt"

	"github.com/sarulabs/di"
)

func main() {

	builder, _ := di.NewBuilder()
	builder.Add(di.Def{
		Name: "UserService",
		Build: func(ctn di.Container) (interface{}, error) {
			service := services.NewUserService()
			return service, nil
		},
	})

	ctn := builder.Build()
	userService := ctn.Get("UserService").(services.IUserService)

	user, _ := userService.GetById()

	fmt.Println(user)
}
