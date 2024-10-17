package UseCase

import (
	"DSAShare/Domain"
)

type ILectureUseCase interface{
	AddLecture(lecture *Domain.Lecture) (int, error)
	GetLectureByID(id string) (*Domain.Lecture, int, error)
	GetAllLectures() (*[]Domain.Lecture, int, error)
}

type ILectureRepository interface{
	AddLecture(lecture *Domain.Lecture) error
	GetLectureByID(id string) (*Domain.Lecture, error)
	GetAllLectures() (*[]Domain.Lecture, error)
}

type ITopicRepository interface{
	UpsertTopics(topics []string) error
	DecrementTopicCount(topic string, by int) error
}