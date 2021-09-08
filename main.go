package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func loadEnv() {
	file, err := ioutil.ReadFile("config/config.txt")
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, "=")
		os.Setenv(split[0], split[1])
	}
}

func main() {
	loadEnv()
	StartHttpServer()
}
