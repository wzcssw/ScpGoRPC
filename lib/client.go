package lib

import (
	"fmt"
	"math"
	"net/rpc"
	"os"
	"time"
)

const Port string = "42586"

const PackageSize int64 = 1024 * 1024 * 5.0

// const PackageSize int64 = 256

func Send(addr string, fileName string) {
	file, fileSize := openFile(fileName)
	defer file.Close()

	fmt.Println("----- 分包数量:", generatePackage(fileSize), "-----")
	file.Seek(0, 0)
	packageCount := generatePackage(fileSize)
	var cur_offset int64 //  偏移量
	for i := 0; i < packageCount; i++ {
		pkgSize := PackageSize
		if i+1 == packageCount { // 最后一个包发送剩余的包大小,其余的发送PackageSize大小
			pkgSize = fileSize - cur_offset
		}
		buf := make([]byte, pkgSize)
		n, errRead := file.Read(buf)
		cur_offset, _ = file.Seek(0, os.SEEK_CUR)
		CheckErr(errRead)
		fmt.Println("正在发送第", i+1, "/", packageCount, "个包,", "包大小:", pkgSize, "预定义包大小:", PackageSize, "读取大小:", n, "偏移量:", cur_offset, "预定义包大小:", PackageSize)
		Package := &Package{}
		Package.Header = Header{Filename: file.Name(), Filesize: fileSize, PackageIndex: i + 1, PackageCount: packageCount, PackageSize: pkgSize}
		Package.Body = buf
		callRPC(addr, Package)
	}

}

// return (File,FileSize)
func openFile(fileName string) (*os.File, int64) {
	file, _ := os.Open(fileName)
	fi, errStat := file.Stat()
	CheckErr(errStat)
	return file, fi.Size()
}

//  generatePackage 文件分包数量
func generatePackage(fileSize int64) int {
	return int(math.Ceil(float64(fileSize) / float64(PackageSize)))
}

func callRPC(addr string, Package *Package) {
	client, err := rpc.Dial("tcp", addr+":"+Port)
	CheckErr(err)
	var reply bool
	times := 4
	for i := 0; i < times; i++ {
		if i > 0 {
			time.Sleep(time.Duration(Fibonacci(i) * 1e9))
			fmt.Println("**** 重试第", i, "次****  包:", Package.Header)
		}
		errCall := client.Call("Listener.Proc", Package, &reply)
		if errCall != nil || !reply {
			if i == 3 {
				panic("上传失败!")
			}
		} else {
			break
		}
	}

}
