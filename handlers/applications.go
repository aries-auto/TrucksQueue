package handlers

import (
	"github.com/bitly/go-nsq"
	"log"
)

func HandleMessage(message *nsq.Message) error {
	log.Print("msg: ", string(message.Body)) //TODO - put these in MongoDB
	message.Finish()
	return nil
}
