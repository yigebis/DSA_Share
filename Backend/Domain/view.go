package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type View struct{
	ID 					primitive.ObjectID
	ViewerUserName 		string
	ViewedLecture		string
	ViewTime			time.Time
}