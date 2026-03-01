# Redis in Go

In-memory database inspired by Redis, implemented in Go.

## Build

```bash
make build
```

## Run

```bash
./bin/server --port 6379 --aof
```

### Parameters

- `--port`: Server port (default: 6379)
- `--aof`: Enable disk persistence (Append-Only File)

## Test

```bash
echo "SET name redis-go" | nc localhost 6379
echo "GET name" | nc localhost 6379
```

## Structure

Project follows clean architecture with separation in layers: domain, use cases, adapters, and infrastructure.
