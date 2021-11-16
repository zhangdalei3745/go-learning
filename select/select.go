package main

import (
	"fmt"
	"time"
)

// select是Go中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
// select是一个随机执行的可运行的case。如果没有case执行，他将阻塞，直到有case执行。一个默认的select子句应该总是可执行的。
func main() {
	// 1、超时控制
	timeControl()
}

func foo() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		for {
			select {
			case data, ok := <-chanInt:
				if ok {
					fmt.Println(data)
				}
				break
			default:
				fmt.Println("全部阻塞")
			}

		}
	}()
	//time.Sleep(time.Second)
	chanInt <- 1
}

func timeControl() {
	// 信号量为2
	stream := make(chan int, 2)
	defer close(stream)
	w := NewWorker(stream, time.Second*10)
	w.Start()
}

// 定义任务
type Worker struct {
	stream  chan int      //可读写chan
	timeout time.Duration //超时
	done    chan struct{} //结束信号
}

// 定义初始化函数
func NewWorker(stream chan int, timeout time.Duration) *Worker {
	return &Worker{
		stream:  stream,
		timeout: timeout,
		done:    make(chan struct{}),
	}
}

// 定义超时控制
func (w *Worker) afterTimeStop() {
	// 经过超时时间后发送结束信号
	go func() {
		time.Sleep(w.timeout)
		// 超时后发送结束信号
		w.done <- struct{}{}
	}()
	// 可读写chan，每隔2秒发送一条消息
	for i := 0; i < 50; i++ {
		if i > 0 && i%2 == 0 {
			time.Sleep(time.Second)
		}
		w.stream <- i
	}
}

//
func (w *Worker) Start() {
	go w.afterTimeStop()
	for {
		select {
		case data, ok := <-w.stream:
			if !ok {
				return
			}
			fmt.Println(fmt.Sprintf(time.Unix(time.Now().Unix(), 0).String()))
			fmt.Println("收到数据：%s", data)
		case <-w.done:
			close(w.done)
			return
		}
	}
}
