package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	// ob := make([]byte, 8)
	for i := 0; ; i++ {
		ob := make([]byte, 8)
		_, err := reader.Read(ob)
		if err != nil {
			break
		}
		fmt.Print(string(ob))

	}

	// fmt.Print(string(ob))
}
