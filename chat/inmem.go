package chat

type memClient struct {
	chanBuff int
	channels map[string]chan Message
}

// NewMemClient initalises new in memory client
func NewMemClient(chanBuffer int) BasicClient {
	return &memClient{
		chanBuff: chanBuffer,
		channels: make(map[string]chan Message),
	}
}

func (client *memClient) channel(name string) chan Message {
	ch, found := client.channels[name]
	if !found {
		ch = make(chan Message, client.chanBuff)
		client.channels[name] = ch
	}
	return ch
}

// Send message to dest channel
func (client *memClient) Send(dest string, msg Message) error {
	// TODO timeout
	client.channel(dest) <- msg
	return nil
}

// Receive gets message from src chan
func (client *memClient) Receive(src string) (Message, error) {
	// TODO timeout
	msg := <-client.channel(src)
	return msg, nil
}
