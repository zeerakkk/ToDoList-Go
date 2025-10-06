// A simple HTTP-based To-Do List web application written in Go.
// This program demonstrates basic web server setup using the "net/http" package,
// routing multiple endpoints, and dynamic HTML/text output in response to client requests.

package main

// Import Required Packages
import (
	"fmt"      // Used for formatted input/output
	"net/http" // Provides HTTP client and server implementations
)

//	Step 1: Define Initial Tasks
//
// Global variables for task storage.
// These represent static tasks displayed by the web server.
var shortGolang = "Practice GO language"
var fullGolang = "Workout at the Gym"
var rewardDessert = "Reward myself with a cookie"

// Slice (dynamic array) holding all initial tasks.
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

// Step 2: Main Function (Program Entry Point)
func main() {
	// Register HTTP route handlers.
	// Each path ("/" and "/show-tasks") maps to a specific function.
	http.HandleFunc("/", helloUser)           // Root endpoint: greets the user
	http.HandleFunc("/show-tasks", showTasks) // Endpoint: displays all tasks

	// Start an HTTP server on port 8080.
	// The second argument (nil) means the default multiplexer (router) is used.
	fmt.Println(" Server is running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

// Step 3: Handler Function: showTasks
// Purpose: Displays all tasks from the to-do list when a user visits "/show-tasks".
// Parameters:
//
//	writer: used to send data back to the client (browser).
//	request: holds details about the HTTP request.
func showTasks(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "###### To-Do List ######")
	for _, task := range taskItems {
		fmt.Fprintln(writer, "- ", task) // Print each task as a new line in the HTTP response
	}
}

//	Step 4: Handler Function: helloUser
//
// Purpose: Sends a greeting message to the user when visiting the root path ("/").
// Parameters:
//
//	writer: response writer to send HTML/text back to the browser.
//	request: incoming client request (unused here but required by the function signature).
func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user!  Welcome to our Todolist App!\nVisit /show-tasks to view your tasks."
	fmt.Fprintln(writer, greeting)
}
