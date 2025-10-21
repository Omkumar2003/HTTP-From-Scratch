package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Panic(err)
	}

	defer f.Close()
	// ob := make([]byte, 8)
	ob := make([]byte, 1)
	fmt.Print("read :")

	for {
		ok, err := f.Read(ob)
		if err != nil {
			break
		}
		if string(ob) == string('\n') {
			fmt.Println("")
			fmt.Print("read :")
			continue
		}
		fmt.Print(string(ob[:ok]))
	}
}
