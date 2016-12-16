package chat_test

import (
	"testing"
	"time"

	"github.com/JSainsburyPLC/manchester-hackathon/xmas-2016/go-chat/chat"
)

func TestWappedClient(t *testing.T) {
	mock := &chat.MockClient{
		SendFunc: func(dest string, msg chat.Message) error {
			return nil
		},
		ReceiveFunc: func(src string) (chat.Message, error) {
			return chat.Message{}, nil
		},
	}
	client := chat.BasicClientWrapper(mock, 100)

	message := chat.Message{
		Body:      "this is a test",
		Author:    "go devs",
		Timestamp: time.Now(),
	}
	t.Run("send", func(t *testing.T) {
		done := make(chan interface{})
		mock.SendFunc = func(dest string, received chat.Message) error {
			defer close(done)
			if received.Body != message.Body {
				t.Errorf("expected Body: %s, got: %s", message.Body, received.Body)
			}

			if received.Author != message.Author {
				t.Errorf("expected Author: %s, got: %s", message.Author, received.Author)
			}

			if received.Timestamp != message.Timestamp {
				t.Errorf("expected Timestamp: %s, got: %s", message.Timestamp, received.Timestamp)
			}
			return nil
		}
		send, err := client.Producer("test")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
			return
		}
		//defer close(send)
		send <- message
		select {
		case <-time.After(10 * time.Second):
			t.Errorf("timeout")
		case <-done:
		}

	})
	t.Run("receive", func(t *testing.T) {
		mock.ReceiveFunc = func(src string) (chat.Message, error) {
			return message, nil
		}
		receive, err := client.Consumer("test")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
			return
		}
		//defer close(receive)
		select {
		case <-time.After(10 * time.Second):
			t.Errorf("timeout")
		case received := <-receive:
			if received.Body != message.Body {
				t.Errorf("expected Body: %s, got: %s", message.Body, received.Body)
			}

			if received.Author != message.Author {
				t.Errorf("expected Author: %s, got: %s", message.Author, received.Author)
			}

			if received.Timestamp != message.Timestamp {
				t.Errorf("expected Timestamp: %s, got: %s", message.Timestamp, received.Timestamp)
			}
		}
	})
}
