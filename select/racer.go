package _select

import (
	"fmt"
	"net/http"
	"time"
)

// 对URL，用time.Now()来记录请求URL前的时间，用http.Get来请求URL的内容，这个函数返回一个http。Response和一个error
// time.since()return the time duration
//func Racer(a, b string) (winner string) {
//	aDuration := measureResponseTime(a)
//	bDuration := measureResponseTime(b)
//
//	if aDuration < bDuration {
//		return a
//	}
//
//	return b
//}
//
//func measureResponseTime(url string) time.Duration {
//	start := time.Now()
//	http.Get(url)
//	return time.Since(start)
//}
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	//select 允许在多个channel等待，第一个发送值的channel胜出，
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
