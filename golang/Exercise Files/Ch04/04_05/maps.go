package main

import (
	"fmt"
	"sort"
)

// store unordered data, process data into an order
func main() {
    states := make(map[string]string)
    fmt.Println(states)

    states["WA"] = "Washington"
    states["CA"] = "California"
    states["OR"] = "Oregon"
    states["UT"] = "Utah"
    fmt.Println(states)

    california := states["CA"]
    fmt.Println(california)

    delete(states, "UT")
    fmt.Println(states)

    for k, v := range states {
        fmt.Printf("%v: %v\n", k, v)
    }

    keys := make([]string, len(states))
    i := 0
    for k := range states {
        keys[i] = k
        i++
    }
    sort.Strings(keys)
    fmt.Println("\nSorted")
    for i := range keys {
        fmt.Println(states[keys[i]])
    }
}