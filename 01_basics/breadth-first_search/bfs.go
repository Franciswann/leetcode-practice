package main

import "fmt"

var List = map[string][]string{
	// First layer
	"Me": {"Ivy", "John", "Mandy"},
	// Second layer
	"Ivy":   {""},
	"John":  {"Josh"},
	"Mandy": {"Josh", "Pug", "Bob"},
	// Third layer
	"Jenny": {""},
}

var check_List []string

var searched_List map[string]bool

func AlreadySearch(name string) {
	if searched_List == nil {
		searched_List = make(map[string]bool)
	}
	// note searched person
	searched_List[name] = true
}

func IsAlreadySearched(name string) bool {
	if searched_List == nil {
		return false
	}
	return searched_List[name]
}

func ShouldSkip(name string) bool {
	// skip empty string
	if name == "" {
		return true
	}
	// skip if already searched
	if IsAlreadySearched(name) {
		return true
	}
	return false
}

func AlienVerifier(name string) bool {

	if name[len(name)-1:] == "g" {
		return true
	}
	return false
}

func BFS() string {
	check_List = append(check_List, List["Me"]...)

	for len(check_List) > 0 {

		// skip if should be skipped (empty or already searched)
		if ShouldSkip(check_List[0]) {
			check_List = check_List[1:]
			continue
		}

		if AlienVerifier(check_List[0]) == true {
			return check_List[0] + " is Alien!"
		}

		// add searched person to searched_List
		AlreadySearch(check_List[0])

		// add new list to queue
		check_List = append(check_List, List[check_List[0]]...)

		// remove first person from queue
		check_List = check_List[1:]
	}

	return "Couldn't found one"
}

func main() {
	result := BFS()
	// print result
	fmt.Println(result)
}
