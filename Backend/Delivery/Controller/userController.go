package Controller

import (
	"DSAShare/Domain"
	"DSAShare/UseCase"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct{
	UserUseCase UseCase.IUserUseCase
	V *validator.Validate
	TokenService UseCase.ITokenService
}

func NewUserController(uuc UseCase.IUserUseCase, ts UseCase.ITokenService) *UserController{
	return &UserController{
		UserUseCase: uuc,
		V: validator.New(),
		TokenService: ts,
	}
}

func (c *UserController) Register(ctx *gin.Context){
	user := Domain.User{}

	err := ctx.ShouldBindJSON(&user)

	if err != nil{
		ctx.JSON(400, gin.H{"error" : "invalid request payload"})
		return
	}

	err = c.V.Struct(user)
	if err != nil{
		ctx.JSON(400, gin.H{"error" : "invalid request payload"})
		return
	}

	code, err := c.UserUseCase.Register(&user)

	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" : "registration successful. Verification has been sent to the Email"})
}

func (c *UserController) VerifyEmail(ctx *gin.Context){
	email := ctx.Query("email")
	token := ctx.Query("token")

	code, err := c.UserUseCase.VerifyEmail(email, token)

	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}


	ctx.JSON(code, gin.H{"message" : "email verified successfully"})
}

func (c *UserController) Login(ctx *gin.Context){
	credential := Domain.Credential{}
	err := ctx.ShouldBindJSON(&credential)

	if err != nil{
		ctx.JSON(400, gin.H{"error" : "invalid request payload"})
		return
	}

	var token, refresher string
	var code int

	if credential.Email != "" &&credential.Password != ""{
		token, refresher, code, err = c.UserUseCase.LoginByEmail(credential.Email, credential.Password)
	}else{
		ctx.JSON(code, gin.H{"error" : "email and password are required"})
		return
	}
	

	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{
		"token" : token,
		"refresher" : refresher,
	})
}

func (c *UserController) GetHeader(ctx *gin.Context) (map[string]interface{}, error){
	authHeader := ctx.GetHeader("Authorization")
	parts := strings.Split(authHeader, " ")
	claims, err := c.TokenService.ValidateToken(parts[1])
	return claims, err
}






