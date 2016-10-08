package main

import (
	"fmt"
	"io"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("error : %v\n", err)
		panic(err)
	}
}

func main() {
	// err := os.Mkdir("newdir", 0777)
	// checkErr(err)

	// err := os.Remove("newdir") // 注意当被删除的对象是目录的时候，目录下面有其他文件或者目录会出错 remove dir: directory not empty
	// checkErr(err)

	// err := os.MkdirAll("test/file", 0777)
	// checkErr(err)

	// err := os.Rename("test", "dir")
	// checkErr(err)

	// file, err := os.Create("dir/file/test.txt")
	// checkErr(err)
	// fmt.Println(file)

	// defer file.Close()

	// len, err := file.WriteString("hello,world")
	// checkErr(err)
	// fmt.Println(len)

	// err = os.Remove("dir/file/test.txt") // 注意删除的时候传递的参数
	// checkErr(err)

	// err := os.RemoveAll("dir") // 递归删除
	// checkErr(err)

	// file, err := os.Create("test.txt")
	// checkErr(err)

	// defer file.Close()

	// len, err := file.WriteString("hello,world")
	// checkErr(err)
	// fmt.Println(len)

	buf := make([]byte, 1024)
	file, err := os.Open("test.txt")
	checkErr(err)
	for {
		n, err := file.Read(buf)
		fmt.Println(n)
		if 0 == n && err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		checkErr(err)
		os.Stdout.Write(buf[:n])
		fmt.Println()
	}

    var a string
    var b int
    fmt.Scan(&a, &b)
    fmt.Println(a, b)
    fmt.Scanln(&a, &b)
    fmt.Println(a, b)

}
