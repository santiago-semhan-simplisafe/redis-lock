package lock

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisLock struct {
	client *redis.Client
}

type Lock struct {
	IsLocked bool
	key      string
	value    string
	exp      time.Duration
}

func NewRedisLock(client *redis.Client) *RedisLock {
	return &RedisLock{client: client}
}

func (rl *RedisLock) Aquire(key string, value string, exp time.Duration) (*Lock, error) {
	ok, err := rl.client.SetNX(key, value, exp).Result()
	if err != nil {
		return nil, err
	}

	return &Lock{IsLocked: ok, key: key, value: value, exp: exp}, nil

}

func (rl *RedisLock) Release(lock *Lock) error {
	// Check if lock is still held by client
	val, err := rl.client.Get(lock.key).Result()
	if err != nil {
		return err
	}

	if val == lock.value {
		// Delete lock
		_, err := rl.client.Del(lock.key).Result()
		if err != nil {
			return err
		}
	}

	return nil
}
