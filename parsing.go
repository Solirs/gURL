package main

import (
	"strings"
	"regexp"
)

func StatusCode(response string) string{
    scode, err := regexp.Compile(`HTTP\/1.0 \s*([^\n\r ]*)*`)

    checkError(err)


    return scode.FindString(response)[len(scode.FindString(response))-3:]
}

func Redirectlocation(response string) string{
    scode, err := regexp.Compile(`Location: \s*([^\n\r ]*)`)

    checkError(err)

    ret := strings.Trim(scode.FindString(response), "Location: ")


    return ret
}