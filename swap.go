package main

import (
//    "encoding/json"
	"fmt"
	"io/ioutil"
//    "log"
	"net/http"
    "os"
)

func main() {
	// Define flags
	action := os.Args[1:][0]
	
	switch action {
	case "list":
		fmt.Printf("tinycore\n")
	case "get":
	case "put":
	case "discover":
		host := os.Args[1:][1]
		fmt.Printf(httpget(host + "/.well-known/swap/root"))
	case "version":
		host := os.Args[1:][1]
		fmt.Printf(httpget(host + "/version"))
	}
/*	
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        for k := range v {
            if k != "workloads" {
                delete(v, k)
            }
        }
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }
*/
}

func httpget(url string) string {
	resp, err := http.Get(url)
	if resp.StatusCode != 200 {
		fmt.Printf("Server returned HTTP Status %v, expected 200\n", resp.StatusCode)
        os.Exit(1)
	}
	if err != nil {
		// handle error
        fmt.Printf("%s\n", err)
        os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
    if err != nil {
        fmt.Printf("%s\n", err)
        os.Exit(1)
    }
    return string(body)
}
