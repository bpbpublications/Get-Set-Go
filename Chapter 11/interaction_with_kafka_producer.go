package main

import (
 "encoding/json"
 "log"

 "github.com/IBM/sarama"
)

// Order struct
type Order struct {
 CustomerName string `json:"customer_name"`
 CoffeeType   string `json:"coffee_type"`
}

func PushOrderToQueue(topic string, message []byte) error {
 brokers := []string{"localhost:9092"}
 // Create connection
 producer, err := sarama.NewSyncProducer(brokers, nil)
 if err != nil {
  return err
 }

 defer producer.Close()

 // Create a new message
 msg := &sarama.ProducerMessage{
  Topic: topic,
  Value: sarama.StringEncoder(message),
 }

 // Send message
 partition, offset, err := producer.SendMessage(msg)
 if err != nil {
  return err
 }

 log.Printf("Order is stored in topic(%s)/partition(%d)/offset(%d)\n",
  topic,
  partition,
  offset)

 return nil
}

func main() {
 orders := []Order{
  {CustomerName: "Zorg Blaster", CoffeeType: "Espresso"},
  {CustomerName: "Nova Quasar", CoffeeType: "Latte"},
  {CustomerName: "Orion Vortex", CoffeeType: "Cappuccino"},
  {CustomerName: "Luna Starlight", CoffeeType: "Americano"},
  {CustomerName: "Echo Nebula", CoffeeType: "Mocha"},
 }

 for _, order := range orders {
  orderInBytes, _ := json.Marshal(order)
  PushOrderToQueue("coffee_orders", orderInBytes)
 }
}
