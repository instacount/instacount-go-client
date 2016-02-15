package instacount

import (
	"time"
)

type Client struct {
	transport *Transport
}

type Increment struct {
	amount int64
	async  bool
}

type Decrement struct {
	amount int64
	async  bool
}

func NewClient(appID, apiKey string) *Client {
	client := new(Client)
	client.transport = NewTransport(appID, apiKey)
	return client
}

func NewClientWithHosts(appID, apiKey string, hosts []string) *Client {
	client := new(Client)
	client.transport = NewTransportWithHosts(appID, apiKey, hosts)
	return client
}

func (c *Client) SetExtraHeader(key string, value string) {
	c.transport.setExtraHeader(key, value)
}

func (c *Client) SetTimeout(connectTimeout int, readTimeout int) {
	c.transport.setTimeout(time.Duration(connectTimeout) * time.Millisecond, time.Duration(readTimeout) * time.Millisecond)
}


func (c *Client) EncodeParams(body interface{}) string {
	return c.transport.EncodeParams(body)
}