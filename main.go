package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/viveksyngh/kube-copy/cmd"
	"k8s.io/client-go/util/homedir"
)

var (
	kubeconfig      *string
	namespace       *string
	targetNamespace *string
)

func defaultKubeconfig() string {
	if home := homedir.HomeDir(); home != "" {
		return filepath.Join(home, ".kube", "config")
	}
	return ""
}

func main() {
	app := &cli.App{
		Name:     "kube-copy",
		HelpName: "kube-copy",
		Commands: []*cli.Command{
			{
				Name:      "configmap",
				Aliases:   []string{"cm"},
				ArgsUsage: "NAME",
				Usage:     "Copy a configmap from one namespace to another",
				Action: func(c *cli.Context) error {
					err := cmd.Validate(c)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:      "secret",
				Aliases:   []string{"sec"},
				ArgsUsage: "NAME",
				Usage:     "Copy a secret from one namespace to another",
				Action: func(c *cli.Context) error {
					err := cmd.Validate(c)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "kubeconfig",
				Usage:       "path to the kubeconfig file to use for CLI requests",
				EnvVars:     []string{"KUBECONFIG"},
				Destination: kubeconfig,
				Value:       defaultKubeconfig(),
			},
			&cli.StringFlag{
				Name:        "namespace",
				Aliases:     []string{"n"},
				Value:       "default",
				Usage:       "namespace of the source resource",
				Destination: namespace,
			},
			&cli.StringFlag{
				Name:        "target-namespace",
				Aliases:     []string{"tn"},
				Value:       "default",
				Usage:       "target namespace in which resource will be copied",
				Destination: targetNamespace,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
