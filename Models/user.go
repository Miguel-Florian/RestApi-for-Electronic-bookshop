package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,onitempty" bson:"_id,onitempty"`
	Username  string             `json:"username,onitempty" bson:"username,onitempty"`
	FirstName string             `json:"firstname,onitempty" bson:"firstname,onitempty"`
	LastName  string             `json:"lastname,onitempty,unique" bson:"lastname,onitempty"`
	Email     string             `json:"email,onitempty" bson:"email,onitempty,unique"`
	Password  string             `json:"password" bson:"password,onitempty"`
	//HashPassword []byte        `json:"hashpassword,omitempty "`
}
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
