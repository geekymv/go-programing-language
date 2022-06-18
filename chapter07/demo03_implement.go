package main

type AA interface {
	hello()
	play()
}

type BB struct {
}

func (b *BB) hello() {

}

func (b *BB) play() {
}

// 让BB具体类型必须实现AA接口中的所有方法
var _ AA = (*BB)(nil)
