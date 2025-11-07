# chan（通道）

最近面试有问到如何理解 `golang` 的 `channel`，几年前看过一些，现在早忘记了，于是回顾一下。

>   以下内容由 `AI` 补充的。

### 什么是 channel

- `channel` 是 `Go` 提供的在 `goroutine` 之间传递数据的管道，类型为 `chan T`，`T` 为传输的值类型。
- `channel` 本身是引用类型，需要使用 `make` 创建（类似 `slice`、`map`）。

### 创建与基本操作

- 创建：`ch := make(chan int)          // 无缓冲`
- 带缓冲：`ch := make(chan int, 5)    // 最大缓存 5 个元素`
- 发送：`ch <- v`
- 接收：`v := <-ch`
- 接收并判断是否关闭：`v, ok := <-ch`

### 缓冲与同步

- 无缓冲 `channel` 在发送和接收双方就绪时完成数据传递（同步点）。
- 带缓冲可以在缓冲未满时立即发送，接收者再延后取数据（异步场景）。

### 关闭 channel

- `close(ch)` 用于通知接收方不再有新的值发送。
- 只能由发送方关闭 `channel`；对已关闭 `channel` 继续发送会导致 `panic`。
- `(v, ok := <-ch)` 从已关闭的 `channel` 接收会得到类型零值，并且 `ok` 返回 `false`。

### 遍历 channel

- 使用 for range 遍历 channel，直到被关闭：

  ```go
  for v := range ch {
      // 处理 v
  }
  ```

### 单向 channel

- 可指定方向：`chan<- int`（只能发送），`<-chan int`（只能接收），用于接口或函数参数以提高安全性。

### select 多路复用

- select 可以同时等待多个 channel 操作，常用于超时、合并多路输入或优先级控制：
  ```go
  select {
  case v := <-ch1:
      // ...
  case ch2 <- x:
      // ...
  case <-time.After(time.Second):
      // 超时处理
  }
  ```

### 常见应用场景

`worker pool`、`pipeline`、任务调度、事件广播（结合多个 `channel` 或 `fan-out/fan-in` 模式）。

### 常见坑与建议

- 避免在多个 `goroutine` 中无序地关闭同一 `channel`（可能导致 `panic`）。
- 关闭 `channel` 用于通知接收方，不要把关闭作为常规同步手段。
- `nil channel` 会永远阻塞，可用于动态启用/禁用 `select` 分支。
- 注意死锁：当所有 `goroutine` 都在等待发送或接收时会发生死锁，程序会 `panic`。

### 简短示例

```go
// 简单管道与工作者示例
func main() {
    jobs := make(chan int, 3)  // 创建一个可以缓冲3个整数的通道
    done := make(chan struct{})  // 创建一个用于信号传递的通道，struct{} 类型不占用空间
    // jobs 通道用于主 goroutine 向工作 goroutine 发送任务（整数）。
    // done 通道用于工作 goroutine 在完成任务后向主 goroutine 发送完成信号。

    go func() {
        for j := range jobs {  // 从jobs通道中接收任务，直到通道关闭
            fmt.Println("处理任务", j)  // 模拟处理任务
        }
        done <- struct{}{}  // 通过 done 通道发送完成信号
    }()  // 注意这里的()表示立即调用匿名函数，启动 goroutine

    // 发送任务并关闭
    for i := 0; i < 3; i++ {  // 发送3个任务
        jobs <- i  // 将任务（整数）发送到 jobs 通道
    }
    close(jobs)  // 关闭 jobs 通道，表示所有任务已发送完毕

    <-done  // 阻塞等待，直到从done通道接收到完成信号

    // 主 goroutine 在发送完所有任务后，会等待从 done 通道接收完成信号，确保工作 goroutine 完成任务后再继续执行。
}
```

### 小结

`channel` 是 `goroutine` 间通信和同步的重要原语。理解缓冲、关闭、`select` 与单向类型，有助于写出正确且高效的并发代码。

