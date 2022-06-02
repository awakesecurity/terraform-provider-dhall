package main

import (
	"flag"

	"github.com/awakesecurity/terraform-provider-dhall/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: provider.New}

	if debugMode {
		opts.Debug = true
		opts.ProviderAddr = "registry.terraform.io/awakesecurity/dhall"
	}

	plugin.Serve(opts)
}
