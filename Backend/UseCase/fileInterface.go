package UseCase

import (
	"mime/multipart"
)

type IFileUseCase interface{
	UploadFile(fileHeader *multipart.FileHeader, uploadDir string) (string, int, error)
}

type IFileRepository interface{
	SaveFile(fileHeader *multipart.FileHeader, fileUrl string) error
}