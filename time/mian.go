package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()

	fmt.Println(currTime)
	fmt.Printf("Type of currTime is %T\n", currTime)

	//for formatting of time, we cant use yyyy-mm-dd, go supports formaating in the format of their launch date(2006-01-02, monday, 03:04:05)
	formatted := currTime.Format("2006-01-02, 03:04:05, Monday")
	fmt.Println("Formatted date is", formatted)

	layout_str := "02/01/2006"

	dateStr := "03/05/2026"
	formattedDate, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted time :", formattedDate)

	newDate := currTime.Add(24 * time.Hour)
	fmt.Println(newDate.Format("02/01/2006, Monday"))
}
