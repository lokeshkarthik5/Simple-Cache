package shard

import "hash/fnv"

func Hash(key string) uint64 {
	h := fnv.New64()
	h.Write([]byte(key))
	return h.Sum64()
}
