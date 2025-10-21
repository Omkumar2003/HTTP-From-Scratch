package main

import (
	"bufio"
	"fmt"

	"log"
	"net"
	"os"
)

// for windows to listen udp
// $udp=New-Object System.Net.Sockets.UdpClient 42069;while($true){$r=New-Object System.Net.IPEndPoint([System.Net.IPAddress]::Any,0);$b=$udp.Receive([ref]$r);$s=[System.Text.Encoding]::UTF8.GetString($b);Write-Host "$($r.Address):$($r.Port) -> $s"}

// udp does not even care if the listner is listening or not it just sends

func main() {
	raddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Panic(err)
	}
	con, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Panic(err)
	}
	defer con.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		lol, err := reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		_, err = con.Write([]byte(lol))
		if err != nil {
			log.Print(err)
		}
	}
}
