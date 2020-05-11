package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ZoneModel struct {
	ID        	primitive.ObjectID 	`json:"zone_id" bson:"zone_id"`
	Name      	string             	`json:"name" bson:"name"`
	Type       	string             	`json:"type" bson:"type"`
	MemberID  	primitive.ObjectID 	`json:"member_id" bson:"member_id"`
	Bio			string				`json:"bio" bson:"bio"`
	RulingParty string             	`json:"ruling_party" bson:"ruling_party"`
	State     	string             	`json:"state" bson:"state"`
	Population  int					`json:"population" bson:"population"`

}

func (z *ZoneModel) SetZoneId() {
	z.ID = primitive.NewObjectID()
}
