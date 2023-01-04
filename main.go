package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}

func main() {
	broker := "tcp://45.32.111.51:1883"
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID("go_mqtt_example")
	options.SetCredentialsProvider(func() (username string, password string) { return "dev", "UATbroker" })
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "en-de/wsm/wsm001/sensor/dht"
	token = client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s\n", topic)

	// key := []byte("the-key-has-to-be-32-bytes-long!")
	// ciphertext := []byte("zfzTvPLBgo3EjH6+S+t5HQ==")
	// painText, err := decrypt(ciphertext, key)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("=> painText: %v\n", painText)

	num := 10
	topicSensor := "wsm/wsm001/sensor/dht"
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("%d", i)
		token = client.Publish(topicSensor, 0, false, text)
		token.Wait()
		time.Sleep(time.Second * 5)
	}
	client.Disconnect(100)

	// EncryptDecrypt()
}
