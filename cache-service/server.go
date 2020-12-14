package main

import (
	"cache-service/common"
	"cache-service/controller"
	"cache-service/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/pdrum/swagger-automation/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	initConfig()

	router := mux.NewRouter()

	router.HandleFunc("/ping", controller.Ping.Ping).Methods("GET")
	router.HandleFunc("/user/{userID}", controller.UserInfo.GetUserInfo).Methods("GET")
	router.HandleFunc("/users", controller.UserInfo.GetAllUsers).Methods("GET")
	router.HandleFunc("/user", controller.UserInfo.AddUser).Methods("POST")
	router.HandleFunc("/user/{userID}", controller.UserInfo.UpdateUser).Methods("PUT")

	log.Println("Starting the application at port 8080")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", router)
}

func initConfig() {
	//getting config for db and cache
	common.Conf = common.GetConfigObj()
	repository.DBClient = new(repository.DBRepo)
	InitMongo()
	InitRedis()
}

//Preparing for connecting to mongo
func InitMongo() {
	common.Conf.MongoSecrets = common.GetSecretValues(common.MONGO_SECRETS)
	err := repository.DBClient.MongoConnect(common.Conf.MongoSecrets)
	if err != nil {
		log.Fatal("Unable to connect to the db...")
	}

}

//Preparing for connecting to Redis
func InitRedis() {
	var err error
	common.Conf.RedisSecrets = common.GetSecretValues(common.REDIS_SECRETS)
	repository.DBClient.RedisClient, err = repository.DBClient.RedisConnect(common.Conf.RedisSecrets)
	if err != nil {
		log.Fatal("Unable to connect to the db...")
	}

}
