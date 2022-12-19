package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	channel := make(chan string, 102)
	var wg sync.WaitGroup
	counter := 110
	// here we can change the number to increase or decrease speed of checking the  Queue
	//  if we use "1" the Queue will go on with a regular speed , one by one
	for i := 0; i < counter; i++ {
		go receiverChannelAndScrollingSites(channel, &wg)
	}
	senderChannel(channel, &wg)
	wg.Wait()
}

func senderChannel(channel chan<- string, wg *sync.WaitGroup) {
	urls := [2]string{"https://www.google.com", "https://www.varzesh3.com"}
	for i := 0; i < 7; i++ {
		for _, url := range urls {
			wg.Add(1)
			channel <- url
		}
	}
}

func receiverChannelAndScrollingSites(channel <-chan string, wg *sync.WaitGroup) {
	for {
		url := <-channel
		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("%s StatusCode : %d", url, res.StatusCode)
		}
		wg.Done()
	}
}
