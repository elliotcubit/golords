package state

// Sets up connection and other variables for the rest of the library

import (
  "os"
  "log"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var client mongo.Client

// Collections
var plusplusColl *mongo.Collection
var qmanagerColl *mongo.Collection

func init(){
  URI := os.Getenv("MONGO_URI")
  if URI == "" {
    log.Fatal("No URI found for MongoDB")
  }
  clientOptions := options.Client().ApplyURI(URI)
  client, err := mongo.Connect(context.TODO(), clientOptions)
  if err != nil {
    log.Fatal(err)
  }
  err = client.Ping(context.TODO(), nil)
  if err != nil {
    log.Fatal(err)
  }
  plusplusColl = client.Database("plusplus").Collection("plusplus")
  qmanagerColl = client.Database("qmngr").Collection("qmngr")
  log.Println("Successfully connected to MongoDB")
}
