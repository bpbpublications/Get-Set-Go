package main

import (
    "fmt"

    "github.com/IBM/sarama"
)

func main() {
    topic := "coffee_orders"

    // Create a new consumer and start it.
    config := sarama.NewConfig()
    config.Consumer.Return.Errors = true

    brokers := []string{"localhost:9092"}
    worker, err := sarama.NewConsumer(brokers, config)
    if err != nil {
        panic(err)
    }

    consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
    if err != nil {
        panic(err)
    }

    defer worker.Close()

    fmt.Println("Consumer started ")

    // Run the consumer / worker.
    for {
        select {
        case err := <-consumer.Errors():
            fmt.Println(err)
        case msg := <-consumer.Messages():
            fmt.Printf("Received order: Topic(%s) | Message(%s) \n", string(msg.Topic), string(msg.Value))
            order := string(msg.Value)
            fmt.Printf("Brewing coffee for order: %s\n", order)
        }
    }
}
