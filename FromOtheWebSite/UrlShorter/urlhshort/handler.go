package urlhshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

func MapHandler(pathUrl map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathUrl[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YamlHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}

	return MapHandler(buildMap(pathUrls), fallback), nil

}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathMap := make(map[string]string)

	for _, pu := range pathUrls {
		pathMap[pu.Path] = pu.URL
	}
	return pathMap
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
