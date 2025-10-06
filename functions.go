// This program demonstrates function creation, slice manipulation (dynamic arrays)

package main

// Importing the "fmt" package for formatted I/O operations.
import "fmt"

// --- Entry Point ---
// main() is the starting point of every Go program.
func main() {
	// Displaying the app header.
	fmt.Println("WELCOME TO MY TO-DO LIST APP!")

	// Step 1: Define Initial Tasks
	// Each variable represents a task in the to-do list.
	var shortGolang = "Practice GO language"
	var fullGolang = "Workout at the Gym"
	var rewardDessert = "Reward myself with a cookie"

	// Storing all tasks in a slice — a resizable list structure in Go.
	var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	//  Step 2: Print the Initial Task List
	printTasks(taskItems)
	fmt.Println() // Blank line for cleaner console spacing

	// Step 3: Dynamically Add More Tasks
	// The addTask() function returns an updated slice with the new task appended.
	taskItems = addTask(taskItems, "Go for a run")
	taskItems = addTask(taskItems, "Practice coding in Python")

	// Step 4: Print the Updated Task List
	fmt.Println("Updated list:")
	printTasks(taskItems)
}

// Function: printTasks
// Purpose: Prints all the tasks in the list in a numbered format.
// Parameters:
// - taskItems []string → A slice of strings containing task descriptions.
func printTasks(taskItems []string) {
	fmt.Println("List of my Todos:")
	for index, task := range taskItems {
		// Print task number (index+1) and description.
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

// Function: addTask
// Purpose: Adds a new task to the existing to-do list.
// Parameters:
// - taskItems []string → The current list of tasks.
// - newTask string → The new task to be added.
// Returns:
// - []string → A new slice that includes the added task.
func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask) // Append adds the new item to the slice
	return updatedTaskItems                           // Return the updated list
}
