package cmd

import (
	ocmd "github.com/open-policy-agent/opa/cmd"
)

var RootCommand = ocmd.RootCommand

func init() {
	RootCommand.Short = "Open Policy Agent (OPA)"
	RootCommand.Long = `Open Policy Agent with extra built-in functions`
}
