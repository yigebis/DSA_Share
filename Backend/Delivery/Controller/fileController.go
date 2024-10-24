package Controller

import (
	"DSAShare/UseCase"

	"net/http"

	"github.com/gin-gonic/gin"
)

type FileController struct{
	UploadDir		string
	FileUseCase		UseCase.IFileUseCase
}

func NewFileController(uploadDir string, fuc UseCase.IFileUseCase) *FileController{
	return &FileController{
		UploadDir : uploadDir,
		FileUseCase : fuc,
	}
}

func (fc *FileController) UploadFile(ctx *gin.Context){
	fileHeader, err := ctx.FormFile("file")

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : "failed uploading file"})
		return
	}

	fileUrl, code, err := fc.FileUseCase.UploadFile(fileHeader, fc.UploadDir)
	if err != nil{
		ctx.JSON(code, gin.H{"error" : err.Error()})
		return
	}

	ctx.JSON(code, fileUrl)
}