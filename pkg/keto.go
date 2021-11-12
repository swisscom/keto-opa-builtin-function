package opa_keto

import (
	"github.com/ory/keto-client-go/client"
	"net/url"
	"os"
)

var cfg = &client.TransportConfig{
	Host:     "localhost:4466",
	BasePath: "/",
	Schemes:  []string{"http"},
}

func Init() {
	if ketoUrl, exists := os.LookupEnv(KetoUrlEnv); exists {
		u, err := url.Parse(ketoUrl)
		if err != nil {
			panic(err)
		}
		cfg.Host = u.Host
		cfg.BasePath = u.Path
		cfg.Schemes = []string{u.Scheme}
	}
}
