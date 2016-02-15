package instacount
import (
	"time"
)

const sharded_counters = "sharded_counters"
const increments = "increments"
const decrements = "decrements"

type Counter struct {
	client      *Client
	Name        string `json:"name"`
	Description string  `json:"description"`
	NumShards   int `json:"numShards"`
	Status      string `json:"status"`
	Created     time.Time `json:"createdDateTime"`
	Count       int64 `json:"count"`
}

func (c *Counter) GetCounter(counterId string) (interface{}, error) {
	return c.client.transport.request("GET", "/" + sharded_counters + "/" + c.client.transport.urlEncode(counterId), nil, read)
}

func (c *Counter) IncrementShardedCounter(counterName string, async bool) (interface{}, error) {
	params := &Increment{
		amount:  1,
		async: async,
	}
	return c.IncrementShardedCounterWithParams(counterName, params)
}

func (c *Counter) IncrementShardedCounterWithParams(counterName string, params interface{}) (interface{}, error) {
	return c.client.transport.request("POST", "/" + sharded_counters + "/" + c.client.transport.urlEncode(counterName) + "/" + increments, params, write)
}


func (c *Counter) DecrementShardedCounter(counterName string, async bool) (interface{}, error) {
	params := &Increment{
		amount:  1,
		async: async,
	}
	return c.DecrementShardedCounterWithParams(counterName, params)
}

func (c *Counter) DecrementShardedCounterWithParams(counterName string, params interface{}) (interface{}, error) {
	return c.client.transport.request("POST", "/" + sharded_counters + "/" + c.client.transport.urlEncode(counterName) + "/" + decrements, params, write)
}