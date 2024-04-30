package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}

	// MQTT connection options
	opts := mqtt.NewClientOptions().
		AddBroker(os.Getenv("MQTT_URL")).
		SetUsername(os.Getenv("MQTT_USERNAME")).
		SetPassword(os.Getenv("MQTT_KEY"))

	// Create MQTT client
	client := mqtt.NewClient(opts)

	// Connect to the MQTT broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %s", token.Error())
	}
	defer client.Disconnect(250)

	// Subscribe to a topic
	topic := "home"
	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic %s: %s", topic, token.Error())
	}

	// Handle graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Keep listening for messages until a termination signal is received
	for {
		select {
		case <-signalChan:
			log.Println("Shutting down...")
			return // Exit the loop and terminate the program
		}
	}
}

// messageHandler is the callback function for incoming messages
func messageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
