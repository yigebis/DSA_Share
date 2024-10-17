package Controller

import (
	"DSAShare/Domain"
	"DSAShare/UseCase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LectureController struct{
	LectureUseCase UseCase.ILectureUseCase
	V *validator.Validate
}

func NewLectureController(luc UseCase.ILectureUseCase) *LectureController{
	return &LectureController{
		LectureUseCase: luc,
		V : validator.New(),
	}
}

func (lc *LectureController) AddLecture(ctx *gin.Context){
	var lecture Domain.Lecture

	if err := ctx.ShouldBindJSON(&lecture); err != nil{
		ctx.JSON(400, gin.H{"error" : "invalid request payload"})
		return
	}

	if err := lc.V.Struct(lecture); err != nil{
		ctx.JSON(400, gin.H{"error" : "invalid request payload"})
		return
	}

	code, err := lc.LectureUseCase.AddLecture(&lecture)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" : "lecture added successfully"})
}

func (lc *LectureController) GetAllLectures(ctx *gin.Context){

	lectures, code, err := lc.LectureUseCase.GetAllLectures()
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, lectures)
}

func (lc *LectureController) GetLecture(ctx *gin.Context){
	id := ctx.Param("id")
	lecture, code, err := lc.LectureUseCase.GetLectureByID(id)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, lecture)
}