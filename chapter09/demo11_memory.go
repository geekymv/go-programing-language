package main

// CPU缓存

// 指令重排序，CPU 和编译器优化指令执行次序。

// make 不是原子操作
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
