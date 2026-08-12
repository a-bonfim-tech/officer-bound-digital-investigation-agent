package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] != "all" {
		fmt.Fprintln(os.Stderr, "usage: verify all")
		os.Exit(2)
	}
	goBin, gofmt := os.Getenv("OBDIA_GO_BIN"), os.Getenv("OBDIA_GOFMT_BIN")
	if goBin == "" {
		goBin = filepath.Join(runtime.GOROOT(), "bin", "go")
	}
	if gofmt == "" {
		gofmt = filepath.Join(runtime.GOROOT(), "bin", "gofmt")
	}
	if !filepath.IsAbs(goBin) || !filepath.IsAbs(gofmt) {
		fmt.Fprintln(os.Stderr, "absolute Go tool paths required")
		os.Exit(2)
	}
	goRoot := filepath.Dir(filepath.Dir(goBin))
	if filepath.Clean(gofmt) != filepath.Join(goRoot, "bin", "gofmt") || filepath.Clean(goBin) != filepath.Join(goRoot, "bin", "go") {
		fmt.Fprintln(os.Stderr, "Go tools must be the go and gofmt pair from one absolute GOROOT")
		os.Exit(2)
	}
	version, err := exec.Command(goBin, "version").Output() // #nosec G204,G702 -- validated absolute binary under runtime GOROOT.
	if err != nil || !strings.HasPrefix(string(version), "go version go1.26.5 ") {
		fmt.Fprintln(os.Stderr, "exact Go 1.26.5 toolchain required")
		os.Exit(1)
	}
	commands := [][]string{{gofmt, "-d", "cmd", "internal"}, {goBin, "vet", "./..."}, {goBin, "test", "-count=1", "./..."}}
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
