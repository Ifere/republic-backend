package memberRepo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	db "republic-backend/config/db"
	apperrors "republic-backend/config/errors"
	models "republic-backend/models"
)

type MembersMongoRepo struct {
	Client db.MongoDB
}

type MemberRepositoryI interface {
	CreateMember(m models.MemberModel) (models.MemberModel, error)
	GetMemberById(ID string) (models.MemberModel, error)
	FetchMembers(filter interface{}) ([] models.MemberModel, error)
	UpdateMemberDetails(ID string, update models.MemberModel) (models.MemberModel, error)
	DeleteMember(ID string) error
}

func (c MembersMongoRepo) CreateMember(b models.MemberModel) (models.MemberModel, error) {
	memCollection := c.Client.MemberCollection()
	_, err := memCollection.InsertOne(nil, b)

	if err != nil {
		if c.Client.IsMongoDuplicateError(err) {
			return b, apperrors.DuplicateError{Resource: "member"}
		}
		return b, err
	}
	return b, nil
}

func (c MembersMongoRepo) GetMemberById(ID string) (models.MemberModel, error) {
	var member models.MemberModel
	memCollection := c.Client.MemberCollection()
	ObjectID, e := primitive.ObjectIDFromHex(ID)
	if e != nil {
		return member, e
	}
	id := bson.D{{"_id", ObjectID}}
	err := memCollection.FindOne(nil, id).Decode(&member)
	if err != nil {
		return member, apperrors.ErrorGetting{Resource: "member"}
	}
	return member, nil
}

func (c MembersMongoRepo) FetchMembers(filter interface{}) ([]models.MemberModel, error) {
	var members []models.MemberModel
	memCollection := c.Client.MemberCollection()
	findOptions := options.Find()
	cur, err := memCollection.Find(nil, bson.D{{}}, findOptions)
	if err != nil {
		return members, apperrors.ErrorGetting{Resource: "members"}
	}
	err = cur.All(nil, &members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (c MembersMongoRepo) UpdateMemberDetails(ID string, update models.MemberModel) (models.MemberModel, error) {
	var member models.MemberModel
	memCollection := c.Client.MemberCollection()
	ObjectId, e := primitive.ObjectIDFromHex(ID)
	if e != nil {
		return models.MemberModel{}, e
	}
	update.ID = ObjectId
	id := bson.D{{"_id", ObjectId}}
	updatePayload := bson.M{"$set": update}
	after := options.After

	findOneOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err := memCollection.FindOneAndUpdate(nil, id, updatePayload, &findOneOptions).Decode(&member)
	if err != nil {
		return member, apperrors.ErrorUpdating{Resource: "member"}
	}
	return member, nil
}

func (c MembersMongoRepo) DeleteMember(ID string) error {
	memCollection := c.Client.MemberCollection()
	ObjectId, e := primitive.ObjectIDFromHex(ID)
	if e != nil {
		return e
	}
	id := bson.D{{"_id", ObjectId}}
	_, err := memCollection.DeleteOne(nil, id, nil)
	if err != nil {
		return apperrors.ErrorDeleting{Resource: "member"}
	}
	return nil
}

func NewMemberRepo(conn db.MongoDB) MemberRepositoryI {
	return MembersMongoRepo{Client: conn}
}