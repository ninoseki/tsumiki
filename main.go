package main

import (
	"fmt"
	"os"

	"github.com/open-policy-agent/opa/cmd"

	_ "github.com/ninoseki/tsumiki/functions"
)

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
