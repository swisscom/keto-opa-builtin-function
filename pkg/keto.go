package opa_keto

import (
	"github.com/ory/keto-client-go/client"
	"os"
)

var cfg = &client.TransportConfig{
	Host:     "localhost:4466",
	BasePath: "/",
	Schemes:  []string{"http"},
}

func Init() {
	if ketoUrl, exists := os.LookupEnv(KetoUrlEnv); exists {
		cfg.Host = ketoUrl
	}
}
