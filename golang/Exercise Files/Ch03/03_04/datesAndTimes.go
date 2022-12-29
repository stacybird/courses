package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("this is a date time %s\n", time1)

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)

    fmt.Println("The month is", time1.Month())
    fmt.Println("The month is", time1.Day())
    fmt.Println("The month is", time1.Weekday())

    tomorrow := time1.AddDate(0, 0, 1)
    fmt.Printf("Tomorrow is %v, %v, %v, %v\n",
        tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

//     own custom model 3:04:05pm in 2006 MST
    longFormat := "Monday, January 2, 2006"
    fmt.Println("The day is", time1.Format(longFormat))

}