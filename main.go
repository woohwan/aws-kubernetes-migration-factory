package main

import (
	"os"
)


type Config struct {
	CurrentContext string `yaml:"current-context"`
}

var config Config

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func main() {



}