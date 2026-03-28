package main

import "fmt"

var List = map[string][]string{
	// First layer
	"Me": {"Ivy", "John", "Mandy"},
	// Second layer
	// "Ivy":   {""},
	"John":  {"Josh"},
	"Mandy": {"Josh", "Pug", "Bob"},
	// Third layer
	// "Jenny": {""},
}

var check_List []string

// var searched_List []string

func AlienVerifier(name string) bool {
	// name = check_List[0]

	if name[len(name)-1:] == "g" {
		return true
	}
	return false
}

func BFS() string {
	check_List = append(check_List, List["Me"]...)

	for len(check_List) > 0 {
		// fmt.Println(check_List[0])
		// fmt.Println(check_List)

		if AlienVerifier(check_List[0]) == true {
			return check_List[0] + " is Alien!"
		}
		// add new list to queue
		check_List = append(check_List, List[check_List[0]]...)

		// // add searched person to searched_List
		// searched_List = append(searched_List, check_List[0])

		// remove first person from queue
		check_List = check_List[1:]
	}

	return "Couldn't found one"
}

func main() {
	// BFS()

	// fmt.Println(check_List)
	// fmt.Println(check_List[0])
	// // x := check_List[0]
	// xy := check_List[0][len(check_List[0])-1:]
	// fmt.Println(xy)

	fmt.Println(BFS())

}

// ```

// list

// 1. add "Me" List

// 2. verify "Me" List[],
// if True -> return
// else -> add a list by someone been removed

// 3. verify List[] again

// 4. End
// ```
