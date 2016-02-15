package instacount
import (
	"time"
)

const sharded_counters = "sharded_counters"
const increments = "increments"
const decrements = "decrements"

type ShardedCounter struct {
	client      *Client `json:"-"`
	Name        string `json:"name"`
	nameEncoded string `json:"-"`
	Description string  `json:"description"`
	NumShards   int `json:"numShards"`
	Status      string `json:"status"`
	Created     time.Time `json:"createdDateTime"`
	Count       int64 `json:"count"`
}

func NewShardedCounter(client *Client, name string) *ShardedCounter {
	shardedCounter := new(ShardedCounter)
	shardedCounter.client = client
	shardedCounter.Name = name
	shardedCounter.nameEncoded = client.transport.urlEncode(name)
	return shardedCounter
}

func (c *ShardedCounter) GetCounter(counterId string) (interface{}, error) {
	return c.client.transport.request("GET", "/" + sharded_counters + "/" + c.client.transport.urlEncode(counterId), nil, read)
}

func (c *ShardedCounter) IncrementShardedCounter(counterName string, async bool) (interface{}, error) {
	params := &Increment{
		amount:  1,
		async: async,
	}
	return c.IncrementShardedCounterWithParams(counterName, params)
}

func (c *ShardedCounter) IncrementShardedCounterWithParams(counterName string, params interface{}) (interface{}, error) {
	return c.client.transport.request("POST", "/" + sharded_counters + "/" + c.client.transport.urlEncode(counterName) + "/" + increments, params, write)
}


func (c *ShardedCounter) DecrementShardedCounter(counterName string, async bool) (interface{}, error) {
	params := &Increment{
		amount:  1,
		async: async,
	}
	return c.DecrementShardedCounterWithParams(counterName, params)
}

func (c *ShardedCounter) DecrementShardedCounterWithParams(counterName string, params interface{}) (interface{}, error) {
	return c.client.transport.request("POST", "/" + sharded_counters + "/" + c.client.transport.urlEncode(counterName) + "/" + decrements, params, write)
}