package chat

// MockClient is a mock implementation chat Client
type MockClient struct {
	ReceiveFunc  func(src string) (Message, error)
	SendFunc     func(dest string, msg Message) error
	ConsumerFunc func(src string) (chan Message, error)
	ProducerFunc func(dest string) (chan Message, error)
}

// Receive is a mock implementation
func (mock *MockClient) Receive(src string) (Message, error) {
	return mock.ReceiveFunc(src)
}

// Send is a mock implementation
func (mock *MockClient) Send(dest string, msg Message) error {
	return mock.SendFunc(dest, msg)
}

// Consumer is a mock implementation
func (mock *MockClient) Consumer(src string) (chan Message, error) {
	return mock.ConsumerFunc(src)
}

// Producer is a mock implementation
func (mock *MockClient) Producer(dest string) (chan Message, error) {
	return mock.ProducerFunc(dest)
}
