package main

import (
	"fmt"
	"math/rand"
	"time"
)

type result string
type search func(query string) result

var (
	web1   = fakeSearch("web1")
	web2   = fakeSearch("web2")
	image1 = fakeSearch("image1")
	image2 = fakeSearch("image2")
	video1 = fakeSearch("video1")
	video2 = fakeSearch("video2")
)

/*
Google Search 3.0
Reduce tail latency using replicated search servers.

https://talks.golang.org/2012/concurrency.slide#50
*/
func first(query string, replicas ...search) result {
	c := make(chan result)
	searchReplica := func(a int) { c <- replicas[a](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := first("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}

func google(query string) (results []result) {
	c := make(chan result)
	go func() { c <- first(query, web1, web2) }()
	go func() { c <- first(query, image1, image2) }()
	go func() { c <- first(query, video1, video2) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func fakeSearch(kind string) search {
	return func(query string) result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}
