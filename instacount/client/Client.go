package instacount

import (
	"time"
)

type Client struct {
	transport *Transport
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

func (c *Client) GetShardedCounter(counterName string) (interface{}, error) {
	return c.transport.request("GET", "/sharded_counters/" + counterName, nil, read)
}

func (c *Client) IncrementShardedCounter(counterName string, async bool) (interface{}, error) {
	params := `{"amount": 1, "async": }`
	params[async] = async
	return c.IncrementShardedCounterWithParams(counterName, params)
}

func (c *Client) IncrementShardedCounterWithParams(counterName string, params interface{}) (interface{}, error) {
	return c.transport.request("POST", "/sharded_counters/" + counterName + "/increments", params, write)
}

func (c *Client) DecrementShardedCounter(counterName string, async bool) (interface{}, error) {
	params := `{"amount": 1, "async": false}`
	params[async] = async
	return c.DecrementShardedCounterWithParams(counterName, params)
}

func (c *Client) DecrementShardedCounterWithParams(counterName string, params interface{}) (interface{}, error) {
	return c.transport.request("POST", "/sharded_counters/" + counterName + "/decrements", params, write)
}

func (c *Client) EncodeParams(body interface{}) string {
	return c.transport.EncodeParams(body)
}