package Repository

import (
	"DSAShare/Domain"
	"DSAShare/UseCase"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (lr *LectureRepository) GetLecturesOf(userName string) (*[]Domain.Lecture, error){
	var lectures []Domain.Lecture

	filter := bson.M{"author" : userName}

	cursor, err := lr.Collection.Find(lr.DbCtx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(lr.DbCtx)

	err = cursor.All(lr.DbCtx, &lectures)
	if err != nil {
		return nil, err
	}

	return &lectures, nil
}

func (lr *LectureRepository) EditLecture(lecture *Domain.Lecture) error{
	filter := bson.M{"_id" : lecture.ID}
	var update = bson.M{}

	update["last_modified_date"] = lecture.LastModifiedDate
	if len(lecture.Content) > 0{
		update["content"] = lecture.Content
	}
	if lecture.Title != ""{
		update["title"] = lecture.Title
	}
	
	_, err := lr.Collection.UpdateOne(lr.DbCtx, filter, bson.M{"$set" : update})
	return err
}

func (lr *LectureRepository) DeleteLecture(id string) error{
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return err
	}

	filter := bson.M{"_id" : objID}
	_, err = lr.Collection.DeleteOne(lr.DbCtx, filter)
	return err
}

func (lr *LectureRepository) AddTopic(topic, lectureID string) error{
	objID, err := primitive.ObjectIDFromHex(lectureID)
	if err != nil{
		return err
	}

	filter := bson.M{"_id" : objID}
	update := bson.M{"$addToSet" : bson.M{"topics" : topic}}

	_, err = lr.Collection.UpdateOne(lr.DbCtx, filter, update)
	return err
}

func (lr *LectureRepository) RemoveTopic(topic, lectureID string) error{
	objID, err := primitive.ObjectIDFromHex(topic)
	if err != nil {
		return err
	}

	filter := bson.M{"_id" : objID}
	update := bson.M{"$pull" : bson.M{"topics" : topic}}

	_, err = lr.Collection.UpdateOne(lr.DbCtx, filter, update)
	return err
}

func (lr *LectureRepository) SearchLectures(query *map[string]interface{}, lastID string, pageSize int) (*[]Domain.Lecture, error){
	var lectures []Domain.Lecture
	filter := bson.M{}
		
	for key, value := range *query{
		if key != "topics"{
			filter[key] = value
		}
	}

	if val, exists := (*query)["topics"]; exists{
		filter["topics"] = bson.M{"$all" : val}
	}

	if lastID != ""{
		objectID, err := primitive.ObjectIDFromHex(lastID)
        if err != nil {
            // fmt.Println("Invalid ObjectID:", err)
            return nil, err
        }
		filter["_id"] = bson.M{"$gt" : objectID}
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))

	cursor, err := lr.Collection.Find(lr.DbCtx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(lr.DbCtx)

	err = cursor.All(lr.DbCtx, &lectures)

	return &lectures, err
}