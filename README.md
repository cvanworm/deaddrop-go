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

MAC Strategy:
Logging: To start off, I fixed my logging system from last assignment. It helped clean up my code a lot and works much better now. Each log is still the same as before, but now it has the time and date included.

Sender information: Next, I began by updating the database to keep track of who sends each message. This was done by adding a new collumn to the "Messages" database and storing the Id of each sender. After that, I also updated the read.go file to allow the recipient of a message to see the username of anyone whos sends them a message. 

MAC: For this, I altered the Messages database a little bit more to include an sha hash of each message. Each message is encoded and decoded with a secret key held in the .env file. This simple encryption allows for us to check the authenticity of each message when a user tries to read their messages. In the read.go file, I included a small function used to verify the hash of the message. This will return true or false based on whether the message was altered or not. If the message is authentic, it will be printed to the reader along with who sent the message. If it is not verified, the system will log that a message has been altered or not been verified, alert the user that the message could not be verified, but print it to be read anyway. 

