package IRC

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
)

func connect() net.Conn {
	conn, err := net.Dial("tcp", "irc.root-me.org")
	if err != nil {
		panic(err)
	}
	return conn
}

func disconnect(conn net.Conn) {
	conn.Close()
}

func readmsg(conn net.Conn) {
	tp := textproto.NewReader(bufio.NewReader(conn))
	for {
		status, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Println(status)
	}
}

func login(conn net.Conn) {
	fmt.Fprintf(conn, "USER Sifouo 8 * :Someone")
	fmt.Fprintf(conn, "NICK SayedMahdi")
}

func sendData(conn net.Conn, message string) {
	fmt.Fprintf(conn, "%s\r\n", message)
}
