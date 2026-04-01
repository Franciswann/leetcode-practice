package main

import "fmt"

// Structure:
// pics/
// ├── 2001/
// │   ├── a.png
// │   └── space.png
// └── odyssey.png
var fileSystem = map[string][]string{
	"pics": {"2001", "odyssey.png"},
	"2001": {"a.png", "space.png"},
}

// Returns true if it's a directory, false if it's a file or doesn't exist.
func isFolder(name string) bool {
	_, exists := fileSystem[name]
	return exists
}

func printBFS(start string) {
	// to-do list (Queue)
	todo := []string{start}

	for len(todo) > 0 {
		// Dequeue: get the first directory from the queue
		current := todo[0]
		todo = todo[1:]

		// Process all items in the current directory
		items := fileSystem[current]
		for _, item := range items {
			if isFolder(item) {
				// If it's a directory, enqueue it for later processing
				todo = append(todo, item)
			} else {
				// If it's a file, print it immediately
				fmt.Println(item)
			}
		}
	}
}

func printDFS(current string) {
	// Get all items in the current directory
	items := fileSystem[current]

	for _, item := range items {
		if isFolder(item) {
			// If it's a directory, recursively traverse it first (pre-order DFS)
			printDFS(item)
		} else {
			// If it's a file, print it immediately
			fmt.Println(item)
		}

	}

}

func main() {
	fmt.Println("-----Breadth-First Search-----")
	printBFS("pics")
	fmt.Println("-----Depth-First Search-----")
	printDFS("pics")
}
