package main

import "fmt"

func linearsearch(datalist []int, key int) bool {

	for i, item := range datalist {

		if item == key {

			fmt.Println("Key:", key, "available at:", i, "position")
			return true
		}
	}
	fmt.Println("Key", key, "not available")
	return false
}

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	linearsearch(items, 58)

}
