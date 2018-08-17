package main

import (
	"time"
	"math/rand"
	"fmt"
)

type result string

// START1 OMIT
func google(query string) (results []result) {
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
	return
}

var (
	web = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
)

type search func(query string) result

func fakeSearch(kind string) search {
	return func(query string) result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := google("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}