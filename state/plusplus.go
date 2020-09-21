package state

import (
  "log"
  "context"
  "fmt"

  "go.mongodb.org/mongo-driver/bson"

  "github.com/bwmarrin/discordgo"
)

type PlusInfo struct {
  User string
  Score int
}

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

  cursor, err := plusplusColl.Find(context.TODO(), filter)
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

func PlusPlus(ident string, amt int) (int, error) {
  filter := bson.M{"user": ident}
  update := bson.M{"$inc": bson.M{"score": amt}}

  var result bson.M
  err := plusplusColl.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)

  if err != nil {
    // Add user
    err = ppCreateUser(ident, amt)
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
  err := plusplusColl.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)

  if err != nil {
    // Create with negative karma lol
    err = ppCreateUser(ident, -amt)
    if err != nil {
      return 0, err
    }
    return -amt, nil
  }

  // Updated score
  return int(result["score"].(int32))-amt, nil
}

func ppCreateUser(ident string, score int) error {
  inf := PlusInfo{User: ident, Score: score}
  _, err := plusplusColl.InsertOne(context.TODO(), inf)
  if err != nil {
    log.Println("Problem pushing new user to mongoDB")
    return err
  }
  return nil
}
