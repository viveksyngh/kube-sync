package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Validate(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		return fmt.Errorf("The resource name is missing")
	}

	if len(ctx.String("kubeconfig")) == 0 {
		return fmt.Errorf("Config for connecting with the Kubernetes cluster is missing. Either set KUBECONFIG environment variable or `--kubeconfig`")
	}

	if ctx.String("namespace") == ctx.String("target-namespace") {
		return fmt.Errorf("source and target namespace can not be the same")
	}

	return nil
}
