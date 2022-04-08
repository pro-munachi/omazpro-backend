package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gallery struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Img				*string					`json:"img" validate:"required,min=1"`
	Category		*string					`json:"category" validate:"required,min=1"`
	Createdat		time.Time				`json:"created_at"`
	Updatedat		time.Time				`json:"updated_at"`
	Picid			string					`json:"user_id"`
}
