package main

import (
	"fmt"
)

// each structure is independent
// keep words simple
type Dog struct {
    Breed string
    Weight int
}

func main() {
    poodle := Dog{"Poodle", 34}
    fmt.Println(poodle)
    fmt.Printf("%+v\n", poodle)
    fmt.Printf("Breed: %v\n Weight: %v\n", poodle.Breed, poodle.Weight)
}
