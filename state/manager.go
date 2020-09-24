package state

import (
	"fmt"
	"log"
)

type Quote struct {
	AddedBy   string
	Text      string
	Timestamp string
}

var createQuoteStatement string = `INSERT INTO quotes(serverID, userID, quote, timestamp) VALUES ('%s', '%s', '%s', '%s')`
var getRandomQuoteStatement string = `SELECT userID, quote, timestamp FROM quotes WHERE serverID='%s' ORDER BY RANDOM() LIMIT 1`

func AddQuote(server, user, quote, timestamp string) {
	_, err := database.Exec(fmt.Sprintf(createQuoteStatement, server, user, quote, timestamp))
	if err != nil {
		log.Println("Failed to add quote to database")
	}
}

func GetRandomQuote(serverID string) (Quote, error) {
	var result Quote
	rows, err := database.Query(fmt.Sprintf(getRandomQuoteStatement, serverID))
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&result.AddedBy, &result.Text, &result.Timestamp)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}
