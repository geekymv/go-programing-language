package main

// b = b + a 不是原子操作

// 数据竞态 发生于付哦个goroutine 并发读写同一个变量，并且其中一个是写入。

// 不要通过共享内存来通信，而应该通过通信来共享内存。

// 互斥机制：保证同一时刻只有一个 goroutine 执行。