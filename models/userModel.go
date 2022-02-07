package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:_id`
	FirstName     string             `json:"first_name" validate:"required, min=2, max=100"`
	LastName      string             `json:"last_name" validate:"required, min=2, max=100"`
	Password      string             `json:"password" validate:"required,min=6, max=100"`
	Email         string             `json:"email" validate:"email, required"`
	Token         string             `json:"token"`
	User_type     string             `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Refresh_token string             `json:"refresh_token"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	userId        string             `json:user_Id`
}
