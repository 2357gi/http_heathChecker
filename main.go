package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const (
		totalExistNum = 6
		maxNum        = 3
	)
	sig := make(chan string, totalExistNum)
	res := make(chan string, maxNum)
	defer close(sig)
	defer close(res)
	fmt.Printf("start concurrency execute %s \n", time.Now())
	gw := new(sync.WaitGroup)
	gw.Add(totalExistNum)
	go countUp()
	for i := 0; i < totalExistNum; i++ {
		go w86Sec(sig, res, i, gw)
	}
	gw.Wait()
	fmt.Printf("exitconcurrency execute %s \n", time.Now())

}

func w86Sec(sig chan string, res chan string, name int, wg *sync.WaitGroup) {
	sig <- fmt.Sprintf("sig %d", name)
	t := 6 + name
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Printf("no.%d: end wait %v sec.\n", name, t)
	wg.Done()
}

func countUp() {
	i := 0
	for {
		fmt.Println(i)
		i++
		time.Sleep(time.Duration(i) + time.Second)
	}
}
