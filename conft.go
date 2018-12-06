package main

import (
	"flag"
	"fmt"
	"os"
	"plugin"
)

// Backend interface
type Backend interface {
	Client()
}

func main() {
	backendPtr := flag.String("backend", "env", "backend")

	flag.Parse()

	var mod string
	mod = "./plugins/" + *backendPtr + ".so"

	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symBackend, err := plug.Lookup("Backend")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var backend Backend
	backend, ok := symBackend.(Backend)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	backend.Client()

}
