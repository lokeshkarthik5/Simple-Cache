package cache

import (
	"github.com/lokeshkarthik5/simple-cache/internal/shard"
)

type Cache struct {
	shards    []*shard.Shard
	numShards uint64
}
