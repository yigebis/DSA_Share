package Controller

import (
	"DSAShare/Domain"
	"DSAShare/UseCase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (lc *LectureController) EditLecture(ctx *gin.Context){
	id := ctx.Param("id")
	var lecture Domain.Lecture

	if err := ctx.ShouldBindJSON(&lecture); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "invalid request payload"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}

	lecture.ID = objID

	code, err := lc.LectureUseCase.EditLecture(&lecture)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" : "update successful"})
}

func (lc *LectureController) DeleteLecture(ctx *gin.Context){
	id := ctx.Param("id")

	code, err := lc.LectureUseCase.DeleteLecture(id)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" : "delete successful"})
}

func (lc *LectureController) AddTopic(ctx *gin.Context){
	topic := ctx.Query("topic")
	lectureID := ctx.Query("lecture_id")

	code, err := lc.LectureUseCase.AddTopic(topic, lectureID)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" : "topic added successfully"})
}

func (lc *LectureController) RemoveTopic(ctx *gin.Context){
	topic := ctx.Query("topic")
	lectureID := ctx.Query("lecture_id")

	code, err := lc.LectureUseCase.RemoveTopic(topic, lectureID)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, gin.H{"message" : "topic removed successfully"})
}