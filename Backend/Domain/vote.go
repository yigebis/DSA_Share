package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vote struct{
	ID 					primitive.ObjectID
	VoterUserName 		string
	VotedLecture		string
	VoteTime			time.Time
	VoteType			int
}