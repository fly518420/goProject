package main

import (
    "fmt"
    "strconv"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    flag, err := strconv.ParseBool("true")
    checkErr(err)

    fmt.Println(flag)

    n1, err := strconv.ParseFloat("23.33", 64)
    checkErr(err)

    fmt.Println(n1)

    n2, err := strconv.ParseInt("17", 10, 64)
    checkErr(err)

    fmt.Println(n2)

    n3, err := strconv.ParseInt("17", 8, 64)
    checkErr(err)

    fmt.Println(n3)

    n4, err := strconv.ParseInt("17", 16, 64)
    checkErr(err)

    fmt.Println(n4)

    s := strconv.FormatBool(false)
    fmt.Println(s)

    s = strconv.FormatFloat(12.11, 'g', 3, 64)
    fmt.Println(s)

    s = strconv.FormatInt(17, 10)
    fmt.Println(s)

    s = strconv.FormatInt(17, 8)
    fmt.Println(s)

    s = strconv.FormatInt(17, 16)
    fmt.Println(s)
}

