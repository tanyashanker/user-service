package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:route POST /user user-tag
// responses:
//   200: Response

// This text will appear as description of your response body.
// swagger:response Response
type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID      string             `json:"userId"  bson:"userId"`
	UserName    string             `json:"userName" bson:"userName"`
	Email       string             `json:"email" bson:"email"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber"`
	Address     Address            `json:"address" bson:"address"`
}

type Address struct {
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
	Pincode int64  `json:"pincode" bson:"pincode"`
}

// swagger:parameters Response
type Response struct {
	// in:body
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}
