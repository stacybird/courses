package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	colors = append(colors[:])
	fmt.Println(colors)

	numbers := make ([]int, 5, 5)
	numbers[0] = 134
	numbers[2] = 134
	numbers[0] = 134
	numbers[0] = 134
	numbers[0] = 134


}