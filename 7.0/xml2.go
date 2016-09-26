package main

import (
    "encoding/xml"
    "fmt"
    "os"
)

type Server struct {
    ServerName string `xml:"serverName"`
    ServerIP string `xml:"serverIP"`
}

type Servers struct {
    XMLName xml.Name `xml:"servers"`
    Version string `xml:"version,attr"`
    Svs []Server `xml:"server"`
}

func main() {
    v := &Servers{Version:"1"}
    v.Svs = append(v.Svs, Server{"Shanghai_VPN", "127.0.0.1"})
    v.Svs = append(v.Svs, Server{"Beijing_VPN", "127.0.0.2"})

    data1, err := xml.Marshal(v)
    if err != nil {
        fmt.Printf("error:%v\n", err)
        return
    }
    data2, err := xml.MarshalIndent(v, "  ", "    ")
    if err != nil {
        fmt.Printf("error:%v\n", err)
        return
    }

    os.Stdout.Write([]byte(xml.Header))
    os.Stdout.Write(data1)

    fmt.Println()
    os.Stdout.Write([]byte(xml.Header))
    os.Stdout.Write(data2)
}