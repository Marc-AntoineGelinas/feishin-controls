package websocket

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

func authenticate() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: viper.GetString("URL"), Path: "/"}

	creds64 := base64.StdEncoding.EncodeToString(fmt.Appendf(nil, "%s:%s", viper.GetString("Username"), viper.GetString("Password")))

	authHeader := map[string]string{
		"event":  "authenticate",
		"header": fmt.Sprintf("Basic %s", creds64),
	}

	connection, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	err = connection.WriteJSON(authHeader)
	if err != nil {
		log.Fatal("Request error:", err)
	}

	return connection
}

func SendRequest(request map[string]string) {
	connection := authenticate()
	if connection == nil {
		log.Fatal("Connection error:")
	}
	defer func() {
		err := connection.Close()
		if err != nil {
			log.Fatal("Request:", err)
		}
	}()

	err := connection.WriteJSON(request)
	if err != nil {
		log.Fatal("Request:", err)
	}
}
