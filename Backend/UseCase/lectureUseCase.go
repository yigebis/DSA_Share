package UseCase

import (
	"DSAShare/Domain"

	"time"
)

type LectureUseCase struct{
	LectureRepository ILectureRepository
	TopicRepository ITopicRepository
	UserRepository IUserRepository
	ErrorService IErrorService
}

func NewLectureUseCase(lr ILectureRepository, topr ITopicRepository, ur IUserRepository, es IErrorService) ILectureUseCase{
	return &LectureUseCase{
		LectureRepository: lr,
		TopicRepository: topr,
		UserRepository: ur,
		ErrorService: es,
	}
}

func (luc *LectureUseCase) AddLecture(lecture *Domain.Lecture) (int, error){
	// views and votes set to zero
	lecture.Views = 0
	lecture.Votes = 0
	// creation and last modified date set to current time
	lecture.CreationDate = time.Now()
	lecture.LastModifiedDate = lecture.CreationDate
	// add the lecture to the database
	if err := luc.LectureRepository.AddLecture(lecture); err != nil{
		return luc.ErrorService.InternalServer()
	}

	// add or update the topics in the lectures
	if err := luc.TopicRepository.UpsertTopics(lecture.Topics); err != nil{
		return luc.ErrorService.InternalServer()
	}
	// update the lecture count of the author
	if err := luc.UserRepository.IncrementLectureCount(lecture.Author); err != nil{
		return luc.ErrorService.InternalServer()
	}

	return luc.ErrorService.NoError()
}

func (luc *LectureUseCase) GetAllLectures() (*[]Domain.Lecture, int, error){
	// read the lectures from the database
	lectures, err := luc.LectureRepository.GetAllLectures()
	if err != nil {
		code, err := luc.ErrorService.InternalServer()
		return nil, code, err
	}

	code, err := luc.ErrorService.NoError()
	return lectures, code, err
}

func (luc *LectureUseCase) GetLectureByID(id string) (*Domain.Lecture, int, error){
	// read the lecture from the database
	lecture, err := luc.LectureRepository.GetLectureByID(id)
	if err != nil {
		code, err := luc.ErrorService.LectureNotFound()
		return nil, code, err
	}

	code, err := luc.ErrorService.NoError()
	return lecture, code, err
}