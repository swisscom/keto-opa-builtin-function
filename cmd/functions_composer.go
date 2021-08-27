package cmd

import (
	"fmt"
	"github.com/open-policy-agent/opa/cmd"
	"keto_opa_function/pkg"
	"os"
)

func Run() {
	pkg.RegisterCheckFuncs()
	pkg.RegisterExpandFuncs()
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}