package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) { //关键字go,开始新的goroutine,通常使用匿名函数
			//send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		//receive expression
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}

// goroutines是Go的基本并发单位，它让我们可以同时检查多个网站
// channels，用来组织和控制不同进程之间的交流，使我们能避免race condition的问题
//
