package database

import (
	"context"
	"errors"

	"github.com/ZondaF12/my-pocket-garage/internal/handlers/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var database string

func GetDBCollection(name string) *mongo.Collection {
	return mongoClient.Database(database).Collection(name)
}

func StartMongoDB(uri, dbName string) error {
	if uri == "" {
		return errors.New("you must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	if dbName == "" {
		return errors.New("you must set your 'DATABASE' environmental variable")
	} else {
		database = dbName
	}

	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		return errors.New("can't verify a connection")
	}

	return nil
}

func CloseMongoDB() {
	err := mongoClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}

type UserVehicle struct {
	UserID        string     `json:"userId" bson:"userId"`
	Active        bool       `json:"active" bson:"active"`
	Registration  string     `json:"registration" bson:"registration"`
	Make          string     `json:"make" bson:"make"`
	Model         string     `json:"model" bson:"model"`
	Year          int        `json:"year" bson:"year"`
	EngineSize    int        `json:"engineSize" bson:"engineSize"`
	Color         string     `json:"color" bson:"color"`
	Registered    string     `json:"registered" bson:"registered"`
	TaxDate       string     `json:"taxDate" bson:"taxDate"`
	MotDate       string     `json:"motDate" bson:"motDate"`
	InsuranceDate string     `json:"insuranceDate" bson:"insuranceDate"`
	ServiceDate   string     `json:"serviceDate" bson:"serviceDate"`
	Activity      []Activity `json:"activity" bson:"activity"`
}

type Activity struct {
	UserID        string `json:"userId" bson:"userId"`
	Registration  string `json:"registration" bson:"registration"`
	Title         string `json:"title" bson:"title"`
	Date          string `json:"date" bson:"date"`
	Cost          string `json:"cost" bson:"cost"`
	ServiceCentre string `json:"serviceCentre" bson:"serviceCentre"`
	Notes         string `json:"notes" bson:"notes"`
}

func AddUserVehicle(userId string, registration string) error {
	result := CheckUserVehicleExists(userId, registration)
	if !result {
		return errors.New("vehicle already added")
	}

	res, err := tools.DoVehicleInfoRequest(registration)
	if err != nil {
		return err
	}

	motRes, err := tools.DoVehicleMotRequest(registration)
	if err != nil {
		return err
	}

	userVehicles, err := GetUserVehicles(userId)
	if err != nil {
		return err
	}
	setActive := userVehicles == nil

	// validate the body
	newUserVehicle := UserVehicle{UserID: userId, Active: setActive, Registration: registration, Make: res.Make, Model: motRes[0].Model, Year: res.YearOfManufacture, EngineSize: res.EngineCapacity, Color: motRes[0].PrimaryColour, Registered: motRes[0].FirstUsedDate, TaxDate: res.TaxDueDate, MotDate: res.MotExpiryDate, InsuranceDate: "", ServiceDate: "", Activity: []Activity{}}

	// create the price alert
	coll := GetDBCollection("User Vehicles")
	_, err = coll.InsertOne(context.Background(), newUserVehicle)
	if err != nil {
		return err
	}

	return nil
}

func CheckUserVehicleExists(userId string, registration string) bool {
	coll := GetDBCollection("User Vehicles")
	filter := bson.D{{Key: "userId", Value: userId}, {Key: "registration", Value: registration}}

	var res UserVehicle
	cur := coll.FindOne(context.Background(), filter).Decode(&res)

	return cur != nil
}

func GetUserVehicles(userId string) ([]UserVehicle, error) {
	coll := GetDBCollection("User Vehicles")
	filter := bson.D{{Key: "userId", Value: userId}}

	var res []UserVehicle
	cur, err := coll.Find(context.Background(), filter)
	if err != nil {
		return []UserVehicle{}, err
	}

	for cur.Next(context.Background()) {
		//Create a value into which the single document can be decoded
		var elem UserVehicle
		err := cur.Decode(&elem)
		if err != nil {
			return []UserVehicle{}, err
		}
		res = append(res, elem)
	}

	return res, err
}

func AddVehicleActivity(arr Activity) error {
	actColl := GetDBCollection("Activity")
	_, err := actColl.InsertOne(context.Background(), arr)
	if err != nil {
		return err
	}

	uvColl := GetDBCollection("User Vehicles")
	filter := bson.D{{Key: "userId", Value: arr.UserID}, {Key: "registration", Value: arr.Registration}}
	update := bson.M{"$push": bson.M{"activity": arr}}

	var updatedDoc Activity
	err = uvColl.FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedDoc)
	if err != nil {
		return err
	}

	return nil
}

func GetActiveVehicle(userId string) (UserVehicle, error) {
	coll := GetDBCollection("User Vehicles")
	filter := bson.D{{Key: "userId", Value: userId}, {Key: "active", Value: true}}

	var active UserVehicle
	err := coll.FindOne(context.Background(), filter).Decode(&active)
	if err != nil {
		return UserVehicle{}, err
	}

	return active, err
}

func SetActiveVehicle(userId string, registration string) error {
	coll := GetDBCollection("User Vehicles")
	filter := bson.D{{Key: "userId", Value: userId}, {Key: "active", Value: true}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "active", Value: false}}}}

	_, err := coll.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return err
	}

	filter = bson.D{{Key: "userId", Value: userId}, {Key: "registration", Value: registration}}
	update = bson.D{{Key: "$set", Value: bson.D{{Key: "active", Value: true}}}}

	var updatedDoc Activity
	err = coll.FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedDoc)
	if err != nil {
		return err
	}

	return err
}