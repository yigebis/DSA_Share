package Repository

import (
	"DSAShare/UseCase"

	"mime/multipart"
	"os"
)

type FileRepository struct{
}

func NewFileRepository() UseCase.IFileRepository{
	return &FileRepository{}
}

func (fr *FileRepository) SaveFile(fileHeader *multipart.FileHeader, fileUrl string) error{
	// open input file
	inFile, err := fileHeader.Open()
	if err != nil {
		return err
	}

	//defer closing of input file
	defer inFile.Close()

	//create the output file
	outFile, err := os.Create(fileUrl)
	if err != nil {
		return err
	}

	//defer closing of output file
	defer outFile.Close()

	//copy from input file to output file
	_, err = inFile.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = outFile.ReadFrom(inFile)
	return err
}