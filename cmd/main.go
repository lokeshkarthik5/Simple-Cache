package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lokeshkarthik5/simple-cache/internal/cache"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: cachecli <set|get|delete>")
		return
	}
	cache := cache.New(32, 100)

	switch os.Args[1] {
	case "set":
		key := os.Args[2]
		val := os.Args[3]
		ttl := 10 * time.Second
		cache.Set(key, val, ttl)
		fmt.Println("OK")

	case "get":
		key := os.Args[2]
		val, err := cache.Get(key)
		if err != nil {
			fmt.Println("MISS")
			return
		}
		fmt.Println(val)

	case "delete":
		key := os.Args[2]
		cache.Delete(key)
		fmt.Println("DELETED")
	}
}
