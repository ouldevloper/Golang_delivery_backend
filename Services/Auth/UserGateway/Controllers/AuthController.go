package auth_usergetway_controllers

import (
	authu_services "delivery/Services/Auth/Application/Services"
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	auth_infrastructure_repository "delivery/Services/Auth/Infrastructure/Repositories"
	auth_requests "delivery/Services/Auth/UserGateway/Requests"
	shared_configs "delivery/Services/Shared/Application/Configs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AuthController struct {
	Validate *validator.Validate
	service  *authu_services.AuthService
}

func NewAuthController() *AuthController {
	config, _ := shared_configs.GetConfig()
	return &AuthController{
		Validate: validator.New(),
		service:  authu_services.NewAuthService(auth_infrastructure_repository.NewAuthRepository(config)),
	}
}

func (obj AuthController) Login(ctx *gin.Context) {

	request := &auth_requests.LoginRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   errorMessage,
			"data":    nil,
		})
		return
	}

	accessToken, _ := obj.service.Login(auth_domain_dtos.LoginDTO{
		Login:    request.Login,
		Password: request.Password,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   accessToken,
		"error":   nil,
		"data":    nil,
	})
}

func (obj AuthController) Register(ctx *gin.Context) {
	request := &auth_requests.RegisterRequest{}
	if err := ctx.BindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	fmt.Print(request)

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   errorMessage,
			"data":    nil,
		})
		return
	}

	accessToken, _ := obj.service.Register(auth_domain_dtos.RegisterDTO{
		UserName
		FullName
		Email
		Address
		Password
		UserType
		UserLevel
		IsOnline
		IsLocked
		LastSeen
	})



	// var dto DTOs.UserDTO
	// SharedUtils.ParseBody(ctx.Request, &dto)
	// fmt.Println("level Register Controller : ", ctx.Request.Body)
	// res := userCase.Register(dto)
	// SharedUtils.Success(ctx, res, 200)
}
