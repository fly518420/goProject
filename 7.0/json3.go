package main

import (
    "encoding/json"
    "fmt"
)

type Server struct {
    ServerName string `json:"serverName"`
    ServerIP string `json:"serverIP"`
}

type ServerSlice struct {
    ServerID string `json:"-"`
    Servers []Server `json:"servers"`
}

func main() {
    var s ServerSlice
    s.ServerID = "20160927"
    s.Servers = append(s.Servers, Server{"beijing", "127.0.0.1"})
    s.Servers = append(s.Servers, Server{"shanghai", "127.0.0.2"})
    data, err := json.Marshal(&s)
    if err != nil {
        fmt.Printf("error : %v\n", err)
        return
    }
    fmt.Println(string(data))
}

