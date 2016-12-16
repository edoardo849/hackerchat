package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/JSainsburyPLC/manchester-hackathon/xmas-2016/go-chat/chat"
	"github.com/JSainsburyPLC/manchester-hackathon/xmas-2016/go-chat/cli"
	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
)

func main() {
	// open a file
	f, err := os.OpenFile("./_logs/go-chat.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// don't forget to close it
	defer f.Close()

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(f)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	ctx := context.Background()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	kc, err := chat.NewKafkaClient([]string{"52.214.102.133:9092"}, config)
	if err != nil {
		log.WithError(err).Fatal("error connecting to kafka")
	}
	client := chat.BasicClientWrapper(kc, 100)
	//receiveYou, err := client.Consumer("you")
	//if err != nil {
	//log.WithError(err).Fatal("failed to set up consumer")
	//}
	//sendMe, err := client.Producer("me")
	//if err != nil {
	//log.WithError(err).Fatal("failed to set up producer")
	//}
	//

	stop, err := cli.Start(ctx, os.Stdout, os.Stdin, client)
	if err != nil {
		panic(err)
	}
	defer stop()

	//cleanMock := make(chan struct{})
	//go func() {
	//log.Debug("Mock started.")
	//for {
	//
	//select {
	//
	//case <-cleanMock:
	//return
	//
	//case msg := <-receiveYou:
	//log.Debugf("Received: %s", msg)
	//time.Sleep(1 * time.Second)
	//log.Debug("Sending BACK")
	//sendMe <- chat.Message{
	//Author:    "auto-reply",
	//Body:      fmt.Sprintf("Hey, received: %s", msg.Body),
	//Timestamp: time.Now(),
	//}
	//log.Debug("SENT")
	//
	//}
	//}
	//
	//}()
	<-c
	//	cleanMock <- struct{}{}
	os.Exit(0)
	fmt.Println("Done!")

}