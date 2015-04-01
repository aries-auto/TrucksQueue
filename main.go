package main

import (
	"github.com/aries-auto/TrucksQueue/handlers"
	"github.com/bitly/go-nsq"

	"log"
	"sync"
)

var (
	NSQDHosts = []string{
		"127.0.0.1:4161",
	}

	ConsumerConcurrency = 100
)

func main() {

	wg := &sync.WaitGroup{}

	config := nsq.NewConfig()
	applications, err := nsq.NewConsumer("applications", "ch", config)
	if err != nil {
		log.Fatal(err.Error())
	}

	applications.AddConcurrentHandlers(nsq.HandlerFunc(handlers.HandleMessage), ConsumerConcurrency)

	running := 0

	err = applications.ConnectToNSQLookupds(NSQDHosts)
	if err == nil {
		running = running + 1
		wg.Add(1)
	}
	if running == 0 {
		wg.Done()
		return
	}
	wg.Wait()

}
