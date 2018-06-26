package lib

import (
	"bufio"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Listener int

const bufferSize = 512

func (l *Listener) Proc(data *Package, ack *bool) error {

	file := getFile(data.Header.Filename)
	writer := bufio.NewWriterSize(file, bufferSize)
	writeBytes, ew := writer.Write(data.Body)
	CheckErr(ew)
	fmt.Println("正在接收第", data.Header.PackageIndex, "/", data.Header.PackageCount, "个包,", "包大小:", len(data.Body), "读取大小:", writeBytes, "文件名:", data.Header.Filename)
	writer.Flush()
	*ack = true
	defer file.Close()
	return nil
}

func Receive() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+Port)
	CheckErr(err)

	inbound, errListenTCP := net.ListenTCP("tcp", addy)
	CheckErr(errListenTCP)

	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}

func getFile(fileName string) *os.File {
	if _, err := os.Stat("./downloads"); os.IsNotExist(err) { // 如果downloads目录不存在则创建
		os.Mkdir("./downloads", 0700)
	}
	if _, err := os.Stat("./downloads/" + fileName); os.IsNotExist(err) { // 如果downloads目录不存在则创建
		_, err1 := os.Create("./downloads/" + fileName) //创建文件
		CheckErr(err1)
	}
	file, errF := os.OpenFile("./downloads/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	CheckErr(errF)
	return file
}
