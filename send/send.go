package send

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

// SendMessage takes a destination username and will
// prompt the user for a message to send to that user
func SendMessage(to string) {
	f, e := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		log.Fatalf("Error opening file: %v", e)
	}
	defer f.Close()
	if !db.UserExists(to) {
		if _, err := f.WriteString("Destination user does not exist: " + to + "\n"); err != nil {
			log.Println(err)
		}
		log.Fatalf("Destination user does not exist")
	}

	username := loginUserName()
	if !db.UserExists(username) {
		if _, err := f.WriteString("User tried to send a message from a user that does not exist: " + username + "\n"); err != nil {
			log.Println(err)
		}
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(username)
	if err != nil {
		if _, err := f.WriteString(username + " failed to login to send a message\n"); err != nil {
			log.Println(err)
		}
		log.Fatalf("Unable to authenticate user")
	}

	message := getUserMessage()
	
	if _, err := f.WriteString(username + " sent a message to " + to + "\n"); err != nil {
		log.Println(err)
	}
	db.SaveMessage(message, to)
}

// getUserMessage prompts the user for the message to send
// and returns it
func getUserMessage() string {
	fmt.Println("Enter your message: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return text
}

func loginUserName() string{
	fmt.Println("Please log in to send a message.\n Username: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return strings.Trim(text, "\n\t ")
}