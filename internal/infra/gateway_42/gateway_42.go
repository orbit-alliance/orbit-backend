package gateway_42

import (
	"sync"
	"time"
)

type Gateway42 struct {
	client *Client
	mu     sync.Mutex
	token  token42
}

func NewGateway42() *Gateway42 {
	return &Gateway42{
		client: NewClient(10*time.Second, 120*time.Second),
	}
}
