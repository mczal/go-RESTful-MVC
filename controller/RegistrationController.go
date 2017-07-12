package controller

import (
	"go-restful-mvc/model"
	"go-restful-mvc/service"

	"github.com/go-siris/siris/context"
)

func PostNewUser(ctx context.Context) {
	var user model.User
	ctx.ReadJSON(&user)

	stat, result := service.NewUser(&user)

	ctx.StatusCode(stat)
	ctx.JSON(result)
}
