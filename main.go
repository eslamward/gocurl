package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eslamward/gocurl/curl"
)

func main() {

	saveInFile := flag.Bool("o", false, "Save In File")
	flag.Parse()

	fmt.Println(os.Args)
	curl.ServeCommand(os.Args, *curl.NewFlags(*saveInFile))

}
