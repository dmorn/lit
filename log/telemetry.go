package log

import (
	"log"
)

type Measurement map[string]int
type Metadata map[string]string

type EventHandler func(string, Measurement, Metadata)

func eventLog(id string, data Measurement, m Metadata) {
	log.Printf("%s: %v %v", id, data, m)
}

func eventDiscard(id string, data Measurement, m Metadata) {}

var handlersLUT = map[string]EventHandler{
	"main": eventLog,

	"lit.SearchLiterature": eventLog,
}

func Event(id string, data Measurement, m Metadata) {
	h, ok := handlersLUT[id]
	if !ok {
		Warn("telemetry handler %q does not exist, yet it is invoked", id)
		return
	}
	h(id, data, m)
}
