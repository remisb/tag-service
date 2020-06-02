package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	conf    *Config
	mongoDb *mongo.Client
)

const (
	defaultDbName = "winawin"
	mongoTimeout = 15 * time.Second
)

type Config struct {
	MongoDb MongoDb `yaml:"mongoDb"`
}

type MongoDb struct {
	DriverName string `yaml:"driverName"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	DbName     string `yaml:"dbName"`
}

func MongoConnect() *mongo.Client {

	if mongoDb != nil {
		return mongoDb
	}

	//c := conf
	c := Config{MongoDb: MongoDb{
		DriverName: "mongodb",
		Host: "18.188.233.74",
		DbName: "winawin",
		Port: "27017",
	}}

	mongoURI := fmt.Sprintf("%s://%s:%s", c.MongoDb.DriverName, c.MongoDb.Host, c.MongoDb.Port)
	clientOptions := options.Client().ApplyURI(mongoURI)

	var err error
	mongoDb, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoDb.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return mongoDb
}

func GetWinaWinDb() *mongo.Database {
	dbName := "winawin"
	//if dbName == "" {
	//	dbName = defaultDbName
	//}
	return MongoConnect().Database(dbName)
}

func GetTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), mongoTimeout)
}

func GetMongoCollection(name string) *mongo.Collection {
	return GetWinaWinDb().Collection(name)
}
