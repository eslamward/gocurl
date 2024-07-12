package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	if len(os.Args) < 2 {
		log.Fatal("please enter valid url")
	}
	if len(os.Args) > 2 {
		numberOfRequest, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err.Error())
		}
		wg.Add(numberOfRequest)
		for i := 0; i < numberOfRequest; i++ {
			go Request(os.Args[1], &wg)
		}

	} else {
		wg.Add(1)
		go Request(os.Args[1], &wg)
	}

	wg.Wait()
}

func Request(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())

	}

	defer res.Body.Close()

	fmt.Printf(
		"url:%s - status:%s - %d - time:%v\n",
		res.Request.Method,
		res.Status,
		res.StatusCode,
		time.Since(start).Round(time.Millisecond))

}
