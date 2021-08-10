package handler

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(rw, r, url, http.StatusFound)
			return
		}

		fallback.ServeHTTP(rw, r)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL.
func YAMLHandler(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlData, &pathUrls)

	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)

	for _, url := range pathUrls {
		pathsToUrls[url.Path] = url.URL
	}

	return MapHandler(pathsToUrls, fallback), nil
}
