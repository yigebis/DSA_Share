package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserName		   string			  `json:"user_name" bson:"user_name" validate:"required"`
	FirstName          string             `json:"first_name" bson:"first_name" validate:"required,min=1,max=50"`
	LastName           string             `json:"last_name" bson:"last_name" validate:"required,min=1,max=50"`
	Email              string             `json:"email" bson:"email" validate:"required,email"`
	Password           string             `json:"password" bson:"password" validate:"required"`
	ProfilePhoto       string             `json:"profile_photo" bson:"profile_photo"`
	LectureCount	   int 			  	  `json:"lecture_count" bson:"lecture_count"`
	RegistrationDate   time.Time          `json:"registration_date" bson:"registration_date"`
	Verified           bool               `json:"verified" bson:"verified"`
}

type Credential struct{
	Identifier string `json:"identifier" bson:"identifier"`
	Password string `json:"password" bson:"password"`
}

type ChangeCredential struct{
	Email string `json:"email" bson:"email"`
	OldPassword string `json:"old_password" bson:"old_password"`
	NewPassword string `json:"new_password" bson:"new_password"`
}