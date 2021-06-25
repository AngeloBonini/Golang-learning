package main

import "fmt"

var messages = []string{
	"At one fell swoop",
	"He will give the Devil his due",
	"Butter wouldn't melt in his mouth",
	"Survival of the fittest",
	"Get on my wick",
	"Right Out of the Gate",
	"It's Not All It's Cracked Up To Be",
	"What Goes Up Must Come Down",
	"The short end of the stick",
	"Bodice ripper",
	"The pen is mightier than the sword",
	"Better to light a candle than curse the darkness",
	"Hark, hark! the lark at heaven's gate sings",
	"Alas, poor Yorick! I knew him, Horatio",
	"Stay with me forever",
}

const consumerCount int = 4

func produce(jobs chan<- string) {
	for _, mesg := range messages {
		jobs <- mesg
	}
	close(jobs)
}

func consume(worker int, jobs <-chan string, done chan<- bool) {
	for mesg := range jobs {
		fmt.Printf("Message %v is consumed by worker %v\n", mesg, worker)
	}
	done <- true
}

func main() {
	jobs := make(chan string)
	done := make(chan bool)

	go produce(jobs)
	for iterator := 1; iterator <= consumerCount; iterator++ {
		go consume(iterator, jobs, done)
	}
	<-done
}
