package utils

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// MessageHandlers handles messages sent to the server
func MessageHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message from topic %s: %s", msg.Topic(), msg.Payload())
}

// ConnectionHandler handles established connection between client/server and broker
func ConnectionHandler(client MQTT.Client) {
	fmt.Printf("Connected\n")
}

// ConnectionLostHandler handles lost connection between client/server and broker
func ConnectionLostHandler(client MQTT.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

// OnMessageReceived starts a goroutine for every received message
func OnMessageReceived(client MQTT.Client, message MQTT.Message) {
	go MessageHandler(client, message)
}

// SetClientOptions sets client options from command line arguments
func SetClientOptions(brokerURL string, username string, password string, clientID string) *MQTT.ClientOptions {
	options := MQTT.NewClientOptions()

	options.AddBroker(brokerURL)
	options.SetUsername(username)
	options.SetPassword(password)
	options.SetClientID(clientID)
	options.SetConnectionLostHandler(ConnectionLostHandler)
	options.SetOnConnectHandler(ConnectionHandler)

	return options
}

// InitiateConnection creates client with given parameters and connects to MQTT broker
func InitiateConnection(brokerURL string, username string, password string, clientID string) MQTT.Client {
	options := SetClientOptions(brokerURL, username, password, clientID)
	client := MQTT.NewClient(options)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}
