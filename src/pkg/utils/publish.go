package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Publish publishes message to given topic
func Publish(file string, brokerURL string, username string, password string, clientID string, topic string) {
	client := InitiateConnection(brokerURL, username, password, clientID)

	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	text := fmt.Sprintln("File:", filepath.Base(file)) // strip path from filename and only send the name of the file itself
	token := client.Publish(topic, 0, true, text)
	token.Wait()

	time.Sleep(3 * time.Second)

	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	client.Disconnect(250)
}
