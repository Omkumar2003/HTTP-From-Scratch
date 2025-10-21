package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// go run . | Tee-Object -FilePath "$env:TEMP\tcp.txt"

// $client = New-Object System.Net.Sockets.TcpClient("127.0.0.1", 42069)
// $stream = $client.GetStream()
// $writer = New-Object System.IO.StreamWriter($stream)
// $writer.AutoFlush = $true
// $writer.WriteLine("Do you have what it takes to be an engineer at TheStartup™?")
// $writer.Close()
// $client.Close()

func getLinesChannel(f io.ReadCloser) <-chan string {
	str := ""
	buf := make([]byte, 1)
	ch := make(chan string)
	go func() {
		defer f.Close()
		defer close(ch)
		defer fmt.Println("connection has been closed")
		for {
			_, err := f.Read(buf)
			if err != nil {
				ch <- str
				// fmt.Print(str)
				return
			}
			if string(buf) == "\n" {
				// fmt.Print(str)
				ch <- str
				str = ""
			} else {
				str += string(buf)
			}
		}

	}()

	return ch
}

func main() {
	l, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Panic(err)
	}
	defer l.Close()

	for {
		con, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("connection has been accepted")
		om := getLinesChannel(con)
		for temp := range om {
			fmt.Println(temp)
		}

	}
}
