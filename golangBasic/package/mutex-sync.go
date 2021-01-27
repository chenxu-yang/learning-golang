package main

import "sync"

// map是非线程安全的
type info struct {
	mu sync.Mutex
}

func update(info *info) {
	info.mu.Lock()
	//do something
	info.mu.Unlock()
}
