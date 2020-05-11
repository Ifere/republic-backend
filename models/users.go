package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID 		primitive.ObjectID		`json:"user_id" bson:"user_id"`
	FirstName	string				`json:"first_name" bson:"first_name"`
	Phone		string				`json:"phone" bson:"phone"`
	ZoneID		primitive.ObjectID	`json:"zone_id" bson:"zone_id"`
}


func (u *UserModel) SetUserID () {
	u.ID = primitive.NewObjectID()
}
