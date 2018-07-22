package concurrency

import (
"fmt"
)

func Entry()  {
	fmt.Println("----- Wellcome to concurrency -----")
	done := make(chan  bool)
	go testConcurrent1(done)
	<-done //读取数据从通道
	fmt.Println("done all deal...")


	//定向数据发送/接收通道
	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)


	//关闭通道上的循环范围
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	fmt.Println("----- End concurrency -----")
}

func testConcurrent1(done chan bool)  {
	for i:= 0; i<10; i++{
		fmt.Println("index =", i);
	}
	done <- true;//写入数据到通道
}

func sendData(sendch chan<- int) {
	sendch <- 10
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}