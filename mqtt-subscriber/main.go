package main

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func main() {

    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }


    var broker = os.Getenv("MQTT_URL")
    var port = 8883
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
    opts.SetUsername(os.Getenv("MQTT_USERNAME"))
    opts.SetPassword(os.Getenv("MQTT_KEY"))

    client := mqtt.NewClient(opts)
    
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
    
    topic := "home"

    if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
        fmt.Println(token.Error())
    }

    client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
        fmt.Printf("Recieved message: %s from topic: %s\n", msg.Payload(), msg.Topic())
    })

}
