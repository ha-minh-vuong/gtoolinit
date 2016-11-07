package main

import (
	"fmt"
	"os"
	"os/exec"
)

var tools = []string{
	"github.com/nsf/gocode",
	// lint
	"github.com/golang/lint/golint",
	"github.com/kisielk/errcheck",
	"github.com/alecthomas/gometalinter",
	// format
	"golang.org/x/tools/cmd/goimports",
	"sourcegraph.com/sqs/goreturns",
	// source analysis
	"golang.org/x/tools/cmd/oracle",
	"github.com/rogpeppe/godef",
	"golang.org/x/tools/cmd/gorename",
	// doc
	"github.com/zmb3/gogetdoc",
	// dependency graph
	"github.com/kisielk/godepgraph",
	// benchmark
	"golang.org/x/tools/cmd/benchcmp",
	"github.com/cespare/prettybench",
	"github.com/ajstarks/svgo/benchviz",
}

func main() {

	if len(os.Getenv("GOPATH")) == 0 {
		fmt.Fprintln(os.Stderr, "GOPATH has not been set.")
		os.Exit(1)
	}

	fmt.Println("Installing go tools ...")

	cmd := "go"
	args := []string{"get", "-u"}

	for _, t := range tools {
		targs := append(args, t)
		if err := exec.Command(cmd, targs...).Run(); err != nil {
			fmt.Print(t, "  x ")
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Printf("%v  √\n", t)
	}

}
