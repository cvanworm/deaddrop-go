package read

import (
	"fmt"
	"log"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

func ReadMessages(user string) {
	f, e := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		log.Fatalf("Error opening file: %v", e)
	}
	defer f.Close()

	if !db.UserExists(user) {
		if _, err := f.WriteString("Failed to read messages for a user that doesn't exist: " + user + "\n"); err != nil {
			log.Println(err)
		}
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		if _, err := f.WriteString("Failed to read messages with the wrong password for: " + user + "\n"); err != nil {
			log.Println(err)
		}
		log.Fatalf("Unable to authenticate user")
	}

	messages := db.GetMessagesForUser(user)
	for _, message := range messages {
		fmt.Println(message)
	}
	if _, err := f.WriteString(user + " read their messages successfully\n"); err != nil {
		log.Println(err)
	}
}
