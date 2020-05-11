package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
	"time"
)


const (
	dbName = "republic"
)

type MongoDB struct {
	Client *mongo.Client
}

// create a connection to mongodb
func (db *MongoDB) Connect() error {
	LocalUrl := "mongodb://localhost:27017/platphom"
	// todo change mongodb shard address
	//url := "mongodb://localhost:27017/+replicaSet=rs"
	//if os.Getenv("env") != "dev" {
	//	url = "mongodb+srv://admin003:HSaY4gr8070bGQdM@saga-production-jghko.mongodb.net/test?retryWrites=true&w=majority"
	//}

	client, err := mongo.NewClient(options.Client().ApplyURI(LocalUrl))

	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		return err
	}

	logrus.Println("db connected")
	db.Client = client

	return nil
}

func connect() *mongo.Client {
	url := "mongodb://localhost:27017/+replicaSet=rs"
	if os.Getenv("env") != "dev" {
		url = "mongodb+srv://admin003:HSaY4gr8070bGQdM@saga-production-jghko.mongodb.net/test?retryWrites=true&w=majority"
	}

	client, _ := mongo.NewClient(options.Client().ApplyURI(url))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	logrus.Println("db connected")

	return client
}

// create a connection to the test db
func (db *MongoDB) ConnectTest() error {
	url := "mongodb://localhost:27017"
	if os.Getenv("env") != "dev" {
		url = os.Getenv("mongo_test_url")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		return err
	}

	logrus.Println("db connected")
	db.Client = client

	return nil
}

// create a connection to mongodb
func (db *MongoDB) TestTearDown(coll string) error {
	c := db.Client.Database(dbName).Collection(coll)

	return c.Drop(nil)
}


func (db MongoDB) CreateTTLIndex(coll *mongo.Collection, expir *int32) error {
	index := mongo.IndexModel{
		Keys:    bson.D{{"created_at", 1}},
		Options: &options.IndexOptions{ExpireAfterSeconds: expir},
	}
	_, err := coll.Indexes().CreateOne(nil, index)

	if err != nil {
		return err
	}

	return nil
}

// returns flag for duplicate error

func (db MongoDB) IsMongoDuplicateError(err error) bool {

	if strings.Contains(err.Error(), "E11000") {
		return true
	}
	return false
}



// returns an instance of user collection
func (db MongoDB) UserCollection() *mongo.Collection {
	return db.Client.Database(dbName).Collection("users")
}
func (db MongoDB) ZoneCollection() *mongo.Collection {
	return db.Client.Database(dbName).Collection("zones")
}
func (db MongoDB) MemberCollection() *mongo.Collection{
	return db.Client.Database(dbName).Collection("members")
}