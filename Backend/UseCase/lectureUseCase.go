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

func (luc *LectureUseCase) GetLecturesOf(userName string) (*[]Domain.Lecture, int, error){
	lectures, err := luc.LectureRepository.GetLecturesOf(userName)
	if err != nil {
		code, err := luc.ErrorService.LectureNotFound()
		return nil, code, err
	}

	code, err :=luc.ErrorService.NoError()
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

func (luc *LectureUseCase) EditLecture(lecture *Domain.Lecture) (int, error){
	// make sure to change last modified date of the lecture
	lecture.LastModifiedDate = time.Now()
	
	err := luc.LectureRepository.EditLecture(lecture)
	if err != nil{
		return luc.ErrorService.InternalServer()
	}

	return luc.ErrorService.NoError()
}

func (luc *LectureUseCase) DeleteLecture(id string) (int, error){
	err := luc.LectureRepository.DeleteLecture(id)
	if err != nil {
		return luc.ErrorService.InternalServer()
	}

	return luc.ErrorService.NoError()
}

func (luc *LectureUseCase) AddTopic(topic, lectureID string) (int, error){
	err := luc.LectureRepository.AddTopic(topic, lectureID)
	if err != nil {
		return luc.ErrorService.InternalServer()
	}

	err = luc.TopicRepository.IncrementTopicCount(topic, 1)
	if err != nil {
		return luc.ErrorService.InternalServer()
	}

	return luc.ErrorService.NoError()
}

func (luc *LectureUseCase) RemoveTopic(topic, lectureID string) (int, error){
	err := luc.LectureRepository.RemoveTopic(topic, lectureID)
	if err != nil {
		return luc.ErrorService.InternalServer()
	}

	err = luc.TopicRepository.DecrementTopicCount(topic, 1)
	if err != nil {
		return luc.ErrorService.InternalServer()
	}

	return luc.ErrorService.NoError()
}

func (luc *LectureUseCase) SearchLectures(query *map[string]interface{}, lastID string) (*[]Domain.Lecture, int, error){
	// define business rules
	pageSize := 20

	lectures, err := luc.LectureRepository.SearchLectures(query, lastID, pageSize)
	if err != nil {
		code, err := luc.ErrorService.InternalServer()
		return nil, code, err
	}

	code, err := luc.ErrorService.NoError()
	return lectures, code, err
}