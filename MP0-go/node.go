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
		 str := strings.Split(scanner.Text(), " ")
		 output := str[0] + " " + node + " " + str[1] + "\n"
		_, err = conn.Write([]byte(output))
		//fmt.Fprintf(conn, stdin)
	}
	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//_, err = conn.Write([]byte(stdin))
	//if err != nil{
	//	return
	//}
	//buf := [512]byte{}
	//n, err := conn.Read(buf[:])
	//fmt.Fprintf(conn, stdin)
	//if err != nil{
	//	fmt.Println("receive failed, err:", err)
	//	return
	//}
	//fmt.Println(string(buf[:n]))
}
