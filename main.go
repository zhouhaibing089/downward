package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// the first option, using environment variable
	fmt.Println(os.Getenv("POD_NAME"))
	// the second option, using volume
	file, err := os.Open("/etc/pod")
	if err == nil {
		fmt.Println(bufio.NewScanner(file).Text())
		file.Close()
	}
	// donot exit the loop
	for {
		time.Sleep(time.Duration(1) * time.Minute)
	}
}
