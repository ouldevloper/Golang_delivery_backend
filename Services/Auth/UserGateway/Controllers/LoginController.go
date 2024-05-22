package Authcontrollers

import (
	Authusecases "delivery/Services/Auth/Application/UseCases"
	"delivery/Services/Auth/Domain/DTOs"
	SharedUtils "delivery/Services/Sahared/Application/Utils"

	"github.com/gin-gonic/gin"
)

type Authcontroller struct {
}

func (obj *Authcontroller) Login(ctx *gin.Context) {
	var data *DTOs.LoginDTO
	SharedUtils.ParseBody(ctx.Request, data)
	result := Authusecases.LoginuserCase(data.getLogin(), data.getPassword())
	SharedUtils.Success(ctx, data, 200)
}