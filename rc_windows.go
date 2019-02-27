package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/// Utils

func rootDir(path string) string {
	path, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	path = strings.Replace(path, string(os.PathSeparator), "/", -1)
	// Skip "C:/"
	start := 3
	i := strings.Index(path[start:], "/")
	if i < 0 {
		return path
	}
	return path[:i+start]
}

func eachDir(path string) (paths []string) {
	path, err := filepath.Abs(path)
	if err != nil {
		return
	}
	path = strings.Replace(path, string(os.PathSeparator), "/", -1)

	paths = []string{path}

	is_root_regex, _ := regexp.Compile("[A-Z]:/$")

	if is_root_regex.MatchString(path) {
		return
	}

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			if !is_root_regex.MatchString(path) {
				path = path[:i]
			}
			paths = append(paths, path)
		}
	}

	return
}

func ToUnixDir(path string) string {
	if os.ExpandEnv("$MSYSTEM") == "MSYS" {
		// Strip any colons.
		re := regexp.MustCompile("^([A-Z]):[/\\\\]")
		path = re.ReplaceAllString(path, "/$1/")
	}

	return strings.Replace(path, "\\", "/", -1)
}

func ToUnixPathList(pathList string) string {
	split := strings.Split(pathList, ";")

	for i, p := range split {
		split[i] = ToUnixDir(p)
	}

	return strings.Join(split, ":")
}
