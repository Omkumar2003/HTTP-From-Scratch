package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {

		defer f.Close()
		// ob := make([]byte, 8)
		ob := make([]byte, 1)
		// fmt.Print("read :")
		// ch <- "read :"
		str := ""
		for {
			_, err := f.Read(ob)
			if err != nil {
				ch <- str
				break
			}
			if string(ob) == string('\n') {
				// fmt.Println("")
				str += string(ob)
				ch <- str
				str = ""
				// ch <- "read :"
				continue
			} else {
				str += string(ob)
			}

			// fmt.Print(string(ob[:ok]))
			// ch <- string(ob)
		}
		close(ch)
	}()


	return ch
}

func main() {

	f, err := os.Open("messages.txt")
	if err != nil {
		log.Panic(err)
	}

	chg := getLinesChannel(f)
	for m := range chg {
		fmt.Print("read :" + m)
	}
}
