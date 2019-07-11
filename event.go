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

	// Magnitude indicates the overall strength of emotion
	// (both positive and negative) between 0.0 and +inf.
	// Sample
	// Clearly Positive*	"score": 0.8, "magnitude": 3.0
	// Clearly Negative*	"score": -0.6, "magnitude": 4.0
	// Neutral				"score": 0.1, "magnitude": 0.0
	// Mixed				"score": 0.0, "magnitude": 4.0
	minMagnitude = toFloat(mustGetEnv("MIN_MAGNITUDE", "0.9"))
)

type eventReceiver struct{}

func (r *eventReceiver) Receive(ctx context.Context, event ce.Event, resp *ce.EventResponse) error {

	//log.Printf("Raw Event: %v", event)

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
		return nil
	}

	//score content
	m, s, e := scoreSentiment(ctx, textValue)
	if e != nil {
		log.Printf("Failed on score sentiment: %s", e.Error())
		return e
	}
	log.Printf("Score: %f Magnitude: %f (min %f))", s, m, minMagnitude)

	// set the extension
	event.SetExtension("sentiment", map[string]float32{
		"score":     s,
		"magnitude": m,
	})

	if s < 0 && m >= minMagnitude {
		// negative
		event.SetType(fmt.Sprintf("%s.negative", event.Type()))
	} else if s > 0 && m >= minMagnitude {
		// positive
		event.SetType(fmt.Sprintf("%s.positive", event.Type()))
	} else {
		log.Printf("Sentiment of not significant enough magnitude: %f (expected %f)",
			m, minMagnitude)
		return nil
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
