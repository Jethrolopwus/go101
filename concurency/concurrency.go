

package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond) 
	}
}

func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("Letter: %c\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers() 
	go printLetters() 

	time.Sleep(3 * time.Second)
	fmt.Println("Main function ends")
}
