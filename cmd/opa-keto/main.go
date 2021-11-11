package main

import (
	"fmt"
	"github.com/open-policy-agent/opa/cmd"
	opa_keto "github.com/swisscom/opa-keto/pkg"
	"os"
)

func main() {
	opa_keto.Init()
	opa_keto.RegisterCheck()
	opa_keto.RegisterExpand()
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
