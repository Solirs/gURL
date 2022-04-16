package main

import (
	"flag"
	"log"
	"net/url"
)


type GURLHTTP struct{

}

type GURLHTTPS struct{

}


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

    ghttp := GURLHTTP{}
    ghttps := GURLHTTPS{}







    checkError(err)

    if url.Scheme == "https"{

        switch httptype {
        case "GET":
            ghttps.GETRequestHTTPS(url)
    
        case "HEAD":
            ghttp.HEADRequest(url)
    
        case "POST":
            ghttps.POSTRequestHTTPS(url)
    
        default:
            log.Fatal("Invalid or unsupported type specified, QUITTING!!!")
        }

    }else if url.Scheme == "http"{

        switch httptype {
        case "GET":
            ghttp.GETRequestHTTP(url)
    
        case "HEAD":
            ghttp.HEADRequest(url)
    
        case "POST":
            ghttp.POSTRequestHTTP(url)
    
        default:
            log.Fatal("Invalid or unsupported type specified, QUITTING!!!")
        }
    }



}








func checkError(err error) {

    if err != nil {

        log.Fatal(err)
    }
}



