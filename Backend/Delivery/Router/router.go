package Router

import (
	"DSAShare/Delivery/Controller"

	"github.com/gin-gonic/gin"
)

type Router struct{
	UserController *Controller.UserController
	JWTSigner string
}

func NewRouter(uc *Controller.UserController, jwtSigner string) *Router{
	return &Router{
		UserController: uc,
		JWTSigner: jwtSigner,
	}
}

func (r *Router) Run(){
	router := gin.Default()


	router.POST("/register", r.UserController.Register)
	router.POST("/login", r.UserController.Login)
	router.GET("/verify", r.UserController.VerifyEmail)

	router.Run()
}