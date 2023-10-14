package leetcode

import (
	"strings"
)

func simplifyPath(path string) string {
	items := strings.Split(path, "/")
	items = items[1:]
	layers := []string{}
	for _, item := range items {
		if item == "." || item == "" {
			continue
		}
		if item == ".." {
			if len(layers) > 0 {
				layers = layers[0 : len(layers)-1]
			}
			continue
		}
		layers = append(layers, item)
	}
	return "/" + strings.Join(layers, "/")
}
