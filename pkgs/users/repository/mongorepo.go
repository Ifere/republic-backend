package userRepo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	db "republic-backend/config/db"
	apperrors "republic-backend/config/errors"
	models "republic-backend/models"
)

type UserMongoRepo struct {
	Client db.MongoDB
}

type UserRepositoryI interface {
	CreateUser(m models.UserModel) (models.UserModel, error)
	GetUserById(ID string) (models.UserModel, error)
	FetchUsers(filter interface{}) ([] models.UserModel, error)
	UpdateUserDetails(ID string, update models.UserModel) (models.UserModel, error)
	DeleteUser(ID string) error
}

func (c UserMongoRepo) CreateUser(b models.UserModel) (models.UserModel, error) {
	userCollection := c.Client.UserCollection()
	_, err := userCollection.InsertOne(nil, b)

	if err != nil {
		if c.Client.IsMongoDuplicateError(err) {
			return b, apperrors.DuplicateError{Resource: "user"}
		}
		return b, err
	}
	return b, nil
}

func (c UserMongoRepo) GetUserById(ID string) (models.UserModel, error) {
	var user models.UserModel
	userCollection := c.Client.UserCollection()
	ObjectID, e := primitive.ObjectIDFromHex(ID)
	if e != nil {
		return user, e
	}
	id := bson.D{{"_id", ObjectID}}
	err := userCollection.FindOne(nil, id).Decode(&user)
	if err != nil {
		return user, apperrors.ErrorGetting{Resource: "user"}
	}
	return user, nil
}

func (c UserMongoRepo) FetchUsers(filter interface{}) ([]models.UserModel, error) {
	var users []models.UserModel
	userCollection := c.Client.UserCollection()
	findOptions := options.Find()
	cur, err := userCollection.Find(nil, bson.D{{}}, findOptions)
	if err != nil {
		return users, apperrors.ErrorGetting{Resource: "users"}
	}
	err = cur.All(nil, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c UserMongoRepo) UpdateUserDetails(ID string, update models.UserModel) (models.UserModel, error) {
	var user models.UserModel
	userCollection := c.Client.UserCollection()
	ObjectId, e := primitive.ObjectIDFromHex(ID)
	if e != nil {
		return models.UserModel{}, e
	}
	update.ID = ObjectId
	id := bson.D{{"_id", ObjectId}}
	updatePayload := bson.M{"$set": update}
	after := options.After

	findOneOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err := userCollection.FindOneAndUpdate(nil, id, updatePayload, &findOneOptions).Decode(&user)
	if err != nil {
		return user, apperrors.ErrorUpdating{Resource: "user"}
	}
	return user, nil
}

func (c UserMongoRepo) DeleteUser(ID string) error {
	userCollection := c.Client.UserCollection()
	ObjectId, e := primitive.ObjectIDFromHex(ID)
	if e != nil {
		return e
	}
	id := bson.D{{"_id", ObjectId}}
	_, err := userCollection.DeleteOne(nil, id, nil)
	if err != nil {
		return apperrors.ErrorDeleting{Resource: "user"}
	}
	return nil
}

func NewUserRepo(conn db.MongoDB) UserRepositoryI {
	return UserMongoRepo{Client: conn}
}
