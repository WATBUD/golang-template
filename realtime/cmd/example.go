package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"mai.today/realtime"
)

// Payload represents the structure of the message to be sent.
//
//	{
//		"user_id": "<USER-ID>",
//		"message": "Hello from Go!",
//		"time": "2024-06-29 08:31:07.781248044 +0000 UTC m=+0.006645668"
//	}
type Payload struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

func main() {
	// Initialize a new Payload with a default message.
	d := Payload{
		UserID:  "<USER-ID>",
		Message: "Hello from Go!",
		Time:    time.Now().String(),
	}

	// Create a new scanner to read from standard input.
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter message (type 'exit' to quit):")

	for {
		// Print a prompt.
		fmt.Print("> ")

		// Read the input.
		scanner.Scan()
		input := scanner.Text()

		// Check for 'exit' to break the loop.
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		// Update the message and time in the payload.
		d.Message = input
		d.Time = time.Now().String()

		// Send the payload.
		send(d)
	}

	// Check for errors during scanning.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

// send sends the payload to the specified user using Centrifugo.
func send(d Payload) {
	// Get the singleton instance of the Realtime client.
	client := realtime.Instance()

	// Publish the payload to the user's channel.
	res, err := client.Publish(context.Background(), d.UserID, d)

	// Log the result or error.
	log.Printf("Publish Result: %v", res)
	if err != nil {
		log.Printf("Publish error: %v", err)
	}
}
