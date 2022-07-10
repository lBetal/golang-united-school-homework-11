package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	goroutineLimiter := make(chan struct{}, pool)
	var wg sync.WaitGroup
	defer wg.Done()
	var mu sync.Mutex
	var i int64
	for i = 0; i < n; i++ {
		goroutineLimiter <- struct{}{}
		wg.Add(1)
		go func(next int64) {
			nextUser := getOne(next)
			mu.Lock()
			res = append(res, nextUser)
			mu.Unlock()
			<-goroutineLimiter
		}(i)
	}
	wg.Wait()
	return res
}
