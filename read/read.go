package read

import (
	"fmt"
	"log"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

func ReadMessages(user string) {

	if !db.UserExists(user) {
		log.Println("Failed to read messages for a user that doesn't exist: " + user + "\n")
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		log.Println("Failed to read messages with the wrong password for: " + user + "\n")
		log.Fatalf("Unable to authenticate user")
	}

	messages := db.GetMessagesForUser(user)
	for _, message := range messages {
		fmt.Println(message)
	}
	log.Println(user + " read their messages successfully\n")
}
