package Repository

import (
	"DSAShare/Domain"
	"DSAShare/UseCase"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LectureRepository struct{
	DbCtx context.Context
	Collection *mongo.Collection
}

func NewLectureRepository(dbCtx context.Context, collection *mongo.Collection) UseCase.ILectureRepository{
	return &LectureRepository{
		DbCtx: dbCtx,
		Collection: collection,
	}
}

func (lr *LectureRepository) AddLecture(lecture *Domain.Lecture) error{
	_, err := lr.Collection.InsertOne(lr.DbCtx, lecture)
	return err
}

func (lr *LectureRepository) GetLectureByID(id string) (*Domain.Lecture, error){
	var lecture Domain.Lecture
	filter := bson.M{"_id" : id}

	err := lr.Collection.FindOne(lr.DbCtx, filter).Decode(&lecture)
	return &lecture, err
}

func(lr *LectureRepository)	GetAllLectures() (*[]Domain.Lecture, error){
	var lectures []Domain.Lecture

	cursor, err := lr.Collection.Find(lr.DbCtx, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments{
			return nil, nil
		}
		return nil, err
	}

	defer cursor.Close(lr.DbCtx)

	err = cursor.All(lr.DbCtx, &lectures)
	return &lectures, err
}