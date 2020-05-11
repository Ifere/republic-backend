package models

import (
"go.mongodb.org/mongo-driver/bson/primitive"
"time"
)

type MemberModel struct {
	ID        	primitive.ObjectID 		`json:"member_id" bson:"member_id"`
	Name      	string             		`json:"name" bson:"name"`
	Age       	int                		`json:"age" bson:"age"`
	Gender    	string             		`json:"gender" bson:"gender"`
	Avatar    	string             		`json:"avatar" bson:"avatar"`
	Bio			string					`json:"bio" bson:"bio"`
	Party     	string             		`json:"party" bson:"party"`
	ZoneID		primitive.ObjectID	 	`json:"zone_id" bson:"zone_id"`
	Committee 	string             		`json:"committee" bson:"committee"`
	Position  	string					`json:"position" bson:"position"`
	StartDate	time.Time				`json:"start_date" bson:"start_date"`
	EndDate		time.Time				`json:"end_date" bson:"end_date"`
	Tenure		int						`json:"tenure" bson:"tenure"`
}

func (m *MemberModel) SetMemberId() {
	m.ID = primitive.NewObjectID()
}
