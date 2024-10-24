package UseCase

import (
	"mime/multipart"
	"path/filepath"
)

type FileUseCase struct{
	FileRepository IFileRepository
	ErrorService IErrorService
}

func NewFileUseCase(fr IFileRepository, es IErrorService) IFileUseCase{
	return &FileUseCase{
		FileRepository: fr,
		ErrorService: es,
	}
}

func (fuc *FileUseCase) UploadFile(fileHeader *multipart.FileHeader, uploadDir string) (string, int, error) {
	fileName := fileHeader.Filename
	fileUrl := filepath.Join(uploadDir, fileName)

	err := fuc.FileRepository.SaveFile(fileHeader, fileUrl)
	if err != nil {
		code, err := fuc.ErrorService.InternalServer()
		return "", code, err
	}

	code, err := fuc.ErrorService.NoError()
	return fileUrl, code, err
}