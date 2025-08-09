package controls

import "github.com/marc-antoinegelinas/feishin-controls/internal/websocket"

type SimpleEvent string

const (
	Next     SimpleEvent = "next"
	Pause    SimpleEvent = "pause"
	Play     SimpleEvent = "play"
	Previous SimpleEvent = "previous"
	Proxy    SimpleEvent = "proxy"
	Repeat   SimpleEvent = "repeat"
	Shuffle  SimpleEvent = "shuffle"
)

func ClientSimpleEvent(event SimpleEvent) {
	request := map[string]string{
		"event": string(event),
	}
	websocket.SendRequest(request)
}
