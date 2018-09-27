package main

import (
	"reflect"
	"testing"
)

func TestPathsToKeys(t *testing.T) {
	paths := []string{
		"target/lambda/lambda.go",
		"target/build.json",
		"target/riff-raff.yaml",
	}

	actual := PathsToKeys(paths, "dotcom:foo/", "target/")
	expected := map[string]string{
		"target/lambda/lambda.go": "dotcom:foo/lambda/lambda.go",
		"target/build.json":       "dotcom:foo/build.json",
		"target/riff-raff.yaml":   "dotcom:foo/riff-raff.yaml",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("PathsToKeys: expected %s, actual %s", expected, actual)
	}
}
