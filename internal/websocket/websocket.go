package websocket

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"runtime"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

var (
	connection *websocket.Conn
	connOnce   sync.Once
)

func Authenticate() {
	connOnce.Do(func() {
		u := url.URL{Scheme: "ws", Host: viper.GetString("URL"), Path: "/"}

		creds64 := base64.StdEncoding.EncodeToString(fmt.Appendf(nil, "%s:%s", viper.GetString("Username"), viper.GetString("Password")))

		authHeader := map[string]string{
			"event":  "authenticate",
			"header": fmt.Sprintf("Basic %s", creds64),
		}

		var err error

		connection, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Fatal("Dial error:", err)
		}

		err = connection.WriteJSON(authHeader)
		if err != nil {
			log.Fatal("Request error:", err)
		}
	})
}

func GetServerMessage() []byte {
	if connection == nil {
		_, file, no, _ := runtime.Caller(1)
		log.Fatalf("Connection error in %s#%d\n", file, no)
	}

	_, message, err := connection.ReadMessage()
	if err != nil {
		log.Fatal("Read:", err)
	}

	return message
}

func SendRequest(request map[string]string) {
	if connection == nil {
		_, file, no, _ := runtime.Caller(1)
		log.Fatalf("Connection error in %s#%d\n", file, no)
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
