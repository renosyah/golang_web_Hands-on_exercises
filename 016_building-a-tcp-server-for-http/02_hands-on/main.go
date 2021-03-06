package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error serve", err)
		os.Exit(1)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("error at accept", err)
			continue
		}

		go func(c net.Conn) {

			defer c.Close()

			// print request
			request(c)

			// send response
			response(c)

		}(conn)
	}

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			fmt.Println("***METHOD", strings.Fields(ln)[0])
			fmt.Println("***URL", strings.Fields(ln)[1])
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
