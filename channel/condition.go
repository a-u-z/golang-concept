package channel

import (
	"log"
	"time"
)

// 如果要判斷一個非同步是否完成，在完成的時候要在 chan 塞入一個值，然後如果可以取出
// 那就代表是真的完成那個非同步，就可以往下執行
// 因為 chan 的特性是從 chan 取不出東西會一直 hang 在那邊

var msg string
var done bool

func setup() {
	msg = "hello world"
	time.Sleep(time.Second * 2)
	done = true
}
func empty() {
	msg = " "
}
func Run() {
	go setup()
	go empty()
	for !done {

	}
	log.Printf(msg)
}

var newDone = make(chan struct{})

func newSetup() {
	msg = "hello world"
	newDone <- struct{}{} // 不佔記憶體資源
}
func NewRun() {
	go newSetup()
	go empty()
	<-newDone
	log.Printf("here is msg:%+v", msg)
}
