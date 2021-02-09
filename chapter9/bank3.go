package chapter9

import "sync"

var (
	mu       sync.Mutex
	balance3 int
)

func Deposit3(amount int) {
	mu.Lock()
	balance3 = balance3 + amount
	mu.Unlock()
}

func Balance3() int {
	mu.Lock()
	b := balance3
	mu.Unlock()
	return b
}
