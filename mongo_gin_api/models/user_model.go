package models

// Import the required dependencies.
// Create a User struct with required properties. We added omitempty and validate:"required" to the struct tag to tell Gin-gonic to ignore empty fields and make the field required, respectively.
// Create a User Endpoint
// we need a model to represent our application data

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Age 	 int             `json:"location,omitempty" validate:"required"`
	Designation    string             `json:"title,omitempty" validate:"required"`
}
