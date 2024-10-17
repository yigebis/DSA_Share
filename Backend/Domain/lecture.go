package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lecture struct{
	ID 						primitive.ObjectID		`json:"id" bson:"_id"`
	Title					string					`json:"title" bson:"title" validate:"required"`
	Author					string					`json:"author" bson:"author" validate:"required"`
	Topics					[]string				`json:"topics" bson:"topics"`
	Content					string					`json:"content" bson:"content" validate:"required,min=50"`
	Votes					int64					`json:"votes" bson:"votes"`
	Views					int64					`json:"views" bson:"views"`
	CreationDate			time.Time				`json:"creation_date" bson:"creation_date"`
	LastModifiedDate		time.Time				`json:"last_modified_date" bson:"last_modified_date"`
}