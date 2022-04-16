package main

import (
	"fmt"
	"io"
	"net"
	"net/url"
)

func (ghttp GURLHTTP)GETRequestHTTP(u *url.URL) {

    fmt.Println(u.Host)

    con, err := net.Dial("tcp", u.Host)

    if err != nil{
        con, err = net.Dial("tcp", u.Host + ":80")
        checkError(err)
    }

    defer con.Close()
    
    req := fmt.Sprintf("GET %s HTTP/1.1\r\n", u.Path)
    req += fmt.Sprintf("Host: %v\r\n", u.Host)
    req += fmt.Sprintf("Connection: close\r\n")
    req += fmt.Sprintf("User-Agent: %v\r\n", uagent)

    for i := 0; i < len(headers); i++{
        req += fmt.Sprintf("%s\r\n", headers[i])
    }

    

    req += fmt.Sprintf("\r\n")

    fmt.Print(req)

    _, err = con.Write([]byte(req))
    checkError(err)


    res, err := io.ReadAll(con)
    checkError(err)

    fmt.Print(string(res))

    fmt.Print(Redirectlocation(string(res)), "\n")
}


func (ghttp GURLHTTP)HEADRequest(u *url.URL){

    con, err := net.Dial("tcp", u.Host)

    if err != nil{
        con, err = net.Dial("tcp", u.Host + ":80")
        checkError(err)
    }

    
    req := "HEAD / HTTP/1.1\r\n"
    req += "Connection: close\r\n"
    for i := 0; i < len(headers); i++{
        req += fmt.Sprintf("%s\r\n", headers[i])
    }
    req += "\r\n"

    _, err = con.Write([]byte(req))
    checkError(err)

    res, err := io.ReadAll(con)
    checkError(err)

    fmt.Println(string(res))

}



func (ghttp GURLHTTP)POSTRequestHTTP(u *url.URL){


    con, err := net.Dial("tcp", u.Host)

    if err != nil{
        con, err = net.Dial("tcp", u.Host + ":80")
        checkError(err)
    }

    
    req := fmt.Sprintf("POST / HTTP/1.1\r\n")
    req += fmt.Sprintf("Host: %v\r\n", u.Host)
    req += fmt.Sprintf("Connection: close\r\n")
    req += fmt.Sprintf("Content-type: text/plain\r\n")
    req += fmt.Sprintf("Content-length: %d\r\n", len(content))
    for i := 0; i < len(headers); i++{
        req += fmt.Sprintf("%s\r\n", headers[i])
    }
    req += fmt.Sprintf("\r\n")
    req += content

    fmt.Print(req)

    _, err = con.Write([]byte(req))
    checkError(err)

    res, err := io.ReadAll(con)
    checkError(err)
    fmt.Println(string(res))


}
