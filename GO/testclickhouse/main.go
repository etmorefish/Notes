package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
1 读取文件、解压 读取文件到内存
2 入库 一次性写入 分多线程测试 入库总量时间
3 每次入库的条数 500、1000 结合线程测试
4 查询测试，首先入库10天到一个月的数据
*/
func main() {
	var path = "D:\\dl\\2022-09-08"
	if IsDir(path) != true {
		return
	}
	defer TimeCost(time.Now())

	//items := Walk(path)
	//fmt.Println(items)
	//newItems := UngzAll(items)
	lineCh := make(chan string, 10000)
	demoItems := []string{"D:\\dl\\2022-09-08\\JB05.txt", "D:\\dl\\2022-09-08\\JB06.txt", "D:\\dl\\2022-09-08\\JB07.txt"}
	go func() {
		ReadRecord(demoItems[0], lineCh)
	}()
	time.Sleep(time.Second * 3)
	//close(lineCh)
	fmt.Println(len(lineCh))
	for i := 0; i < 4; i++ {
		line := <-lineCh
		fmt.Println(line)
	}
}

func TimeCost(start time.Time) {
	// 统计程序运行时间
	tc := time.Since(start)
	fmt.Printf("time cost = %v\n", tc)
}

func Walk(path string) (items []string) {
	// 遍历目录下(包含子目录)所有文件及目录
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		filePath := filepath.Join(path, file.Name())
		items = append(items, filePath)
	}
	return items
}

func WalkDir(path string) (items []string) {

	// 获取当前目录下的所有文件或目录信息
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		fmt.Printf(" %s\n", path)
		fmt.Println(path) // 打印path信息   (带绝对路径)
		//fmt.Println(info.Name()) // 打印文件或目录名(无绝对路径)
		items = append(items, path)
		return nil
	})
	return items
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func UnGzip(filePath string) (content []byte, err error) {
	gzipFile, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer gzipFile.Close()
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		return
	}
	defer gzipReader.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, gzipReader)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func UngzAll(items []string) (newItems []string) {
	// 将所有.gz 文件解压至当前目录
	for _, item := range items {
		content, err := UnGzip(item)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// fmt.Println(string(content))
		fileName := fmt.Sprintf("%s.txt", strings.Split(item, ".")[0])
		//fmt.Println(fileName)
		newItems = append(newItems, fileName)
		//写入文件
		if err = ioutil.WriteFile(fileName, content, 0666); err != nil {
			fmt.Println("写入错误：", err)
		}
	}
	return newItems
}

func ReadRecord(filename string, lineCh chan string) {
	log.Println(filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Println(filename + " error")
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text() // line就是每行文本
		// 对line进行处理
		lineCh <- line
	}
}
