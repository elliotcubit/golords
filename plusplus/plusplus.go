package plusplus

import (
  "os"
  "log"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"
)

// Called when module first loaded
func init() {
  err := LoadPlusPlus()
  if err != nil {
    log.Fatal(err)
  }
}

type Info struct {
  User string
  Score int
}

var client mongo.Client
var collection *mongo.Collection

func LoadPlusPlus() error {
  URI := os.Getenv("MONGO_URI")
  if URI == "" {
    log.Fatal("No URI found for MongoDB")
  }
  clientOptions := options.Client().ApplyURI(URI)
  client, err := mongo.Connect(context.TODO(), clientOptions)
  if err != nil {
    return err
  }
  err = client.Ping(context.TODO(), nil)
  if err != nil {
    return err
  }
  log.Print("Successfully connected to MongoDB")
  collection = client.Database("plusplus").Collection("plusplus")

  return nil
}

func PlusPlus(ident string) {
  query := bson.M{
    "user": ident,
  }

  var result bson.M
  err := collection.FindOne(context.TODO(), query).Decode(&result)
  if err != nil {
    log.Println(err)
  }

  log.Println(result)
}

func MinusMinus(ident string) {

}
