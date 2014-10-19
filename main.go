package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	items := flag.Int("items", 65535, "Total number of items the server can store")
	port := flag.String("port", "11212", "The port the server listens on")

	flag.Parse()

	fmt.Println(*port, *items)

	l, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go func(c net.Conn) {
			defer c.Close()
			r := bufio.NewReader(c)
			for {
				line, err := r.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						return
					}
					printErrorTo(c, err)
					continue
				}
				line = strings.TrimSpace(line)
				fmt.Println(line)
			}
		}(conn)
	}
}

func printErrorTo(w io.Writer, err error) {
	w.Write([]byte("ERROR " + err.Error() + "\r\n"))
}
