package quotemanager

import (
  "os"
  "log"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"
)

type Quote struct {
  AddedBy   string
  Text      string
  Timestamp string
}

var client mongo.Client
var collection *mongo.Collection

func LoadQuoteList() error {
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
  collection = client.Database("qmngr").Collection("qmngr")

  return nil
}

func AddQuote(by string, text string, timestamp string){
  quote := Quote{AddedBy: by, Text: text, Timestamp: timestamp}
  _, err := collection.InsertOne(context.TODO(), quote)
  if err != nil {
    log.Fatal("Problem pushing new quote to MongoDB")
  }
}

func GetRandomQuote() Quote {
  query := bson.D{
    {"$sample", bson.D{
      {"size", 1},
    }},
  }
  cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{query})
  if err != nil {
    log.Fatal(err)
  }
  var results[]bson.M
  if err = cursor.All(context.TODO(), &results); err != nil {
    log.Fatal(err)
  }
  var ret Quote
  var ok bool
  ret.AddedBy, ok = results[0]["addedby"].(string)
  if !ok {
    log.Fatal("Expected string from mongo results")
  }
  ret.Text, ok = results[0]["text"].(string)
  if !ok {
    log.Fatal("Expected string from mongo results")
  }
  ret.Timestamp, ok = results[0]["timestamp"].(string)
  if !ok {
    log.Fatal("Expected string from mongo results")
  }

  return ret
}
