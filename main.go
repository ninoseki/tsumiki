package main

import (
	"fmt"
	"os"

	"github.com/ninoseki/tsumiki/cmd"

	_ "github.com/ninoseki/tsumiki/functions"
)

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
