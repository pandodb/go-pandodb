package main

import (
    "fmt"    
    "os"
    "log"
    "io"
    "bytes"
    "io/ioutil"
    "encoding/json"
    "net/http"    
    "github.com/spf13/cobra" 
)

type Configuration struct {
    IPFSNodeAddress string
    EthereumNodeAddress string
}

func main() {        		    
    configuration := loadConfiguration();    
    addCommands(configuration);
}

func addCommands(configuration Configuration) {
    var cmdVersion = &cobra.Command{
        Use:   "version",
        Short: "Show version of app",        
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("pandodb version: 0.0.1")
        },
    }
    
    var cmdIPFSNodeAddress = &cobra.Command{
        Use:   "ipfsnodeaddress",
        Short: "Show address of IPFS node",        
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("IPFS node address: " + configuration.IPFSNodeAddress)
        },
    }        
    
    var cmdPing = &cobra.Command{
        Use:   "ping",
        Short: "ping [ipfs/ethereum]",
    }        
    
    var cmdPingIPFS = &cobra.Command{
        Use:   "ipfs",
        Short: "Pings the IPFS node",        
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Pinging: " + configuration.IPFSNodeAddress)
            pingIPFS();
        },
    }        
    
    var cmdPingEthereum = &cobra.Command{
        Use:   "ethereum",
        Short: "Pings the Ethereum node",        
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Pinging: " + configuration.EthereumNodeAddress)
            pingEthereum();
        },
    }        
    
    var cmdDB = &cobra.Command{
        Use:   "db",
        Short: "Database",
    }        
    
    var cmdDBNew = &cobra.Command{
        Use:   "new",
        Short: "Create a new Database [name]",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Created New Database: " + args[0])            
        },
    }    
	
	var rootCmd = &cobra.Command{Use: "app"}
    cmdPing.AddCommand(cmdPingIPFS)
    cmdPing.AddCommand(cmdPingEthereum)
    
    cmdDB.AddCommand(cmdDBNew)
    
    rootCmd.AddCommand(cmdVersion, cmdPing, cmdDB, cmdIPFSNodeAddress)    
    rootCmd.Execute()
}

func pingIPFS() {
    response, err := http.Get("http://127.0.0.1:5001/api/v0/version")
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

func pingEthereum() {
    url := "http://localhost:8545"
    var jsonStr = []byte(`{"jsonrpc":"2.0","method":"net_version","params":[],"id":67}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()    
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response:", string(body))
}


func loadConfiguration() Configuration {
    file, _ := os.Open("config.json")
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
    }    
    return configuration    
}