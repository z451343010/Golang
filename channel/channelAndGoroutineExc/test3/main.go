/*
	作业2：goroutine + channel 配合完成排序，并写入文件
	要求：
	1）开一个协程 writeDataToFile，随机生成1000个数据，存放到文件中
	2）当writeDataToFile 完成写1000个数据到文件后，
	让sort协程从文件中读取1000个文件，
	并完成排序，重新写入到另外一个文件
	3）考察点：协程和管道+文件的综合使用
	4）功能扩展：开10个协程 writeDataToFile,
	每个协程随机生成1000个数据，存放到10文件中
	5）当10个文件都生成了，让10个sort协程从10文件中读取1000个文件，
	并完成排序，重新写入到10个结果文件
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// 写文件
func writeDataToFile(loopNum int, filePath string, numChan chan int) {

	rand.Seed(time.Now().UnixNano())

	var count int

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}

	writer := bufio.NewWriter(file)

	for i := 0; i < loopNum; i++ {

		// 统计执行写入的次数
		count++

		randNum := rand.Intn(1000)
		// intChan <- randNum
		writeStr := fmt.Sprintf("%d\n", randNum)
		writer.WriteString(writeStr)

		numChan <- randNum
		fmt.Println("write working write number ", randNum)

	}

	writer.Flush()
	defer file.Close()

	if count == loopNum {
		close(numChan)
	}

}

// 读文件
func sort(readFilePath string, writeFilePath string,
	exitReadChan chan bool, numChan chan int,
	numSlice []int, loopNum int) {

	readFile, readErr := os.Open(readFilePath)
	if readErr != nil {
		fmt.Println("open readFile err = ", readErr)
	}

	reader := bufio.NewReader(readFile)

	for {

		_, ok := <-numChan

		if !ok {
			break
		}

		str, readErr := reader.ReadString('\n')
		if readErr == io.EOF {
			break
		}

		rs := []rune(str)
		str = string(rs[0 : len(str)-1])

		readNum, transErr := strconv.Atoi(str)

		if transErr != nil {
			fmt.Println("translate to int err = ", transErr)
			break
		}

		fmt.Println("read working read number", readNum)

		numSlice = append(numSlice, readNum)

		if len(numSlice) == loopNum {

			// 对切片进行冒泡排序
			for i := len(numSlice) - 1; i >= 0; i-- {

				for j := 0; j < i; j++ {

					if numSlice[j] > numSlice[j+1] {

						var temp int
						temp = numSlice[j]
						numSlice[j] = numSlice[j+1]
						numSlice[j+1] = temp

					}

				}

			}

			// 将排序后的切片数据存入文件
			writeFile, writeErr := os.OpenFile(writeFilePath, os.O_WRONLY|os.O_CREATE, 0666)

			if writeErr != nil {
				fmt.Println("open file writeErr = ", writeErr)
				return
			}

			writer := bufio.NewWriter(writeFile)

			for i := 0; i < len(numSlice); i++ {

				writeStr := fmt.Sprintf("%d\n", numSlice[i])
				writer.WriteString(writeStr)

			}

			writer.Flush()
			defer writeFile.Close()

			exitReadChan <- true
			close(exitReadChan)

		}

		defer readFile.Close()

	}

	// exitReadChan <- true
	// close(exitReadChan)

}

func main() {

	numChan := make(chan int, 100)
	exitReadChan := make(chan bool, 1)
	var loopNum int = 100
	var numSlice []int

	var readFilePath string = "F:\\Golang\\goProgramProject" +
		"\\src\\channel\\channelAndGoroutineExc\\test3\\file\\read.txt"

	var writeFilePath string = "F:\\Golang\\goProgramProject" +
		"\\src\\channel\\channelAndGoroutineExc\\test3\\file\\write.txt"

	go writeDataToFile(loopNum, readFilePath, numChan)
	go sort(readFilePath, writeFilePath,
		exitReadChan, numChan, numSlice, loopNum)

	for {

		_, ok := <-exitReadChan

		if !ok {
			break
		}

	}

}
