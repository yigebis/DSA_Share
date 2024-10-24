package main

import (
	"DSAShare/Delivery/Controller"
	"DSAShare/Delivery/Router"
	"DSAShare/Error"
	"DSAShare/Infrastructure"
	"DSAShare/Repository"
	"DSAShare/UseCase"

	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
	if err := godotenv.Load(); err != nil{
		log.Fatalf("error loading .env file")
	}

	// setting up the usecase

	//user usecase
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	uri := "mongodb+srv://" + username + ":" + password + "@cluster0.isgee.mongodb.net/"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to db!")
	user_collection := client.Database("DSAShare").Collection("users")
	lecture_collection := client.Database("DSAShare").Collection("lectures")
	token_collection := client.Database("DSAShare").Collection("refreshers")
	topic_collection := client.Database("DSAShare").Collection("topics")

	user_context := context.TODO()
	lecture_context := context.TODO()
	token_context := context.TODO()
	topic_context := context.TODO()

	ur := Repository.NewUserRepository(user_context, user_collection)
	lr := Repository.NewLectureRepository(lecture_context, lecture_collection)
	topr := Repository.NewTopicRepository(topic_collection, topic_context)
	tr := Repository.NewTokenRepository(token_context, token_collection)
	fr := Repository.NewFileRepository()

	jwtSecret := os.Getenv("JWT_SECRET")
	ps := Infrastructure.NewPasswordService()
	ts := Infrastructure.NewTokenService(jwtSecret)
	ms := Infrastructure.NewMailService(os.Getenv("SENDER_EMAIL"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("FROM"))
	es := Error.NewErrorService()
	
	ex := os.Getenv("EMAIL_EXPIRY")
	tx := os.Getenv("TOKEN_EXPIRY")
	rx := os.Getenv("REFRESHER_EXPIRY")

	uuc := UseCase.NewUserUseCase(ur, ps, tr, ts, ms, es, ex, tx, rx)
	luc := UseCase.NewLectureUseCase(lr, topr, ur, es)
	fuc := UseCase.NewFileUseCase(fr, es)
	uploadDir := os.Getenv("UPLOAD_DIR")

	// setting up the controllers
	user_controller := Controller.NewUserController(uuc, ts)
	lecture_controller := Controller.NewLectureController(luc)
	file_controller := Controller.NewFileController(uploadDir, fuc)

	// setting up the router
	router := Router.NewRouter(user_controller, lecture_controller, file_controller, jwtSecret)
	router.Run()
}