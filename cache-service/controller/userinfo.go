package controller

import (
	"cache-service/common"
	"cache-service/models"
	"cache-service/services"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserInfoHandler struct{}

var UserInfo UserInfoHandler

//GetUserInfo to get the info of a single user
func (u *UserInfoHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["userID"]

	log.Println(ID)
	w.Header().Set("Content-Type", "application/json; charset= UTF-8")
	user, err := services.GetUserInfo(ID)
	if err != nil {
		response := common.ResponseMapper(models.User{}, http.StatusBadRequest, "Bad Request")
		json.NewEncoder(w).Encode(response)
		return
	}
	if user == nil {
		response := common.ResponseMapper(models.User{}, http.StatusNotFound, "User Not Found")
		json.NewEncoder(w).Encode(response)
		return
	}
	response := common.ResponseMapper(user, http.StatusOK, "User details fetched successfully")
	json.NewEncoder(w).Encode(response)
}

//AddUser to add user info
func (u *UserInfoHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	json.Unmarshal(body, &user)

	w.Header().Set("Content-Type", "application/json; charset= UTF-8")
	err = services.AddUser(user)
	if err != nil {
		response := common.ResponseMapper("", http.StatusBadRequest, "Bad Request")
		json.NewEncoder(w).Encode(response)
		return
	}
	response := common.ResponseMapper("", http.StatusCreated, "User Created Successfully")
	json.NewEncoder(w).Encode(response)
}

//GetAllUsers to get all the users
func (u *UserInfoHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset= UTF-8")

	queryParam := r.URL.Query()

	size, _ := strconv.Atoi(queryParam.Get("pageSize"))
	pageNum, _ := strconv.Atoi(queryParam.Get("pageNumber"))
	log.Println(size, pageNum)

	users, err := services.GetAllUsers(size, pageNum)
	if err != nil {
		response := common.ResponseMapper([]models.User{}, http.StatusBadRequest, "Bad Request")
		json.NewEncoder(w).Encode(response)
		return
	}
	if len(users) == 0 {
		response := common.ResponseMapper([]models.User{}, http.StatusNotFound, "No users to display")
		json.NewEncoder(w).Encode(response)
		return
	}
	response := common.ResponseMapper(users, http.StatusOK, "Fetched users successfully")
	json.NewEncoder(w).Encode(response)
}

//UpdateUser to update the user
func (u *UserInfoHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	json.Unmarshal(body, &user)

	vars := mux.Vars(r)
	userID := vars["userID"]

	w.Header().Set("Content-Type", "application/json; charset= UTF-8")
	err = services.UpdateUser(user, userID)
	if err != nil {
		response := common.ResponseMapper("", http.StatusBadRequest, "Bad Request")
		json.NewEncoder(w).Encode(response)
		return
	}
	response := common.ResponseMapper("", http.StatusOK, "User Created Successfully")
	json.NewEncoder(w).Encode(response)
}
