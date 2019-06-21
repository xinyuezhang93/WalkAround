// channel_example.go
package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		messages <- "done"
	}()

	msg := <-messages

	fmt.Println(msg)
}
