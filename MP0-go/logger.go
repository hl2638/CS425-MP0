package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	//log "github.com/sirupsen/logrus"
	//"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)
type Msg struct{
	TimeStamp float64
	Delay float64
	Size int
}
var statLogger *log.Logger

func processLine(recvTime float64, msg string, p_nodeName *string, p_firstMsg *bool, size int){
	strFields := strings.Split(msg, " ")
	evTime, _ := strconv.ParseFloat(strFields[0], 64)
	if *p_firstMsg{
		*p_firstMsg = false
		*p_nodeName = strFields[1]
		//logFileName = fmt.Sprintf("%s_%d.json", nodeName, int(recvTime))
		//fmt.Printf("%.7f - %s connected\n", recvTime, nodeName)
		//eventLogger.Printf("%s - %s connected\n", strFields[0], nodeName)
		log.Printf("%s - %s connected\n", strFields[0], *p_nodeName)
	}
	//log.Printf("Received %d bytes.\n", n)

	log.Print(msg)

	//fmt.Print(recvStr)
	//fmt.Printf("Receive time  is %.7f\n", recvTime)
	//fmt.Printf("Delay is %.7f\n", recvTime-evTime)
	logMsg := Msg{TimeStamp: recvTime, Delay: recvTime-evTime, Size: size}
	dat, _ := json.Marshal(logMsg)
	statLogger.Println(string(dat))
	//fmt.Println(string(dat))
}

func process(conn net.Conn) {

	defer conn.Close()
	firstMsg := true
	var nodeName string
	//logFileName := "unnamed_log.json"
	//msgArr := make([]Msg, 10)

	reader := bufio.NewReader(conn)
	for {
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		recvTime := float64(time.Now().UnixNano())/1000000000
		if err != nil {
			//if err == io.EOF{
			//	fmt.Fprintln(os.Stderr, "EOF!")
			//}
			//fmt.Fprintln(os.Stderr, "read from client failed, err: ", err)
			log.Printf("%.7f - %s disconnected\n", recvTime, nodeName)
			//fmt.Printf("%.7f - %s disconnected\n", recvTime, nodeName)
			break
		}

		recvStr := string(buf[:n])
		msgArr := strings.Split(strings.Trim(recvStr, "\n"), "\n")

		//log.Println(msgArr)

		for _, msg := range msgArr{
			go processLine(recvTime, msg, &nodeName, &firstMsg, n)
		}


	}

}
//var connMap map[string]bool
func main() {
	var port, logFileName string
	if len(os.Args) > 1{
		port = os.Args[1]
	}else{
		port = "9999"
	}

	if len(os.Args) > 2{
		logFileName = os.Args[2]
	}else{
		logFileName = fmt.Sprintf("%d.json", time.Now())
	}

	log.SetFlags(0)

	listen, err := net.Listen("tcp", "0.0.0.0:" + port)
	log.Println("Server started.")
	//fmt.Println("Server started.")
	if err != nil {
		fmt.Fprintln(os.Stderr,"listen() failed, err:", err)
		return
	}

	file, _ := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	statLogger = log.New(file, "", 0)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Accept() failed, err", err)
			continue
		}
		go process(conn)
	}
}