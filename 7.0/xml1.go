package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

type Server struct {
    XMLName xml.Name `xml:"server"`
    ServerName string `xml:"serverName"`
    ServerIP string `xml:"serverIP"`
}

type Recurlyservers struct {
    XMLName xml.Name `xml:"servers"`
    Version string `xml:"version,attr"`
    Servers []Server `xml:"server"`
    Description string `xml:",innerxml"`
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main()  {
    file, err := os.Open("servers.xml")
    checkErr(err)

    defer file.Close()
    
    data, err := ioutil.ReadAll(file)
    checkErr(err)

    v := Recurlyservers{}
    err = xml.Unmarshal(data, &v)
    checkErr(err)

    fmt.Println(v)
    fmt.Println(v.XMLName)
    fmt.Println(v.Version)
    fmt.Println(v.Servers)
    fmt.Println(v.Description)

}