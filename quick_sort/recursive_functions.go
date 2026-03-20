package main

// ex:
// [2, 4, 6] -> 2 + [4, 6] -> 2 + 4 + 6

// recursive sum
func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + Sum(arr[1:])
}

// ex:
// [5, 7, 100, 42] -> 4

// recursive count
func Count(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else {
		return 1 + Count(arr[1:])
	}
}

func Find_Max(arr []int) int {

	if len(arr) == 1 {
		return arr[0]
	}

	resOfRest := Find_Max(arr[1:])

	if arr[0] > resOfRest {
		return arr[0]
	}
	return resOfRest

}

// ex:
// 10, 5, 2, 3

// Less pivot greater
// [5, 2, 3] [10] []
//     |
// [2, 3][5][]
//   |
// [2][3]

func Quick_Sort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	var less []int
	pivot := arr[0]
	var greater []int

	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	result := append(Quick_Sort(less), pivot)
	result = append(result, Quick_Sort(greater)...)
	return result
}
