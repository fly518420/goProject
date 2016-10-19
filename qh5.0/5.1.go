package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    var c byte
    var str string

    // 初始化标准的输入设备
    r := bufio.NewReader(os.Stdin)
    // 初始化标注的输出设备
    w := bufio.NewWriter(os.Stdout)

    var tip string = "来自标准输出设备:"
    for i := 0; i < len(tip); i++ {
        w.WriteByte(byte(tip[i]))
    }

    for i := 0; ; i++ {
        c, _ = r.ReadByte()
        if c == 10 {
            w.Flush()
            break // 如果是回车停止接收退出
        } else {
            w.WriteByte(c)
            str += string(c)
        }
    
    } 
    fmt.Printf("\n")
    fmt.Println(str)
}

