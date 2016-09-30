package main

import (        
    "log"
    "os"
    "io"    
    "net/http"    
)

func RunIPFSCommand(url string) {    
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    } else {
        defer response.Body.Close()
        _, err := io.Copy(os.Stdout, response.Body)
        if err != nil {
            log.Fatal(err)
        }
    }    
}

func IPFSAdd() {            
    client := &http.Client{}
    req, _ := http.NewRequest("GET", "http://10.1.1.54:5001/api/v0/add", nil)
    req.Header.Set("Content-Type", "multipart/form-data")
    res, _ := client.Do(req)
    io.Copy(os.Stdout, res.Body)    
    
    /*
    response, err := http.Get("http://10.1.1.54:5001/api/v0/add")
    if err != nil {
        log.Fatal(err)
    } else {
        defer response.Body.Close()
        _, err := io.Copy(os.Stdout, response.Body)
        if err != nil {
            log.Fatal(err)
        }
    }
    */
}