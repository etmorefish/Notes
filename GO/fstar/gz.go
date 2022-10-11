package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCsGzip(t *testing.T) {
	//压缩
	// err := doGzip([]byte("hello，golang"), "/data/aaa.gz", "test.txt")
	// if err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }

	//解压
	content, err := unGzip("/home/lei/Downloads/1011log/JB05.txt.gz")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(string(content))
}

func main() {
	var name = "/home/lei/Downloads/1011log/JB06.txt.gz"
	content, err := unGzip("/home/lei/Downloads/1011log/JB06.txt.gz")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println(string(content))
	fileName := strings.Split(name, ".")[0]
	fmt.Println(fileName)
	//写入文件
	if err = ioutil.WriteFile(fileName, content, 0666); err != nil {
		fmt.Println("写入错误：", err)
	}

	//读取文件
	_, err = ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("读取错误：", err)
		return
	}
	// fmt.Println("读取成功，文件内容：", string(fileContent))

}

/*
*
压缩bytes内容
1.根据指定目录创建文件
2.根据文件资源对象生成gzip Writer对象
3.往gzip Writer对象写入内容
*/
func doGzip(content []byte, path string, fileName string) error {
	gzFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer gzFile.Close()
	gzipWriter := gzip.NewWriter(gzFile)

	defer gzipWriter.Close()
	gzipWriter.Name = fileName
	_, err = gzipWriter.Write(content)
	if err != nil {
		return err
	}
	return nil
}

/*
*
解压
1.根据指定路径打开gzip文件资源
2.由文件资源创建gzip Reader对象
3.将gizp 内存流copy到指定buf并返回bytes
*/
func unGzip(path string) (content []byte, err error) {
	gzipFile, err := os.Open(path)
	fmt.Println(gzipFile.Name())
	if err != nil {
		return
	}
	defer gzipFile.Close()
	gzipReader, err := gzip.NewReader(gzipFile)
	// fmt.Println(gzipReader)

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
