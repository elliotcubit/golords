package plusplus

import (
  "os"
  "log"
  "context"
  "fmt"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"

  "github.com/bwmarrin/discordgo"
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

func PeopleQuery(users []*discordgo.User) (string, error) {
  if len(users) == 0 {
    return "", fmt.Errorf("No users provided")
  }

  var users_query []string
  for _, user := range users {
    users_query = append(users_query, user.String())
  }

  log.Println(users_query)

  filter := bson.M{
    "user": bson.M{
      "$in": users_query,
    },
  }

  cursor, err := collection.Find(context.TODO(), filter)
  if err != nil {
    // TODO all these people actually just have zero stacks
    return "", fmt.Errorf("No documents")
  }

  var results []bson.M
  if err := cursor.All(context.TODO(), &results); err != nil {
    return "", fmt.Errorf("Error unmarshalling results")
  }

  outStr := ""
  for _, result := range results {
    if err != nil {
      return "", fmt.Errorf("Something awful happened")
    }
    score := result["score"].(int32)
    name := result["user"].(string)
    outStr = outStr + fmt.Sprintf("%v has %d stacks\n", name, score)
  }

  return outStr, nil
}

// TODO look into $max with $limit to solve this problem
func TopQuery() (string, error) {
  return "", fmt.Errorf("Not implemented lol fuck you")
}

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

func PlusPlus(ident string, amt int) (int, error) {
  filter := bson.M{"user": ident}
  update := bson.M{"$inc": bson.M{"score": amt}}

  var result bson.M
  err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)

  if err != nil {
    // Add user
    err = CreateUser(ident, amt)
    if err != nil {
      return 0, err
    }
    return amt, nil
  }

  // Updated score
  return int(result["score"].(int32))+amt, nil
}

func MinusMinus(ident string, amt int) (int, error) {
  filter := bson.M{"user": ident}
  update := bson.M{"$inc": bson.M{"score": -amt}}

  var result bson.M
  err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)

  if err != nil {
    // Create with negative karma lol
    err = CreateUser(ident, -amt)
    if err != nil {
      return 0, err
    }
    return -amt, nil
  }

  // Updated score
  return int(result["score"].(int32))-amt, nil
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
