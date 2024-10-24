package Repository

import (
	"DSAShare/Domain"
	"DSAShare/UseCase"
	"fmt"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DbCtx      context.Context
	Collection *mongo.Collection
}

func NewUserRepository(dbCtx context.Context, collection *mongo.Collection) UseCase.IUserRepository {
	return &UserRepository{
		DbCtx:      dbCtx,
		Collection: collection,
	}
}

func (ur *UserRepository) CreateUser(user *Domain.User) error {
	_, err := ur.Collection.InsertOne(ur.DbCtx, user)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (ur *UserRepository) GetUserByEmail(email string) (*Domain.User, error) {
	var user Domain.User
	filter := bson.M{"email": email}
	err := ur.Collection.FindOne(ur.DbCtx, filter).Decode(&user)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByUserName(user_name string) (*Domain.User, error){
	var user Domain.User

	filter := bson.M{"user_name" : user_name}

	err := ur.Collection.FindOne(ur.DbCtx, filter).Decode(&user)
	return &user, err
}

func (ur *UserRepository) VerifyUser(user *Domain.User) error {
	filter := bson.M{"email": user.Email}
	update := bson.M{"$set": bson.M{
		"verified": true,
	}}

	_, err := ur.Collection.UpdateOne(ur.DbCtx, filter, update)
	fmt.Println(err)
	return err
}

func (ur *UserRepository) UpdateUser(user *Domain.User) error {
	filter := bson.M{"email": user.Email}
	var update = bson.M{}

	if user.FirstName != "" {
		update["first_name"] = user.FirstName
	}
	if user.LastName != "" {
		update["last_name"] = user.LastName
	}
	if user.ProfilePhoto != "" {
		update["profile_photo"] = user.ProfilePhoto
	}
	if user.Password != "" {
		update["password"] = user.Password
	}
	if user.LectureCount != 0 {
		update["lecture_count"] = user.LectureCount
	}

	_, err := ur.Collection.UpdateOne(ur.DbCtx, filter, bson.M{"$set": update})
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (ur *UserRepository) GetUserByVerificationToken(token string) (*Domain.User, error) {
	return nil, nil
}

func (ur *UserRepository) GetUserCount() (int64, error) {
	return 0, nil
}

func (ur *UserRepository) UpdatePasswordByEmail(email string, newPassword string) error {
	return nil
}

func (ur *UserRepository) StoreResetToken(email string, resetToken string) error {
	return nil
}

func (ur *UserRepository) InvalidateResetToken(email string) error {
	return nil
}

func (ur *UserRepository) GetResetTokenByEmail(email string) (string, error) {
	return "", nil
}

func (ur *UserRepository) DeleteUser(email string) error {
	filter := bson.M{"email": email}

	_, err := ur.Collection.DeleteOne(ur.DbCtx, filter)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func (ur *UserRepository) IncrementLectureCount(user_name string) error{
	filter := bson.M{"user_name" : user_name}
	update := bson.M{
		"$inc" : bson.M{"lecture_count" : 1},
	}

	_, err := ur.Collection.UpdateOne(ur.DbCtx, filter, update)
	return err
}