package main

import (
	"fmt"
	"sort"
	"strings"
)

// Person represents a person with name and age
type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

// Address represents a physical address
// The omitempty can be modernized to omitzero in Go 1.24
type Address struct {
	Street  string `json:"street,omitempty"`
	City    string `json:"city,omitempty"`
	ZipCode int    `json:"zip_code,omitempty"`
}

func main() {
	// Example 1: if/else conditional assignment (can use min)
	a := 10
	b := 20

	var minVal int
	if a < b {
		minVal = a
	} else {
		minVal = b
	}
	fmt.Println("Minimum:", minVal)

	// Example 2: sort.Slice usage (can use slices.Sort)
	names := []string{"Alice", "Bob", "Charlie"}
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
	fmt.Println("Sorted names:", names)

	// Example 3: interface{} usage (can use any)
	var data interface{} = "hello"
	fmt.Println("Data:", data)

	// Example 4: append with nil slice (can use slices.Clone)
	original := []int{1, 2, 3}
	clone := append([]int(nil), original...)
	fmt.Println("Clone:", clone)

	// Example 5: []byte(fmt.Sprintf) (can use fmt.Appendf)
	id := 123
	byteData := []byte(fmt.Sprintf("ID: %d", id))
	fmt.Println("Byte data length:", len(byteData))

	// Example 6: loop with range checking for max (can use max)
	numbers := []int{5, 8, 2, 9, 1}
	maxNum := numbers[0]
	for _, n := range numbers[1:] {
		if n > maxNum {
			maxNum = n
		}
	}
	fmt.Println("Maximum:", maxNum)

	// Example 7: deleting an element from a slice
	fruits := []string{"apple", "banana", "cherry", "date"}
	i := 1 // Remove "banana"
	fruits = append(fruits[:i], fruits[i+1:]...)
	fmt.Println("Fruits after removal:", fruits)

	// Example 8: traditional for loop (can use range int in Go 1.22)
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// Example 9: using strings.Split in for range (can use SplitSeq in Go 1.24)
	text := "apple,banana,cherry"
	for _, fruit := range strings.Split(text, ",") {
		fmt.Println("Fruit:", fruit)
	}

	// Example 10: map copying (can use maps.Clone in Go 1.21)
	originalMap := map[string]int{"one": 1, "two": 2, "three": 3}
	newMap := make(map[string]int)
	for k, v := range originalMap {
		newMap[k] = v
	}
	fmt.Println("Copied map:", newMap)
}
