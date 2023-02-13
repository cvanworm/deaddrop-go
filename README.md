# deaddrop-go

A deaddrop utility written in Go. Put files in a database behind a password to be retrieved at a later date.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). This is the base repository to be forked and updated for various assignments. Alternative language versions are available in:
- [Javascript](https://github.com/andey-robins/deaddrop-js)
- [Rust](https://github.com/andey-robins/deaddrop-rs)

## Versioning

`deaddrop-go` is built with:
- go version go1.19.4 linux/amd64

## Usage

`go run main.go --help` for instructions

Then run `go run main.go -new -user <username here>` and you will be prompted to create the initial password.

## Database

Data gets stored into the local database file dd.db. This file will not by synched to git repos. Delete this file if you don't set up a user properly on the first go

## Logging Strategy

Creating a user: Each time an existing user attempts to create a new user, the system will log the name trying to create the user and then if it was successful or not. 
Reading a message: Each time a user tries to read their messages, the system will log whether the user doesn't exist, the authentication fails, or if the user is able to read their messages successfully.
Sending a message: Users now have to login to an existing account in order to send a message. The system will log whether the user tries to send a message from an account that doesn't exist, the authentication fails, or if the message is sent successfully. If successful, the log will state the user that sent a message and who recieves the message. For confidentiality, the messages are still kept private, and the sender remains annonymous to the recipient.
