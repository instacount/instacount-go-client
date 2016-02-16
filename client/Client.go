package instacount

import (
	"time"
	"encoding/json"
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

func NewClient(appID string, apiKey string) *Client {
	client := new(Client)
	client.transport = NewTransport(appID, apiKey)
	return client
}

func NewClientWithHosts(appID string, apiKey string, hosts []string) *Client {
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

func (c *Client) InitShardedCounter(counterName string) *ShardedCounter {
	return NewShardedCounter(c, counterName)
}

func (c *Client) GetShardedCounter(counterName string) (*ShardedCounter, *Errors, error) {
	if r, e, err := c.transport.request("GET", "/" + sharded_counters + "/" + c.transport.urlEncode(counterName), nil, read); err != nil {
		return nil, nil, err
	} else if e != nil {
		return nil, e, nil
	} else {
		// Try to unmarshal a ShardedCounter
		sc := toShardedCounter(r)
		return sc, e, nil
	}
}

func (c *Client) IncrementShardedCounter(counterName string, async bool) (interface{}, *Errors, error) {
	params := &Increment{
		amount:  1,
		async: async,
	}
	return c.IncrementShardedCounterWithParams(counterName, params)
}

func (c *Client) IncrementShardedCounterWithParams(counterName string, params interface{}) (interface{}, *Errors, error) {
	return c.transport.request("POST", "/" + sharded_counters + "/" + c.transport.urlEncode(counterName) + "/" + increments, params, write)
}

func (c *Client) DecrementShardedCounter(counterName string, async bool) (interface{}, *Errors, error) {
	params := &Increment{
		amount:  1,
		async: async,
	}
	return c.DecrementShardedCounterWithParams(counterName, params)
}

func (c *Client) DecrementShardedCounterWithParams(counterName string, params interface{}) (interface{}, *Errors, error) {
	return c.transport.request("POST", "/" + sharded_counters + "/" + c.transport.urlEncode(counterName) + "/" + decrements, params, write)
}

// Helper to convert a map[string][]json.RawMessage containing ShardedCounter JSON into the ShardedCounter struct.
func toShardedCounter(rawJson interface{}) (*ShardedCounter) {

	// Cast to the proper object
	var rm = rawJson.(json.RawMessage)

	sc := ShardedCounter{}
	if err := json.Unmarshal(rm, &sc); err != nil {
		panic(err)
	}

	return &sc
}