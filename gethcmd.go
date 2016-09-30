package main

import (
    "fmt"                    
    "bytes"
    "io/ioutil"    
    "net/http"    
)

func RunGethCommand(url string, jsonStr []byte) {            
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))    
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)    
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}