package curl

import (
	"fmt"
	"os"
)

func SaveInFile(fileName string, data []byte) {

	err := os.WriteFile(fileName, data, 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
