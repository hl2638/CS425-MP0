package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()
	firstMsg := true
	var nodeName string
	reader := bufio.NewReader(conn)
	for {
		recvTime := float64(time.Now().UnixNano())/1000000000
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			//if err == io.EOF{
			//	fmt.Fprintln(os.Stderr, "EOF!")
			//}
			//fmt.Fprintln(os.Stderr, "read from client failed, err: ", err)
			fmt.Printf("%.7f - %s disconnected\n", recvTime, nodeName)
			break
		}
		recvStr := string(buf[:n])
		if firstMsg{
			firstMsg = false
			strFields := strings.Split(recvStr, " ")
			nodeName = strFields[1]
			//fmt.Printf("%.7f - %s connected\n", recvTime, nodeName)
			fmt.Printf("%s - %s connected\n", strFields[0], nodeName)
		}

		fmt.Print(recvStr)
		//conn.Write([]byte(recvStr))
	}
}
//var connMap map[string]bool
func main() {
	var port string
	if len(os.Args) > 1{
		port = os.Args[1]
	}else{
		port = "9999"
	}
	fmt.Println(port)
	listen, err := net.Listen("tcp", "0.0.0.0:" + port)
	//connMap = make(map[string]bool)
	fmt.Println("Server started.")
	if err != nil {
		fmt.Fprintln(os.Stderr,"listen() failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Accept() failed, err", err)
			continue
		}
		go process(conn)
	}
}