package zoneRepo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	db "republic-backend/config/db"
	apperrors "republic-backend/config/errors"
	models "republic-backend/models"
)

type ZoneMongoRepo struct {
	Client db.MongoDB
}

type ZoneRepositoryI interface {
	CreateZone(m models.ZoneModel) (models.ZoneModel, error)
	GetZoneById(ID string) (models.ZoneModel, error)
	FetchZones(filter interface{}) ([] models.ZoneModel, error)
	UpdateZoneDetails(ID string, update models.ZoneModel) (models.ZoneModel, error)
	DeleteZone(ID string) error
}

func (c ZoneMongoRepo) CreateZone(b models.ZoneModel) (models.ZoneModel, error) {
	zoneCollection := c.Client.ZoneCollection()
	_, err := zoneCollection.InsertOne(nil, b)

	if err != nil {
		if c.Client.IsMongoDuplicateError(err) {
			return b, apperrors.DuplicateError{Resource: "user"}
		}
		return b, err
	}
	return b, nil
}

func (c ZoneMongoRepo) GetZoneById(ID string) (models.ZoneModel, error) {
	var zone models.ZoneModel
	userCollection := c.Client.UserCollection()
	ObjectID, e := primitive.ObjectIDFromHex(ID)
	if e != nil {
		return zone, e
	}
	id := bson.D{{"_id", ObjectID}}
	err := userCollection.FindOne(nil, id).Decode(&zone)
	if err != nil {
		return zone, apperrors.ErrorGetting{Resource: "user"}
	}
	return zone, nil
}

func (c ZoneMongoRepo) FetchZones(filter interface{}) ([]models.ZoneModel, error) {
	var zones []models.ZoneModel
	zoneCollection := c.Client.ZoneCollection()
	findOptions := options.Find()
	cur, err := zoneCollection.Find(nil, bson.D{{}}, findOptions)
	if err != nil {
		return zones, apperrors.ErrorGetting{Resource: "zones"}
	}
	err = cur.All(nil, &zones)
	if err != nil {
		return nil, err
	}
	return zones, nil
}

func (c ZoneMongoRepo) UpdateZoneDetails(ID string, update models.ZoneModel) (models.ZoneModel, error) {
	var zone models.ZoneModel
	zoneCollection := c.Client.ZoneCollection()
	ObjectId, e := primitive.ObjectIDFromHex(ID)
	if e != nil {
		return models.ZoneModel{}, e
	}
	update.ID = ObjectId
	id := bson.D{{"_id", ObjectId}}
	updatePayload := bson.M{"$set": update}
	after := options.After

	findOneOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err := zoneCollection.FindOneAndUpdate(nil, id, updatePayload, &findOneOptions).Decode(&zone)
	if err != nil {
		return zone, apperrors.ErrorUpdating{Resource: "user"}
	}
	return zone, nil
}

func (c ZoneMongoRepo) DeleteZone(ID string) error {
	zoneCollection := c.Client.ZoneCollection()
	ObjectId, e := primitive.ObjectIDFromHex(ID)
	if e != nil {
		return e
	}
	id := bson.D{{"_id", ObjectId}}
	_, err := zoneCollection.DeleteOne(nil, id, nil)
	if err != nil {
		return apperrors.ErrorDeleting{Resource: "zone"}
	}
	return nil
}

func NewZoneRepo(conn db.MongoDB) ZoneRepositoryI {
	return ZoneMongoRepo{Client: conn}
}
