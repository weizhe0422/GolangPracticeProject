package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/weizhe0422/GolangPractice/FromOtheWebSite/UrlShorter/urlhshort"
)

func main() {
	mode := flag.String("mode", "map", "mode to select(map/yaml)")
	var handler http.Handler

	switch *mode {
	case "map":
		pathToUrls := map[string]string{
			"/yahoo":          "https://tw.yahoo.com/?p=us",
			"/weizhe0422Blog": "https://weizhe0422.github.io/",
		}
		handler := urlhshort.MapHandler(pathToUrls, defaultMux())
		http.ListenAndServe(":8080", handler)
	case "yaml":
		yamlUrls := `
	- path: /urlshort
	  url: https://github.com/gophercises/urlshort
  	- path: /urlshort-final
	  url: https://github.com/gophercises/urlshort/tree/solution	
	`
		handler, err := urlhshort.YamlHandler([]byte(yamlUrls), defaultMux())
		if err != nil {
			panic(err)
		}
		http.ListenAndServe(":8080", handler)
	default:
		http.ListenAndServe(":8080", handler)
	}

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is default response message!")
}
