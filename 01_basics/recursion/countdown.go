package main

import "fmt"

func CountDown(n int) string {
	fmt.Println(n)

	if n <= 1 {
		return "1"
	}

	return fmt.Sprintf("%d\n%s", n, CountDown(n-1))
}

func main() {
	CountDown(5)
}

// CountDown(5)
// |-- fmt.Println(5) -> 5
// |-- return fmt.Sprintf("%d\n%s", 5, CountDown(4))
// |	CountDown(4)
// |	|-- fmt.Println(4) -> 4
// |	|-- return fmt.Sprintf("%d\n%s", 4, CountDown(3))
// |	|	CountDown(3)
// |	|	|-- fmt.Println(3) -> 3
// |	|	|-- return fmt.Sprintf("%d\n%s", 3, CountDown(2))
// |	|	|	CountDown(2)
// |	|	|	|-- fmt.Println(2) -> 2
// |	|	|	|-- return fmt.Sprintf("%d\n%s", 2, CountDown(1))
// |	|	|	|	CountDown(1)
// |	|	|	|	|-- fmt.Println(1) -> 1
// |	|	|	|	|-- return "1"

// fmt.Sprintf("%d\n%s", n, CountDown(n - 1))
// "1" -> "2\n1" -> "3\n2\n1" -> "4\n3\n2\n1" -> "5\n4\n3\n2\n1"
