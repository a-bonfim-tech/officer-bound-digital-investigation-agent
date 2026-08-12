package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] != "all" {
		fmt.Fprintln(os.Stderr, "usage: verify all")
		os.Exit(2)
	}
	goBin, gofmt := os.Getenv("OBDIA_GO_BIN"), os.Getenv("OBDIA_GOFMT_BIN")
	if goBin == "" || gofmt == "" || !filepath.IsAbs(goBin) || !filepath.IsAbs(gofmt) {
		fmt.Fprintln(os.Stderr, "absolute OBDIA_GO_BIN and OBDIA_GOFMT_BIN required")
		os.Exit(2)
	}
	goRoot := filepath.Dir(filepath.Dir(goBin))
	if filepath.Clean(gofmt) != filepath.Join(goRoot, "bin", "gofmt") || filepath.Clean(goBin) != filepath.Join(goRoot, "bin", "go") {
		fmt.Fprintln(os.Stderr, "Go tools must be the go and gofmt pair from one absolute GOROOT")
		os.Exit(2)
	}
	commands := [][]string{{gofmt, "-d", "internal/referenceslice"}, {goBin, "vet", "./internal/referenceslice"}, {goBin, "test", "-count=1", "./internal/referenceslice"}}
	for _, args := range commands {
		// #nosec G204,G702 -- executable paths are absolute, paired under one GOROOT,
		// and arguments are fixed constants rather than caller-controlled values.
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		cmd.Env = append(os.Environ(), "GOENV=off", "GOTOOLCHAIN=local", "GOPROXY=off", "GOSUMDB=off", "GOVCS=*:off", "GOWORK=off", "CGO_ENABLED=0", "GOFLAGS=-mod=readonly")
		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, "verification failed:", filepath.Base(args[0]))
			os.Exit(1)
		}
	}
}
