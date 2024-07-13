package curl

import (
	"log"
	"strconv"
	"sync"
)

func ServeCommand(args []string, flags Flags) {
	var wg sync.WaitGroup
	if len(args) < 2 {
		log.Fatal("please enter valid url")
	}

	nOfRequest := numberOfRequest(args)
	if nOfRequest > 1 {

		wg.Add(nOfRequest)
		for i := 0; i < nOfRequest; i++ {
			if flags.SaveInFile {
				go RequestMany(args[3], &wg, flags)

			} else {
				go RequestMany(args[1], &wg, flags)

			}
		}

		wg.Wait()

	} else {
		if flags.SaveInFile {
			RequestOnce(args[3], flags)

		} else {
			RequestOnce(args[1], flags)

		}
	}
}

func numberOfRequest(args []string) int {

	for _, val := range args {
		numberOfRequest, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		return numberOfRequest
	}
	return 1
}
