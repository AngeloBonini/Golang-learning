package main

import "fmt"

var messages = []string{
	"I am the one who knocks",
	"THe Horse is gotta bust",
	"You can be jealus at me",
	"a cup of tea make everything simpler",
	"our mind blind us until we make ourselves truthly free",
	"everyone assume many things about the world",
}

func producer(Link chan<- string) {
	for _, m := range messages {
		Link <- m
	}
	close(Link)
}

func consumer(Link <-chan string, done chan<- bool) {
	for b := range Link {
		fmt.Println(b)
	}
	done <- true
}

func main() {
	Link := make(chan string)
	done := make(chan bool)
	go producer(Link)
	go consumer(Link, done)
	<-done
}
