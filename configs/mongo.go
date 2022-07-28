package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DBClient  *mongo.Client
	DBContext context.Context
	DB        *mongo.Database
)

func ConnectDB() (*mongo.Client, context.Context) {
	var (
		mongoURI       = os.Getenv("MONGO_URI")
		replicaSetName = "rs"
		database       = os.Getenv("DATABASE_NAME")
	)

	DBContext = context.Background()
	client, err := mongo.Connect(DBContext,
		options.Client().ApplyURI(mongoURI),
		options.Client().SetReplicaSet(replicaSetName),
	)
	if err != nil {
		log.Fatal(err)
	}

	DBClient = client

	err = client.Ping(DBContext, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to Database as: ", mongoURI)

	DB = DBClient.Database(database)

	return DBClient, DBContext
}
