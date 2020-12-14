package services

import (
	"cache-service/models"
	"cache-service/repository"
	"cache-service/utils"
	"log"
)

//GetUserInfo to get user details based on userID
func GetUserInfo(userID string) (*models.User, error) {
	user, err := repository.DBClient.GetUserInfo(userID)
	if err != nil {
		log.Println("Error occurred while fetching the user details : ", err.Error())
		return user, err
	}
	return user, nil
}

//AddUser to add user details
func AddUser(user models.User) error {

	user.UserID = utils.GenerateUUID()
	err := repository.DBClient.AddUser(user)
	if err != nil {
		log.Println("Error occurred while adding the user details : ", err.Error())
		return err
	}
	return nil
}

//GetAllUsers to get all the users based on size and page number
func GetAllUsers(size, pageNum int) ([]models.User, error) {
	if pageNum <= 1 {
		pageNum = 1
	}
	if size == 0 {
		size = 1
	}

	skip := (pageNum - 1) * size
	users, err := repository.DBClient.GetAllUsers(int64(size), int64(skip))
	if err != nil {
		log.Println("Error occurred while fetching the user details : ", err.Error())
		return users, err
	}
	return users, nil
}

//UpdateUser to update the user details
func UpdateUser(user models.User, userID string) error {

	user.UserID = userID
	err := repository.DBClient.UpdateUser(user, userID)
	if err != nil {
		log.Println("Error occurred while updating the user details : ", err.Error())
		return err
	}
	return nil
}
