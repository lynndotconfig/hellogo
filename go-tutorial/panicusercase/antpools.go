package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"time"
)

func main() {
	go task(1)
	go task(2)
	go task(3)
	time.Sleep(time.Second * time.Duration(30))}

func InnerFunc(timeout int) {
	for i:=0; i < 100; i+=1 {
		if i == timeout {
			time.Sleep(time.Second * time.Duration(i))
		}
		if timeout == 5 {
			panic(fmt.Sprintf("timeout painc=%d\n", timeout))
		}
	}
	fmt.Printf("timeout=%d, finished\n", timeout)
}


func task(idx int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("catch panic for task, idx=%s, err=%s", idx, err)
			if err, ok := err.(error); ok {
				fmt.Printf("reconver for task=%d, err=%s\n", idx, err)
			}
		}
		fmt.Printf("finished task=%d\n", idx)
	}()
	var wg sync.WaitGroup
	pool, err := ants.NewPoolWithFunc(10, func(param interface{}) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("catch panic for pool, param=%s, ere%s", param, err)
				if err, ok := err.(error); ok {
					fmt.Printf("reconver for pool, params=%s, err=%s", param, err)
				}
			}
			fmt.Printf("finished pool,parma=%s\n", param)
		}()
		defer wg.Done()
		timeout := param.(int)
		InnerFunc(timeout)
	})
	defer pool.Release()
	if err != nil {
		fmt.Printf("init pool err=%s\n", err)
		return
	}
	for i := 0; i < 20; i += 1 {
		wg.Add(1)
		invokeErr := pool.Invoke(i)
		if invokeErr != nil {
			fmt.Printf("invoke err=%s\n", err)
			return
		}
	}
	wg.Wait()
}
