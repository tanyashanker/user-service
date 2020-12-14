package common

import "cache-service/models"

//Mapping the response of the operation
func ResponseMapper(data interface{}, code int, message string) models.Response {
	var response models.Response
	response = models.Response{
		Data:    data,
		Code:    code,
		Message: message,
	}
	return response

}
