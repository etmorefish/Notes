package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var cnt int

func main() {
	// cur_dir, _ := os.Getwd()
	// Walk(cur_dir)
	pwd := "/home/lei/Downloads/1011log/"

	WalkDir()
	WalkFiles(pwd)
	fmt.Printf("Total files count: %d.", cnt)
}

// 1.遍历目录下(包含子目录)所有文件及目录
func Walk(path string) {
	dir, _ := os.ReadDir(path)
	for _, file := range dir {
		file_path := filepath.Join(path, file.Name())
		cnt++
		fmt.Printf("%-4d: %s\n", cnt, file_path)
		if file.IsDir() {
			Walk(file_path)
		}
	}
}

// 2.遍历目录下的所有文件及目录信息
// 注意：此方法会将当前工作目录也一并输出，所以数量会多一个
func WalkDir() {
	pwd, _ := os.Getwd()
	pwd = "/home/lei/Downloads/1011log/"

	// 获取当前目录下的所有文件或目录信息
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		cnt++
		fmt.Printf("%-4d. %s\n", cnt, path)
		fmt.Println(path)        // 打印path信息   (带绝对路径)
		fmt.Println(info.Name()) // 打印文件或目录名(无绝对路径)
		return nil
	})
}

// 3.遍历当前目录下所有文件(包含子目录下)
func WalkFiles(path string) {
	dir, _ := os.ReadDir(path)
	for _, file := range dir {
		file_path := filepath.Join(path, file.Name())
		if file.IsDir() {
			Walk(file_path)
		} else {
			cnt++
			fmt.Printf("%-4d: %s\n", cnt, file_path)
		}
	}
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
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

/*
// 创建文件
os.Create(name string)

// 删除文件
os.Remove(name string)
// 使用strings.HasSuffix()判断是否包含此后缀名或者使用数组判断

// 重命名文件
os.Rename(oldpath, newpath string)

// 读取文件
os.ReadFile(name string)

// 重写文件
os.WriteFile(name string, data []byte, perm FileMode)	// FileMode: os.ModePerm

// 目录操作
// 创建目录
os.Mkdir(name string, perm FileMode)
// 创建多级目录
os.MkdirAll(path string, perm FileMode)

// 删除目录
os.RemoveAll(path string)

// 获取当前工作目录
os.Getwd()

// 切换当前工作目录
os.Chdir(path string)

// 获取临时目录？
os.TempDir()

// 读取目录
os.ReadDir(name string)
*/
// https://blog.csdn.net/weixin_56461542/article/details/125240529
