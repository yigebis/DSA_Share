package UseCase

import (
	"DSAShare/Domain"
)

type ILectureUseCase interface{
	AddLecture(lecture *Domain.Lecture) (int, error)
	GetLectureByID(id string) (*Domain.Lecture, int, error)
	GetAllLectures() (*[]Domain.Lecture, int, error)
	GetLecturesOf(userName string) (*[]Domain.Lecture, int, error)
	EditLecture(lecture *Domain.Lecture) (int, error)
	DeleteLecture(id string) (int, error)
	AddTopic(topic string, lectureID string) (int, error)
	RemoveTopic(topic string, lectureID string) (int, error)
	SearchLectures(query *map[string]interface{}, lastID string) (*[]Domain.Lecture, int, error)
}

type ILectureRepository interface{
	AddLecture(lecture *Domain.Lecture) error
	GetLectureByID(id string) (*Domain.Lecture, error)
	GetAllLectures() (*[]Domain.Lecture, error)
	GetLecturesOf(userName string) (*[]Domain.Lecture, error)
	EditLecture(lecture *Domain.Lecture) error
	DeleteLecture(id string) error
	AddTopic(topic string, lectureID string) error
	RemoveTopic(topic string, lectureID string) error
	SearchLectures(query *map[string]interface{}, lastID string, pageSize int) (*[]Domain.Lecture, error)
}

type ITopicRepository interface{
	UpsertTopics(topics []string) error
	IncrementTopicCount(topic string, by int) error
	DecrementTopicCount(topic string, by int) error
}