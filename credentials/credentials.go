package credentials

import (
  "os"
  "errors"
)

// Struct storings credentials
type Credentials struct {
  Id     string
  Secret string
  Token  string
}

// Returns credentials populated from a JSON filepath provided.
func LoadCreds() (Credentials, error) {
  var creds Credentials

  creds.Id = os.Getenv("DISCORD_ID")
  creds.Secret = os.Getenv("DISCORD_SECRET")
  creds.Token = os.Getenv("DISCORD_TOKEN")

  if creds.Id == "" || creds.Secret == "" || creds.Token == "" {
    return creds, errors.New("Discord credentials not found in environment variables")
  }
  return creds, nil
}
