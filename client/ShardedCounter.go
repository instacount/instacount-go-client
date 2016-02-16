package instacount
import "time"

const sharded_counters = "sharded_counters"
const increments = "increments"
const decrements = "decrements"

type Href struct {
	Href string `json:"href"`
}

type ShardedCounterMeta struct {
	SelfLink       Href `json:"@self"`
	IncrementsLink Href  `json:"@increments"`
	DecrementsLink Href `json:"@decrements"`
}

type ShardedCounter struct {
	Meta        ShardedCounterMeta `json:"meta"`
	Name        string `json:"name"`
	Description string  `json:"description"`
	NumShards   int `json:"numShards"`
	Status      string `json:"status"`
	Created     time.Time `json:"createdDateTime"`
	Count       int64 `json:"count"`
}

// Helper to construct a new ShardedCounter
func NewShardedCounter(client *Client, name string) *ShardedCounter {
	shardedCounter := new(ShardedCounter)
	shardedCounter.Name = name
	return shardedCounter
}

type ShardedCounterOperation struct {
	client       *Client `json:"-"`
	Name         string `json:"name"`
	AppliedCount int64
}