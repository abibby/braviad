package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/abibby/braviad/config"
	"github.com/abibby/braviad/tv"
	winlog "github.com/ofcoursedude/gowinlog"
)

type Event struct {
	Data []EventData `xml:"EventData>Data"`
}
type EventData struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:",innerxml"`
}

func (e *Event) Get(name string) (string, bool) {
	for _, d := range e.Data {
		if d.Name == name {
			return d.Value, true
		}
	}
	return "", false
}

func main() {
	err := config.Load("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tv.BraviaConfig{}
	err = config.Get(cfg)
	if err != nil {
		log.Fatal(err)
	}
	var t tv.TV = tv.NewBravia(cfg)

	err = t.Activate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting...")
	watcher, err := winlog.NewWinLogWatcher()
	if err != nil {
		fmt.Printf("Couldn't create watcher: %v\n", err)
		return
	}
	// Recieve any future messages on the Application channel
	// "*" doesn't filter by any fields of the event
	err = watcher.SubscribeFromNow("System", "*")
	if err != nil {
		fmt.Printf("Couldn't subscrbe to events: %v\n", err)
		return
	}
	log.Print("Waiting for events")
	for {
		select {
		case evt := <-watcher.Event():
			if evt.EventId == 566 {
				data := &Event{}
				err := xml.Unmarshal([]byte(evt.Xml), data)
				if err != nil {
					log.Print(err)
					continue
				}
				nextSessionType, ok := data.Get("NextSessionType")
				if !ok {
					log.Print("no next session")
					continue
				}
				previousSessionType, ok := data.Get("PreviousSessionType")
				if !ok {
					log.Print("no previous session")
					continue
				}
				log.Print(previousSessionType, " -> ", nextSessionType)
				switch nextSessionType {
				case "1": // Screensaver
					t.DisplayOff()
				case "2": // ?

				case "3": // Sleep
					t.PowerOff()
				case "0":
					t.Activate()
				}
			} else if evt.EventId == 1 {
				t.PowerOff()
			}

		case err := <-watcher.Error():
			log.Printf("\nError: %v\n\n", err)
		default:
			// If no event is waiting, need to wait or do something else, otherwise
			// the the app fails on deadlock.
			<-time.After(1 * time.Millisecond)
		}
	}
}
