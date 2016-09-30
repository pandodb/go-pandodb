package main

import (
    "fmt"    
    "os"
    "errors"    
    "encoding/json"        
    "github.com/spf13/cobra" 
)

type Configuration struct {        
    IPFSNodeAddress string
    GethNodeAddress string
}

func main() {        		    
    configuration, err := loadConfiguration();    
    if(err != nil) {
        fmt.Println(err)
        return
    }    
    addCommands(configuration);
}

func AddOne(x int) int {
    return x + 1
}

func addCommands(configuration Configuration) {
    var cmdVersion = &cobra.Command{
        Use:   "version",
        Short: "Show version of app",        
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("pandodb version: 0.0.001")
        },
    }        
    
    var cmdIPFS = &cobra.Command{
        Use:   "ipfs",
        Short: "ipfs [subcommand]",
    }        
    
    var cmdGeth = &cobra.Command{
        Use:   "geth",
        Short: "geth [subcommand]",
    }        
    
    var cmdIPFSVersion = &cobra.Command{
        Use:   "version",
        Short: "verion",        
        Run: func(cmd *cobra.Command, args []string) {
            //RunIPFSCommand(configuration.IPFSNodeAddress + "/api/v0/version")
            IPFSAdd();
        },
    }        
    
    var cmdGethVersion = &cobra.Command{
        Use:   "version",
        Short: "verion",        
        Run: func(cmd *cobra.Command, args []string) {
            var jsonStr = []byte(`{"jsonrpc":"2.0","method":"net_version","params":[],"id":67}`)
            RunGethCommand(configuration.GethNodeAddress, jsonStr);
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
    cmdIPFS.AddCommand(cmdIPFSVersion)
    cmdGeth.AddCommand(cmdGethVersion)
    
    cmdDB.AddCommand(cmdDBNew)
    
    rootCmd.AddCommand(cmdVersion, cmdGeth, cmdIPFS, cmdDB)    
    rootCmd.Execute()
}

func loadConfiguration() (Configuration, error) {
    configuration := Configuration{}

    if _, err := os.Stat("/path/to/whatever"); err == nil {        
        return configuration, errors.New("Configuration file not found.")
    }

    file, _ := os.Open("config.json")
    decoder := json.NewDecoder(file)    
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
    }    
    return configuration, nil
}