package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	// the first option, using environment variable
	fmt.Println(os.Getenv("POD"))
	// the second option, using volume
	bytes, err := ioutil.ReadFile("/etc/pod")
	if err == nil {
		fmt.Println(string(bytes))
	}
	// donot exit the loop
	for {
		time.Sleep(time.Duration(1) * time.Minute)
	}
}
