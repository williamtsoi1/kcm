package main

import (
	"context"
	"fmt"
	"log"

	ce "github.com/cloudevents/sdk-go"
	"github.com/tidwall/gjson"
)

var (
	textPath = mustGetEnv("TEXT_PATH", "")
)

type eventReceiver struct{}

func (r *eventReceiver) Receive(ctx context.Context, event ce.Event, resp *ce.EventResponse) error {

	log.Printf("Raw Event: %v", event)

	// get content
	var textValue string
	if event.DataContentType() == "text/plain" {
		if err := event.DataAs(textValue); err != nil {
			log.Printf("Failed to DataAs string: %s", err.Error())
			return err
		}
	} else if event.DataContentType() == "application/json" {
		content, err := event.DataBytes()
		if err != nil {
			log.Printf("Failed to DataAs bytes: %s", err.Error())
			return err
		}
		textValue = gjson.GetBytes(content, textPath).String()
	} else {
		return fmt.Errorf("Invalid Data Content Type: %s. Only application/json and text/plain supported",
			event.DataContentType())
	}

	log.Printf("Text to score: %s", textValue)

	// empty handler
	if textValue == "" {
		resp = &ce.EventResponse{
			Status:  200,
			Event:   &event,
			Reason:  "Text not found",
			Context: ctx,
		}
		return nil
	}

	//score content
	m, s, e := scoreSentiment(ctx, textValue)
	if e != nil {
		log.Printf("Failed on score sentiment: %s", e.Error())
		return e
	}

	sr := map[string]float32{
		"magnitude": m,
		"score":     s,
	}
	log.Printf("Score: %v", sr)

	event.SetType(fmt.Sprintf("%s.scored", event.Type()))
	event.SetExtension("sentiment", sr)

	log.Printf("Processed Event: %v", event)

	resp = &ce.EventResponse{
		Status:  200,
		Event:   &event,
		Reason:  "Scored",
		Context: ctx,
	}

	return nil

}
