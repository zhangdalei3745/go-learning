package main

import (
	"fmt"
	"time"
)

// 1、通道一定要收发，不然会阻塞，产生死锁。
// 2、有缓存通道和无缓存通道，有缓存通道发送后就会阻塞。
// 3、有缓存通道，缓存满后无通道读时会阻塞。
// 4、通道包含单双通道和多向通道
func main() {
	c := make(chan int, 1)
	done := make(chan struct{})
	defer close(done)
	go send(c)
	go receive(c)
	go end(done)
	<-done
	return
}

//只能向chan里写数据
func send(c chan int) {
	for i := 0; i < 20; i++ {
		fmt.Println("通道写取数据：", i, "写入时间：", fmt.Sprintf(time.Unix(time.Now().Unix(), 0).String()))
		c <- i
	}
}

//只能取channel中的数据
func receive(c chan int) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		data := <-c
		fmt.Println("通道读取数据：", data, "读取时间：", fmt.Sprintf(time.Unix(time.Now().Unix(), 0).String()))
	}
}

func end(done chan struct{}) {
	time.Sleep(time.Second * 60)
	done <- struct{}{}
}
