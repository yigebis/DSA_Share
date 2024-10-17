package Router

import (
	"DSAShare/Delivery/Controller"

	"github.com/gin-gonic/gin"
)

type Router struct{
	UserController *Controller.UserController
	LectureController *Controller.LectureController
	JWTSigner string
}

func NewRouter(uc *Controller.UserController, lc *Controller.LectureController, jwtSigner string) *Router{
	return &Router{
		UserController: uc,
		LectureController: lc,
		JWTSigner: jwtSigner,
	}
}

func (r *Router) Run(){
	router := gin.Default()


	router.POST("/register", r.UserController.Register)
	router.POST("/login", r.UserController.Login)
	router.GET("/verify", r.UserController.VerifyEmail)

	router.POST("/lecture/add", r.LectureController.AddLecture)
	router.GET("/lecture/all", r.LectureController.GetAllLectures)
	router.GET("/lecture/:id", r.LectureController.GetLecture)

	router.Run()
}