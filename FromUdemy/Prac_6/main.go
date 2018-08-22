package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func helloWorld() string {
	return "Hello World from weizhe!"
}
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler http request")
	fmt.Fprintf(w, "%s, I love %s", helloWorld(), r.URL.Path[1:])
}

func pinger(port string) error {
	resp, err := http.Get("http://localhost:" + port)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("server return non-200 status code")
	}

	return nil
}

func main() {

	var port string
	var ping bool
	flag.StringVar(&port, "p", "8080", "Server Port")
	flag.BoolVar(&ping, "ping", false, "Server health")
	flag.Parse()

	if ping {
		if err := pinger(port); err != nil {
			log.Printf("ping server error: %s", err)
		}

		return
	}
	http.HandleFunc("/", handler)
	log.Printf("Server run on port: %s", port)
	http.ListenAndServe(":"+port, nil)
}
