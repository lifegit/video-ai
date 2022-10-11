package main

import (
	"github.com/lifegit/video/cmd"
	"log"
)

//go:generate go build -o video -ldflags "-s -w"
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
