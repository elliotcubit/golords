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

func PlusPlus(ident string) (int, error) {
  filter := bson.M{"user": ident}
  update := bson.M{"$inc": bson.M{"score": 1}}

  var result bson.M
  err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)

  if err != nil {
    // Add user
    err = CreateUser(ident, 1)
    if err != nil {
      return 0, err
    }
    // Always 1
    return 1, nil
  }

  // Updated score
  return result["score"].(int), nil
}

func MinusMinus(ident string) (int, error) {
  filter := bson.M{"user": ident}
  update := bson.M{"$inc": bson.M{"score": -1}}

  var result bson.M
  err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)

  if err != nil {
    // Create with negative karma lol
    err = CreateUser(ident, -1)
    if err != nil {
      return 0, err
    }
    // Always -1
    return -1, nil
  }

  // Updated score
  return result["score"].(int), nil
}

func CreateUser(ident string, score int) error {
  inf := Info{User: ident, Score: score}
  _, err := collection.InsertOne(context.TODO(), inf)
  if err != nil {
    log.Println("Problem pushing new user to mongoDB")
    return err
  }
  return nil
}
