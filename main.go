package main

import (
	"fmt"
	"time"
)

// Simulate fetching data from a URL
func fetchData(url string, ch chan<- string) {
    start := time.Now()
    // Simulating delay (replace this line with an actual request if desired)
    time.Sleep(time.Duration(2+time.Now().UnixNano()%3) * time.Second) // Simulate variable delay
    ch <- fmt.Sprintf("Fetched data from %s in %s", url, time.Since(start))
}

func main() {
    urls := []string{
        "http://example.com",
        "http://jsonplaceholder.typicode.com/posts",
        "http://jsonplaceholder.typicode.com/comments",
        "http://jsonplaceholder.typicode.com/users",
    }

    // Channel to receive results
    results := make(chan string)

    // Start fetching data concurrently
    for _, url := range urls {
        go fetchData(url, results)
    }

    // Collect results
    for range urls {
        res := <-results
        fmt.Println(res)
    }
}
