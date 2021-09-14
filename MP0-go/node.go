package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	node := "node"
	address := "127.0.0.1:9999"
	if len(args) > 0{
		node = args[0]
	}
	if len(args) > 1{
		address = args[1] + ":" + args[2]
	}
	//fmt.Println(args)
	fmt.Println("Client " + node + " started.")
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer conn.Close()
	//defer fmt.Println("Close");
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//recvTime := float64(time.Now().UnixNano())/1000000000
		 str := strings.Split(scanner.Text(), " ")
		 output := str[0] + " " + node + " " + str[1] + "\n"
		_, err = conn.Write([]byte(output))
		//fmt.Fprintf(conn, stdin)
		//fmt.Printf("At time %.7f, To send: %s", recvTime, output)
	}
}
