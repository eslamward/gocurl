package curl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func request(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())

	}

	return res

}

func RequestOnce(url string, flags Flags) {

	res := request(url)
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	switch true {
	case flags.SaveInFile:
		SaveInFile(os.Args[2], bytes)
	case !flags.SaveInFile:
		fmt.Println(string(bytes))
	}
}

func RequestMany(url string, wg *sync.WaitGroup, flags Flags) {
	start := time.Now()
	defer wg.Done()
	res := request(url)
	defer res.Body.Close()

	switch true {
	case flags.SaveInFile:
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		SaveInFile(os.Args[2], bytes)
	case !flags.SaveInFile:
		fmt.Printf(
			"url:%s - status:%s - %d - time:%v\n",
			res.Request.Method,
			res.Status,
			res.StatusCode,
			time.Since(start).Round(time.Millisecond))
	}

}
