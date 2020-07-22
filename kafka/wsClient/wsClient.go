package wsClient

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/KumKeeHyun/PDK/kafka/setting"
	"github.com/sacOO7/gowebsocket"
)

func SetupAndStart() *gowebsocket.Socket {
	Repo = NewRegisterRepo()
	socket := gowebsocket.New(setting.WebsocketSetting.URL)

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("Recieved connect error ", err)
	}

	socket.OnTextMessage = MessageHandler

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	socket.Connect()
	return &socket
}

func MessageHandler(message string, socket gowebsocket.Socket) {
	var kafkaMsg KafkaMessage
	if err := json.Unmarshal([]byte(message), &kafkaMsg); err != nil {
		fmt.Printf("json unmarshal fail : %s\n", err.Error())
		return
	}

	jsonBody, err := json.Marshal(kafkaMsg.Msg)
	if err != nil {
		fmt.Printf("json marshal fail : %s\n", err.Error())
	}

	switch kafkaMsg.Type {
	case Init:
		InitMessage(jsonBody)
	case NewNode:
		NewNodeMessage(jsonBody)
	case DeleteNode:
		DeleteNodeMessage(jsonBody)
	case NewSensor:
		NewSensorMessage(jsonBody)
	case DeleteSensor:
		DeleteSensorMessage(jsonBody)
	}
}
