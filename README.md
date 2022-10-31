# Redis distributed lock POC

Based on the [Redis](https://redis.io/docs/reference/patterns/distributed-locks/) distributed lock algorithm.

## Limitations

This version is running on a single node Redis instance. It is not a cluster.


## Usage

### Running with a local GO installation

- First run the local Redis instance with `docker-compose up -d`
- Then run the application with `go run command.go`

```bash
  go mod download
  go run command.go --clients=4 --comands=2 --delay=250
```

### Running with Docker

```bash
  docker-compose up -d
  docker build -t redis-lock .
  docker run -it --rm redis-lock --clients=4 --commands=2 --delay=250 --host=docker.for.mac.localhost
```


### Arguments
  
  ```bash
    -clients int
          Number of clients (default 1)
    -commands int
          Number of commands per client (default 1)
    -delay int
          Delay between commands in milliseconds (default 250)
    -host string
          Redis host (default "localhost")
  ```