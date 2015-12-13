package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGen(t *testing.T) {
	g := new(generator)

	// TODO: figure out a better way of current importpath
	pkgPath := "github.com/golang/mock/mockgen/ignore/orig"
	interfaces := "Adder,HasVariadic"
	pkg, err := Reflect(pkgPath, strings.Split(interfaces, ","))
	if err != nil {
		t.Fatalf("unagle to reflect to find %s %s: %s", pkgPath, interfaces, err)
	}

	g.srcPackage = pkgPath
	g.srcInterfaces = interfaces

	if err := g.Generate(pkg, "orig_test"); err != nil {
		t.Fatalf("unable to generate orig mocks: %s")
	}

	dir, err := ioutil.TempDir("", "mock_orig_test_")
	if err != nil {
		t.Fatalf("unable to make tempdir: %s", err)
	}

	mockDir := filepath.Join(dir, "orig_test")
	err := os.Mkdir(mockDir)
	if err != nil {
		t.Fatalf("unable to make mockDir %#v: %s", mockDir, err)
	}
	err := ioutil.WriteFile(filepath.Join(mockDir, "mock_orig.go"), g.Output(), 0644)
	if err != nil {
		t.Fatalf("unable to write mocks: %s", err)
	}

}
