package main

import (
	"fmt"
	"sync"
)

var (
	sum     int
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

func main() {
	//c := make(chan int)
	//for i := 0; i < 5; i++ {
	//	worker := &Worker{id: i}
	//	go worker.process(c)
	//}
	//
	//for {
	//	c <- rand.Int()
	//	time.Sleep(time.Millisecond * 5000)
	//}
	//run()
	//doOnce()
	pointer()
}

func pointer() {
	name := "Hello Golang"
	nameP := &name

	fmt.Println("name变量的值为:", name)
	fmt.Println("name变量的内存地址为:", nameP)
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
func run() {
	var wg sync.WaitGroup
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			add(10)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("i 的 值 为：%d ；和为：%d\n", i, readSum())
		}()
	}
	wg.Wait()
}
func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}
func readSum() int {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	b := sum
	return b
}

//type Worker struct {
//	id int
//}
//
//func (w *Worker) process(c chan int) {
//	for {
//		data := <-c
//		fmt.Printf("worker %d got %d\n", w.id, data)
//	}
//}
