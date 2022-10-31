# Redis distributed lock POC

Based on the [Redis](https://redis.io/docs/reference/patterns/distributed-locks/) distributed lock algorithm.

This POC is a simple implementation of the redis lock algorithm that is used to lock the access to a resource
based on a key. The lock is configured per client, that means each client can have a different lock
resource. The retry mechanism is an infinite loop that will try to acquire the lock until it succeeds. There is a 
delay between each retry, it's configured to 250ms by default. There is no timeout applied to the retry mechanism.

The script create as many clients as you want, each client will try to acquire the lock and will
release it after a 2-5 seconds interval. The script will stop when all the clients have released the lock.

## Limitations

This version is running on a single node Redis instance. Currently working on the cluster version.


## Usage

### Running with a local GO installation

- First run the local Redis instance with `docker-compose up -d`
- Install go modules with `go mod download`
- Then run the application with `go run command.go`

```bash
  docker-compose up -d
  go mod download
  go run command.go --clients=4 --comands=2 --delay=250 --host=docker.for.mac.localhost
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