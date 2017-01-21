package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Printf("%s: start\n", time.Now().String())
	time.Sleep(3 * time.Second)
	fmt.Printf("%s: after sleep\n", time.Now().String())

	wg.Add(1)
	go func() {
		t := time.Now()
		fmt.Printf("t: %s, nano: %d\n", t.String(), t.UnixNano())
		time.Sleep(time.Second)
		t = time.Now()
		fmt.Printf("t: %s, nano: %d\n", t.String(), t.UnixNano())
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("completed\n")
}
