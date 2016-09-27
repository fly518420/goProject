package main

import (
    "encoding/json"
    "fmt"
)

type Server struct {
    ServerName string
    ServerIP string
}

type ServerSlice struct {
    Servers []Server
}

func main()  {
    var s ServerSlice
    s.Servers = append(s.Servers, Server{"beijing", "127.0.0.1"})
    s.Servers = append(s.Servers, Server{"shanghai", "127.0.0.2"})
    data, err := json.Marshal(s)
    if err != nil {
        fmt.Printf("error : %v\n", err)
        return
    }
    fmt.Println(string(data))
}