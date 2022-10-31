package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/santiago-simplisafe/redis-lock/internal/lock"
)

var wg sync.WaitGroup

func main() {
	clientsPtr := flag.Int("clients", 1, "Number of clients")
	commandsPtr := flag.Int("commands", 10, "Number of commands per client")
	delayPtr := flag.Int("delay", 250, "Delay between commands in milliseconds")
	hostAddrPtr := flag.String("host", "localhost", "Redis host address")
	modePtr := flag.String("mode", "single", "single or cluster")

	flag.Parse()

	redisLock := lock.NewRedisLock(*hostAddrPtr, *modePtr)

	start := time.Now()

	for i := 0; i < *clientsPtr; i++ {
		clientUUID := uuid.New().String()
		for j := 0; j < *commandsPtr; j++ {
			wg.Add(1)
			go sendCommand(redisLock, clientUUID, j, time.Duration(*delayPtr))
		}
	}

	wg.Wait()
	fmt.Println("All commands sent in", time.Since(start))

}

func sendCommand(redisLock *lock.RedisLock, clientUUID string, index int, delay time.Duration) {

	defer wg.Done()

	// Send command to basestation
	for {
		if lockAndSend(redisLock, clientUUID, index) {
			break
		}

		time.Sleep(delay * time.Millisecond)
		// fmt.Println("Retrying:", clientUUID, "index:", index)
	}

}

func lockAndSend(redisLock *lock.RedisLock, clientUUID string, index int) bool {
	lock, _ := redisLock.Aquire(clientUUID, strconv.Itoa(index), 60*time.Second)

	if !lock.IsLocked {
		// fmt.Println("🔒 Lock not acquired. Index:", index)
		return false
	}

	fmt.Println(time.Now().Format("15:04:05.000"), "✅ Lock acquired.   ", clientUUID, "index:", index)

	rand.Seed(time.Now().UnixNano())
	// Simulate sending command to basestation. Random wait between 2 and 5 seconds
	seconds := rand.Intn(4) + 2
	fmt.Println(time.Now().Format("15:04:05.000"), "🚀 Sending command. ", clientUUID, "index:", index, "wait:", seconds, "seconds")
	time.Sleep(time.Duration(seconds) * time.Second)

	// Release lock
	redisLock.Release(lock)
	fmt.Println(time.Now().Format("15:04:05.000"), "🎉 Lock released.   ", clientUUID, "index:", index)
	return true
}
