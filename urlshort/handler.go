package urlshort

import (
	"fmt"
	"net/http"
	"gopkg.in/yaml.v2"
    "encoding/json"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		url, ok := pathToUrls[req.URL.Path]
		if !ok {
			fallback.ServeHTTP(res, req)
			return
		}

		res.Header().Add("Location", url)
		res.WriteHeader(301)
		fmt.Printf("%s to %s\n", pathToUrls[req.URL.Path], url)
	}
}

type pair struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func parseYAML(yml []byte) ([]pair,error) {
	var pairs []pair
	err := yaml.Unmarshal(yml, &pairs)
	return pairs, err
}

func buildMap(pairs []pair) map[string]string{
    pathToUrls := make(map[string]string, len(pairs))
    for _, item := range pairs {
        pathToUrls[item.Path] = item.Url
    }
    return pathToUrls
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yml)
	if err != nil {
        fmt.Println("There is error in the parsing yaml")
		return nil, err
	}
	pathMap := buildMap(parsedYAML)
	return MapHandler(pathMap, fallback), nil
}

func JSONHandler(jsonData[]byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pairs []pair
    pathToUrls := make(map[string]string, len(pairs))
    err := json.Unmarshal(jsonData, &pairs)
    for _, item := range pairs {
        pathToUrls[item.Path] = item.Url
    }
    if err != nil {
        return nil, err
    }
    return MapHandler(pathToUrls, fallback), nil
}
