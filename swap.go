package main

import (
    "encoding/json"
	"fmt"
	"io/ioutil"
//    "log"
	"net/http"
    "os"
)

func main() {
	// Define flags
	action := os.Args[1:][0]
	
    var manifest struct {
        Workloads []struct {
            Name              string
            Title             string
            Components        []struct {
            	Label         string
				File          string
				Checksum      string
            }
        }
    }
	
	switch action {
	case "list":
		fmt.Printf("tinycore\n")
		
	case "get":
		
	case "put":
		
	case "parse":
		host := os.Args[1:][1]
		r := goget(host + "/manifest")
		dec := json.NewDecoder(r.Body)
		dec.Decode(&manifest)
		
		for _, w := range manifest.Workloads {
//			fmt.Printf("%s (%s): [%s]%s\n", w.Title, w.Name, w.Components[0].Label, w.Components[0]].File)
			fmt.Printf("%s (%s)", w.Title, w.Name)
			for _, c := range w.Components {
				fmt.Printf(" [%s]%s", c.Label, c.File)
			}
			fmt.Printf("\n")
		}
		
	case "discover":
		host := os.Args[1:][1]
		r := goget(host + "/.well-known/swap/root")
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Printf(string(body))
		r.Body.Close()
		
	case "server":
		
	case "version":
		host := os.Args[1:][1]
		r := goget(host + "/version")
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Printf(string(body))
		r.Body.Close()
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

func goget(url string) *http.Response {
	r, err := http.Get(url)
	if r.StatusCode != 200 {
		fmt.Printf("Server returned HTTP Status %v, expected 200\n", r.StatusCode)
        os.Exit(1)
	}
	if err != nil {
		// handle error
        fmt.Printf("%s\n", err)
        os.Exit(1)
	}
//	defer r.Body.Close()
	
	return r
//	body, err := ioutil.ReadAll(resp.Body)
	
//    if err != nil {
//        fmt.Printf("%s\n", err)
//        os.Exit(1)
//    }
//    return string(body)
}
