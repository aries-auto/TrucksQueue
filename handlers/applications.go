package handlers

import (
	"github.com/aries-auto/trucksplusapi/helpers/reports/aces"
	"github.com/aries-auto/trucksplusapi/models/vehicles"
	"github.com/bitly/go-nsq"
	"log"
)

func HandleMessage(message *nsq.Message) error {

	app, err := aces.DecodeJson(message.Body)
	if err != nil {
		return err
	}
	err = vehicles.PutAppInDatabase(app)
	if err != nil {
		log.Print("IN DB ERR ", err)
		return err
	}

	message.Finish()
	return nil
}
