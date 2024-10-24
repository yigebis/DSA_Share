package Domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Topic struct{
	ID primitive.ObjectID 		`json:"id" bson:"_id"`
	Name string					`json:"name" bson:"name"`
	TagCount int64				`json:"tag_count" bson:"tag_count"`
}