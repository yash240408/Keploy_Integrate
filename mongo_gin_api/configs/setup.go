package configs


// === What this file does is :- 
//1. Import the required dependencies.
//2. Create a ConnectDB function that first configures the client to use the correct URI and check for errors. Secondly, we defined a timeout of 10 seconds we wanted to use when trying to connect. Thirdly, check if there is an error while connecting to the database and cancel the connection if the connecting period exceeds 10 seconds. Finally, we pinged the database to test our connection and returned the client instance.
//3. Create a DB variable instance of the ConnectDB. This will come in handy when creating collections.
//4. Create a GetCollection function to retrieve and create collections on the database.

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"github.com/keploy/go-sdk/integrations/kmongo" // add the Keploy import
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *kmongo.Collection { // update the return type to use kmongo.Collection
	db := client.Database("golangAPI")
	col := kmongo.NewCollection(db.Collection(collectionName)) // use kmongo.NewCollection
	return col
}
