package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
)

var httptype string
var u string
var uagent string
var content string

func initflags(){

    flag.StringVar(&httptype, "X", "invalid", "GET POST or HEAD")
    flag.StringVar(&u, "u", "invalid", "URL")
    flag.StringVar(&uagent, "ua", "Iorem Ipsum", "User Agent") 
    flag.StringVar(&content, "c", "", "content") 


}

func main(){

    initflags()

    flag.Parse()

    url, err := url.Parse(u)







    checkError(err)

    if url.Scheme == "https"{

        switch httptype {
        case "GET":
            GETRequestHTTPS(url)
    
        case "HEAD":
            HEADRequest(url)
    
        case "POST":
            POSTRequestHTTPS(url)
    
        default:
            log.Fatal("Invalid or unsupported type specified, QUITTING!!!")
        }

    }else if url.Scheme == "http"{

        switch httptype {
        case "GET":
            GETRequestHTTP(url)
    
        case "HEAD":
            HEADRequest(url)
    
        case "POST":
            POSTRequestHTTP(url)
    
        default:
            log.Fatal("Invalid or unsupported type specified, QUITTING!!!")
        }
    }



}





func GETRequestHTTP(u *url.URL) {

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

    

    req += fmt.Sprintf("\r\n")

    fmt.Print(req)

    _, err = con.Write([]byte(req))
    checkError(err)


    res, err := io.ReadAll(con)
    checkError(err)

    fmt.Print(string(res))

    fmt.Print(Redirectlocation(string(res)), "\n")
}

func GETRequestHTTPS(u *url.URL) {

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

func checkError(err error) {

    if err != nil {

        log.Fatal(err)
    }
}

func HEADRequest(u *url.URL){

    con, err := net.Dial("tcp", u.Host)

    if err != nil{
        con, err = net.Dial("tcp", u.Host + ":80")
        checkError(err)
    }

    
    req := "HEAD / HTTP/1.1\r\n"
    req += "Connection: close\r\n"
    req += "\r\n"

    _, err = con.Write([]byte(req))
    checkError(err)

    res, err := io.ReadAll(con)
    checkError(err)

    fmt.Println(string(res))

}

func POSTRequestHTTP(u *url.URL){


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
    req += fmt.Sprintf("\r\n")
    req += content

    fmt.Print(req)

    _, err = con.Write([]byte(req))
    checkError(err)

    res, err := io.ReadAll(con)
    checkError(err)
    fmt.Println(string(res))


}


func POSTRequestHTTPS(u *url.URL){


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




