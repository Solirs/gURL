package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
)

func (ghttps GURLHTTPS)GETRequestHTTPS(u *url.URL) {

    con, err := net.Dial("tcp", u.Host)

    if err != nil{
        con, err = net.Dial("tcp", u.Host + ":443")
        checkError(err)
    }

    client := tls.Client(con, &tls.Config{
        ServerName: u.Host,
    })

    defer client.Close()

    if err := client.Handshake(); err != nil {
        log.Fatal(err)
    }

    
    req := fmt.Sprintf("GET %s HTTP/1.1\r\n", u.Path)
    req += fmt.Sprintf("Host: %v\r\n", u.Host)
    req += fmt.Sprintf("Connection: close\r\n")
    req += fmt.Sprintf("User-Agent: %v\r\n", uagent)
    //req += fmt.Sprintf("Filename: %s\r\n", u.Path)

    req += fmt.Sprintf("\r\n")

    fmt.Print(req)

    _, err = client.Write([]byte(req))
    checkError(err)


    res, err := io.ReadAll(client)
    checkError(err)

    fmt.Print(string(res))

    fmt.Print(Redirectlocation(string(res)), "\n")
}




func (ghttps GURLHTTPS)POSTRequestHTTPS(u *url.URL){


    con, err := net.Dial("tcp", u.Host)

    if err != nil{
        con, err = net.Dial("tcp", u.Host + ":443")
        checkError(err)
    }


    client := tls.Client(con, &tls.Config{
        ServerName: u.Host,
    })

    defer client.Close()

    if err := client.Handshake(); err != nil {
        log.Fatal(err)
    }

    
    req := fmt.Sprintf("POST / HTTP/1.1\r\n")
    req += fmt.Sprintf("Host: %v\r\n", u.Host)
    req += fmt.Sprintf("Connection: close\r\n")
    req += fmt.Sprintf("Content-type: text/plain\r\n")
    req += fmt.Sprintf("Content-length: %d\r\n", len(content))
    req += fmt.Sprintf("\r\n")
    req += content

    fmt.Print(req)

    _, err = client.Write([]byte(req))
    checkError(err)

    res, err := io.ReadAll(client)
    checkError(err)
    fmt.Println(string(res))


}
