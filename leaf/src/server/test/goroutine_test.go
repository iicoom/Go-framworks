package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}

