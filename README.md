# Sharded In-Memory Cache in Go

A **high-performance, sharded in-memory cache** written in Go, featuring:

- LRU eviction
- TTL-based expiration
- Shard-level locking for concurrency
- Clean separation between cache core and storage shards
- CLI interface for testing (`set`, `get`, `delete`)

This project focuses on **systems-level design**, correctness, and concurrency rather than frameworks.

---

## ✨ Features

- **Sharded architecture**  
  Cache keys are hashed across multiple shards to reduce lock contention.

- **LRU eviction (O(1))**  
  Implemented using a map + doubly linked list.

- **TTL support**  
  Entries expire automatically on access.

- **Thread-safe**  
  Each shard owns its own mutex.

- **CLI-driven**  
  Simple command-line interface for validating behavior before adding HTTP or middleware layers.

---


### Responsibilities

- **Cache core**
  - Routes keys to shards
  - Exposes `Get`, `Set`, `Delete`
  - No locking, no data storage

- **Shard**
  - Owns data
  - Owns LRU eviction
  - Owns synchronization
  - Enforces TTL

---

## 🚀 Running the CLI

> Run all commands from the project root.

### Set a value

```bash
go run ./cmd/main.go set foo bar --ttl 5s
```

### Get a value
```bash
go run ./cmd/main.go get foo
```

### Delete a value

```bash 
go run ./cmd/main.go delete foo
```

## NOTE

Still in development and it was more like a fun project to build. Would recommend people to try building one for themself


