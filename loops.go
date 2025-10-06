// A simple console-based To-Do List program written in Go.
// This program demonstrates core Go concepts such as variables, slices, loops,

package main

// Importing the "fmt" package for input/output formatting and printing to console.
import "fmt"

// --- Entry Point ---
// Every Go program begins execution from the main() function.
func main() {
	// --- Step 1: Define Tasks ---
	// Declaring string variables for each task item.
	var shortGolang = "Practice GO language"          // Short learning task
	var fullGolang = "Workout at the Gym"             // Physical activity task
	var rewardDessert = "Reward myself with a cookie" // Small motivational reward

	// --- Step 2: Store Tasks in a Slice ---
	// A slice in Go is a dynamic list structure similar to arrays, but resizable.
	var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	// --- Step 3: Display the Header ---
	fmt.Println("WELCOME TO MY TO-DO LIST APP!")  // Welcome message
	fmt.Println("Here are my tasks for the day:") // Introduction text

	// --- Step 4: Loop Through and Display Tasks ---
	// The range keyword allows iteration through slices with both index and value.
	for index, task := range taskItems {
		// fmt.Printf is used for formatted printing.
		// %d prints an integer (the task number) and %s prints a string (the task name).
		fmt.Printf("%d: %s\n", index+1, task)
	}

	// --- End of Program ---
	// Once the loop ends, all tasks are displayed in a numbered list.
}
