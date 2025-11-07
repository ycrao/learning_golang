package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	fmt.Println("示例 1：使用缓冲 channel 作为信号量限制并发")
	semaphoreExample(20, 5, rng)

	fmt.Println("\n示例 2：使用 token-channel（token bucket）做速率限流") // 消费限流
	rateLimiterExample(10, 2, 3)                              // total=10, rate=2 tokens/sec, burst=3
}

// 使用缓冲 channel 做信号量，限制同时运行的 goroutine 数量
func semaphoreExample(totalTasks, maxConcurrency int, rng *rand.Rand) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrency) // 容量 = 最大并发数

	for i := 0; i < totalTasks; i++ {
		wg.Add(1)
		sem <- struct{}{} // acquire token (若已满会阻塞，限制并发)
		go func(id int) {
			defer wg.Done()
			defer func() { <-sem }() // release token

			// 模拟工作
			d := time.Duration(rng.Intn(400)+100) * time.Millisecond
			fmt.Printf("[task %2d] start at %v (will take %v)\n", id, time.Now().Format("15:04:05.000"), d)
			time.Sleep(d)
			fmt.Printf("[task %2d] done at %v\n", id, time.Now().Format("15:04:05.000"))
		}(i)
	}

	wg.Wait()
	fmt.Println("semaphoreExample 完成")
}

// token bucket：用带缓冲的 channel 保存 token，ticker 周期性填充 token
// rate: 每秒产生 token 的速率（tokens/sec），burst: token 的最大突发容量
func rateLimiterExample(totalJobs, rate, burst int) {
	if rate <= 0 || burst <= 0 {
		fmt.Println("rate and burst must be > 0")
		return
	}

	tokens := make(chan struct{}, burst)

	// 预先填充突发容量
	for i := 0; i < burst; i++ {
		tokens <- struct{}{}
	}

	// 以固定速率补充 token
	interval := time.Second / time.Duration(rate)
	ticker := time.NewTicker(interval)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				select {
				case tokens <- struct{}{}:
				default:
					// 达到 burst 上限则丢弃多余 token
				}
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	var wg sync.WaitGroup
	for j := 0; j < totalJobs; j++ {
		wg.Add(1)

		// 请求 token（若没有 token 则在这里阻塞，达到流控效果）
		go func(id int) {
			defer wg.Done()
			<-tokens // consume one token

			// 处理请求
			fmt.Printf("[job %2d] allowed at %v\n", id, time.Now().Format("15:04:05.000"))
			// 模拟短任务
			time.Sleep(150 * time.Millisecond)
		}(j)
	}

	wg.Wait()
	// 停止补 token goroutine
	close(done)
	fmt.Println("rateLimiterExample 完成")
}
