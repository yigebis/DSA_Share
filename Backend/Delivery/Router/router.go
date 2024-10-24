package Router

import (
	"DSAShare/Delivery/Controller"

	"github.com/gin-gonic/gin"
)

type Router struct{
	UserController *Controller.UserController
	LectureController *Controller.LectureController
	FileController *Controller.FileController
	JWTSigner string
}

func NewRouter(uc *Controller.UserController, lc *Controller.LectureController, fc *Controller.FileController, jwtSigner string) *Router{
	return &Router{
		UserController: uc,
		LectureController: lc,
		FileController: fc,
		JWTSigner: jwtSigner,
	}
}

func (r *Router) Run(){
	router := gin.Default()


	router.POST("/register", r.UserController.Register)
	router.POST("/login/email", r.UserController.LoginByEmail)
	router.POST("/login/user_name", r.UserController.LoginByUserName)
	router.GET("/verify", r.UserController.VerifyEmail)

	router.POST("/lecture/add", r.LectureController.AddLecture)
	router.GET("/lecture/all", r.LectureController.GetAllLectures)
	router.GET("/lectures/:user_name", r.LectureController.GetLecturesOf)
	router.GET("/lecture/:id", r.LectureController.GetLecture)
	router.PUT("/lecture/edit/:id", r.LectureController.EditLecture)
	router.POST("/lecture/delete/:id", r.LectureController.DeleteLecture)
	router.PUT("/lecture/topic/add", r.LectureController.AddTopic)
	router.PUT("/lecture/topic/remove", r.LectureController.RemoveTopic)
	router.GET("/lecture/search", r.LectureController.SearchLectures)

	router.POST("/upload", r.FileController.UploadFile)

	
	router.Run()
}