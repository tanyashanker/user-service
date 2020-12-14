package repository

import (
	"cache-service/common"
	"cache-service/models"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DBInteractions interface contains database methods
type User interface {
	GetUserInfo(userID string) (*models.User, error)
	AddUser(user models.User) error
	GetAllUsers(size, pageNum int) ([]models.User, error)
	UpdateUser(user models.User, userID string) error

	CacheUserData(user models.User, count int64)
	GetCachedUserData(size, skip int64) ([]models.User, error)
}

//GetUserInfo to get user details based on userID
func (u *DBRepo) GetUserInfo(userID string) (*models.User, error) {
	var user *models.User

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := u.MongoClient.Database(common.Conf.MongoSecrets.DBName).Collection("user")

	cursor := collection.FindOne(ctx, bson.M{"userId": userID})
	if err := cursor.Err(); err != nil {
		log.Println("DB Error :", err)
		return user, errors.New("DB error")
	}
	cursor.Decode(&user)

	return user, nil
}

//AddUser adds user to mongodb and redis cache
func (u *DBRepo) AddUser(user models.User) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := u.MongoClient.Database(common.Conf.MongoSecrets.DBName).Collection("user")

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Mongodb Error :", err)
		return err
	}
	log.Println("User added in db")

	err = u.CacheUserData(user)
	if err != nil {
		log.Println("Redis-cache Error :", err)
		return err
	}
	log.Println("User details added in redis cache")
	return nil
}

//GetAllUsers to get all the users from redis and mongo db
func (u *DBRepo) GetAllUsers(size, skip int64) ([]models.User, error) {
	var users []models.User
	var err error
	users, err = u.GetCachedUserData(size, skip)
	if err != nil || len(users) == 0 {
		log.Println("Getting user details from cache")
		log.Println("Error while fetching user details from cache. Fetching details from db...")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		collection := u.MongoClient.Database(common.Conf.MongoSecrets.DBName).Collection("user")

		opts := options.FindOptions{
			Skip:  &skip,
			Limit: &size,
		}
		cursor, err := collection.Find(ctx, bson.M{}, &opts)
		if err != nil {
			log.Println("DB Error :", err)
			return users, errors.New("DB error")
		}
		err = cursor.All(ctx, &users)
		if err != nil {
			log.Println("Failed to decode cursor results :", err)
			return users, errors.New("Failed to decode cursor results")
		}
	}
	return users, nil
}

//UpdateUser to update the user details
func (u *DBRepo) UpdateUser(user models.User, userID string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := u.MongoClient.Database(common.Conf.MongoSecrets.DBName).Collection("user")

	_, err := collection.UpdateOne(ctx, bson.M{"userId": userID}, bson.M{"$set": user})
	if err != nil {
		log.Println("DB Error :", err)
		return errors.New("DB error")
	}

	return nil
}
