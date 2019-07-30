package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	ce "github.com/cloudevents/sdk-go"
	"github.com/tidwall/gjson"
)

var (
	textPath = mustGetEnv("TEXT_PATH", "text")
)

type eventReceiver struct{}

func (r *eventReceiver) Receive(ctx context.Context, event ce.Event, resp *ce.EventResponse) error {

	//log.Printf("Raw Event: %v", event)

	// get content
	var textValue string

	// check tranlated ext
	err := event.Context.ExtensionAs("translation", &textValue)
	log.Printf("Parsing translation ext: %v", err)

	if err != nil || textValue == "" {
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
	}

	log.Printf("Text to score: %s", textValue)

	// empty handler
	if textValue == "" {
		return nil
	}

	//score content
	m, s, e := scoreSentiment(ctx, textValue)
	if e != nil {
		log.Printf("Failed on score sentiment: %s", e.Error())
		event.SetType(fmt.Sprintf("%s.noneng", event.Type()))
		resp.RespondWith(200, &event)
		return nil
	}
	log.Printf("Score: %f Magnitude: %f)", s, m)

	// set the extension
	event.SetExtension("sentiment", map[string]float32{
		"score":     s,
		"magnitude": m,
	})

	if s < 0 {
		// negative
		event.SetType(fmt.Sprintf("%s.negative", event.Type()))
	} else {
		// positive, ye, really stretching the meaning of the word here
		event.SetType(fmt.Sprintf("%s.positive", event.Type()))
	}

	log.Printf("Classified event: %v", event.Context)
	resp.RespondWith(200, &event)

	return nil

}

func toFloat(s string) float32 {
	f, e := strconv.ParseFloat(s, 32)
	if e != nil {
		log.Fatalf("Value not a float (%s)", s)
	}
	return float32(f)
}
