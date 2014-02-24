package main

import (
    "encoding/json"
	"fmt"
	"io/ioutil"
//    "log"
	"net/http"
    "os"
//	"path/filepath"
	"strconv"
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
		r := goget(host + "/manifest.json")
		dec := json.NewDecoder(r.Body)
		dec.Decode(&manifest)
		
		for _, w := range manifest.Workloads {
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
		goserve(8080, "/tmp")
	case "version":
		host := os.Args[1:][1]
		r := goget(host + "/version")
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Printf(string(body))
		r.Body.Close()
	}
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

func scanroot(root string) string {	
/*	func visit(path string, f os.FileInfo, err error) error {
		// look for VMs
	}
	err := filepath.Walk(root, visit)
*/
	return(`{"version": 1, "workloads": [ {"name": "tinycore", "title": "Tiny Core Linux", "components": [ { "label": "root", "file": "tinycore.vmdk" } ] } ]}`)
}

func goserve(port int, root string) {
	manifestjson := scanroot(root)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hello, World!")})
	http.HandleFunc("/manifest.json", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, manifestjson)})
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "swap-0.1")})
	
	if err:= http.ListenAndServe(":" + strconv.Itoa(port), nil); err != nil {
		fmt.Printf("ListenAndServe error: %v\n", err)
	}
}