package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	workers := runtime.GOMAXPROCS(0)
	fmt.Printf("NUmber of workers:::%v \n", workers)

	trafficController := make(chan int, workers)
	
	var wg sync.WaitGroup
	var counter uint

	for i:=0;i <29;i++ {
		counter += 1
		go PrintCounter(counter, &wg, trafficController)	
	}

	wg.Wait()

}

func PrintCounter(counter uint, wg *sync.WaitGroup, tc chan int) {
	tc <- 1

	wg.Add(1)
	defer wg.Done()

	time.Sleep(time.Second * 4)
	fmt.Println(fmt.Sprintf("COUNTER::%v::", counter))

	<- tc
}
