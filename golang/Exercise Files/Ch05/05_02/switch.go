package main

import (
	"fmt"
	"math/rand"
	"time"	
) 

func main() {

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	result := ""

	switch dow {
	    case 1:
	        result = "It's Sunday!"
	    case 7
	        result = "It's Saturday!"
	    default:
	        result = "It's a weekday!"
	}

	fmt.Println ("Day", dow, ",", result)

}
