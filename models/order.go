package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	Name	*string		`json:"name"`
	Price	*int		`json:"price"`
	Qty		*int		`json:"qty"`
	Image	*string		`json:"image"`
	Product	string		`json:"product"`
}



type Order struct{
	ID				primitive.ObjectID		`bson:"_id"`
	OrderItem		[]OrderItem				`json:"orderitem"`
	TotalPrice		*int				`json:"totalprice"`
	IsPaid		*bool				`json:"ispaid"`
	PaymentMethod		*string				`json:"paymentmethod"`
	User		string				`json:"user"`
	Createdat		time.Time				`json:"createdat"`
	Updatedat		time.Time				`json:"updatedat"`
	Orderid		string					`json:"orderid" `
}