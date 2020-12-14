package repository

import (
	"cache-service/common"
	"cache-service/models"
	"encoding/json"
	"log"
	"strings"
)

func (u *DBRepo) CacheUserData(user models.User) error {
	var key strings.Builder

	_, err := key.WriteString(common.REDIS_USER_KEY)
	if err != nil {
		log.Println("Error in creating key for cache :", err)
		return err
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Println("Error in marshaling user details to json")
		return err
	}

	err = u.RedisClient.HSet(key.String(), user.UserID, string(userJSON)).Err()
	if err != nil {
		log.Println("Error in caching the user details", err)
		return err
	}

	return nil

}

func (u *DBRepo) GetCachedUserData(size, skip int64) ([]models.User, error) {
	var key strings.Builder
	var user models.User
	var users []models.User
	_, err := key.WriteString(common.REDIS_USER_KEY)
	if err != nil {
		log.Println("Error in creating key for cache :", err)
		return users, err
	}

	result, err := u.RedisClient.HGetAll(key.String()).Result()
	if err != nil {
		log.Println("Error in getting the user details from cache", err)
		return users, err
	}

	for _, val := range result {
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			log.Println("Error in unmarshling the user details from cache", err)
			return users, err
		}
		users = append(users, user)
	}

	log.Println(users)

	return users, nil

}
