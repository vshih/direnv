package main

import (
	"testing"
)

func TestSomething(t *testing.T) {
	paths := eachDir("C:/foo/b//bar/")
	if len(paths) != 4 {
		t.Error(paths)
	}
	paths = eachDir("C:/")
	if len(paths) < 1 || paths[0] != "C:/" || len(paths) > 1 {
		t.Error(paths)
	}
}

func TestRootDir(t *testing.T) {
	var r string
	r = rootDir("C:/foo")
	if r != "C:/foo" {
		t.Error(r)
	}

	r = rootDir("C:/foo/bar")
	if r != "C:/foo" {
		t.Error(r)
	}
}
