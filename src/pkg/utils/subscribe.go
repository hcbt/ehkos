package utils

import "time"

// Subscribe subscribes to a given topic
func Subscribe(brokerURL string, username string, password string, clientID string, topic string) {
	time.Sleep(15 * time.Second) // sleep until MQTT broker starts

	client := InitiateConnection(brokerURL, username, password, clientID)

	if token := client.Subscribe(topic, 0, OnMessageReceived); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Create channel to keep client running
	c := make(chan int)
	<-c

	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Disconnect(250)
}
