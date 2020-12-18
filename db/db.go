package db

import (
	"context"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb://localhost:27017/mydb?compressors=disabled&gssapiServiceName=mongodb"

func MongoDbConnect() (*mongo.Database, context.Context) {
	clientConfs := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientConfs)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()
	return client.Database("mydb"), ctx

}
