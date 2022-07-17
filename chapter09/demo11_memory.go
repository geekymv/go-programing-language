package main

import (
	"image"
	"sync"
)

// 通道通信 、互斥锁 都会导致处理器把本地缓存中的数据刷回内存并提交。
// 可见性、有序性
// 互斥锁保证可见性和有序（可以防止指令重排序）

// CPU缓存

// 指令重排序，CPU 和编译器优化指令执行次序。

// make 不是原子操作，包含多个步骤
/*
icons = make(map[string]image.Image)

1.分配一块内存 M
2.在内存 M 上初始化 map 对象
3.然后 M 的地址赋值给变量 icons

指令重排的结果可能是
1.分配一块内存 M
2.M 的地址赋值给变量 icons
3.在内存 M 上初始化 map 对象

*/

var rw sync.RWMutex
var icons map[string]image.Image

func Icon(name string) image.Image {
	// 获取读锁（共享锁）
	rw.RLock()
	if icons != nil {
		icon := icons[name]
		rw.RUnlock()
		return icon
	}
	// 释放读锁
	rw.RUnlock()

	rw.Lock()
	defer rw.Unlock()

	if icons == nil {
		// loadIcons 延迟加载，只需要加载一次
		loadIcons()
	}
	return icons[name]
}

func loadIcons() {
	icons = map[string]image.Image{
		// TODO
	}
}
