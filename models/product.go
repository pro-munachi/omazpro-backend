package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Name		*string					`json:"name" validate:"required,min=1,max=100"`
	Price		*int					`json:"price" validate:"required,min=1"`
	Countinstock		*int					`json:"countinstock" validate:"required,min=0"`
	Outofstock		*bool					`json:"outofstock"`
	Description		*string					`json:"description" `
	Category		*string					`json:"category" `
	Image		*string					`json:"image" `
	Createdat		time.Time				`json:"createdat"`
	Updatedat		time.Time				`json:"updatedat"`
	Productid		string					`json:"productid" `
}