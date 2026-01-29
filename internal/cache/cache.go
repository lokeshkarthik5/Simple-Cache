package cache

import (
	"time"

	"github.com/lokeshkarthik5/simple-cache/internal/shard"
)

type Cache struct {
	shards    []*shard.Shard
	numShards uint64
}

func New(numShards, capacity int) *Cache {

	shards := make([]*shard.Shard, numShards)
	for i := 0; i < numShards; i++ {
		shards[i] = shard.New(capacity)
	}
	return &Cache{
		shards:    shards,
		numShards: uint64(numShards),
	}
}

func (c *Cache) pickShard(key string) *shard.Shard {
	idx := shard.Hash(key) % c.numShards
	return c.shards[idx]
}

func (c *Cache) Get(key string) (any, error) {
	return c.pickShard(key).Get(key)
}

func (c *Cache) Set(key string, val any, ttl time.Duration) {
	c.pickShard(key).Set(key, val, ttl)
}

func (c *Cache) Delete(key string) error {
	return c.pickShard(key).Delete(key)
}
