package common

import (
	"cache-service/models"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/labstack/gommon/log"
)

var Conf *models.DBConfig

func logFileName(filePath string) {
	log.Info(fmt.Sprintf("Reading file %s", filePath))
}

func ReadSecret(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func UnMarshellSecret(respStruct interface{}, contents interface{}) {
	err := json.Unmarshal(contents.([]byte), respStruct)
	if err != nil {
		log.Fatal("Error while json unmarshalling the response. Error: " + err.Error())
	}
}

//Getting secret values from specific files
func GetSecretValues(filePath string) *models.Secret {
	logFileName(filePath)

	contents, err := ReadSecret(filePath)
	if err != nil {
		log.Error("Failed to read secret file. Error: " + err.Error())
	}
	SecretObject := &models.Secret{}
	UnMarshellSecret(SecretObject, contents)

	return SecretObject
}

func GetConfigObj() *models.DBConfig {
	if Conf == nil {
		Conf = &models.DBConfig{}
		Conf.MongoSecrets = &models.Secret{}
		Conf.RedisSecrets = &models.Secret{}
	}
	return Conf
}
