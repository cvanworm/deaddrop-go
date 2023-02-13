package new

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

// Create a NewUser as authorized by the user 'user'
func NewUser(user string) {
	f, e := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		log.Fatalf("Error opening file: %v", e)
	}
	defer f.Close()

	if _, err := f.WriteString(user + " is creating a new user\n"); err != nil {
		log.Println(err)
	}

	if !db.NoUsers() && !db.UserExists(user) {
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		if _, err := f.WriteString(user + " couldn't log in while creating a new user\n"); err != nil {
			log.Println(err)
		}
		log.Fatalf("Unable to authenticate user")
	}

	newUser := getNewUsername()
	newPassHash, err := session.GetPassword()
	if err != nil {
		log.Fatalf("Unable to get password hash")
	}

	err = db.SetUserPassHash(newUser, newPassHash)
	if err != nil {
		log.Fatalf("Unable to create new user")
	}

	if _, err := f.WriteString(user + " successfully created a new user: " + newUser + "\n"); err != nil {
		log.Println(err)
	}
}

// getUserMessage prompts the user for the message to send
// and returns it
func getNewUsername() string {

	fmt.Println("Enter the username for the new user: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	return strings.Trim(text, "\n\t ")
}
