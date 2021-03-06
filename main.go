package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/viveksyngh/kube-sync/cmd"
	"github.com/viveksyngh/kube-sync/pkg/syncer"
	"github.com/viveksyngh/kube-sync/pkg/version"
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
		Name:        "kube-sync",
		Usage:       "A CLI to sync Kubernetes resources",
		Description: "A CLI to sync resources in a Kubernetes cluster",
		HelpName:    "kube-sync",
		Commands: []*cli.Command{
			{
				Name:      "configmap",
				Aliases:   []string{"cm"},
				ArgsUsage: "NAME",
				Usage:     "Copy a ConfigMap from one namespace to another",
				Action: func(c *cli.Context) error {
					err := cmd.Validate(c)
					if err != nil {
						return err
					}
					client, err := cmd.CreateClientset(c.String("kubeconfig"))
					if err != nil {
						return err
					}
					cs := &syncer.ConfigMapSyncer{}
					return cs.Sync(client, c.Args().First(), c.String("namespace"), c.String("target-namespace"))
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
					client, err := cmd.CreateClientset(c.String("kubeconfig"))
					if err != nil {
						return err
					}
					cs := &syncer.SecretSyncer{}
					return cs.Sync(client, c.Args().First(), c.String("namespace"), c.String("target-namespace"))
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "print version",
				Action: func(c *cli.Context) error {
					fmt.Printf("Version: %s\n", version.Version)
					fmt.Printf("SHA: %s\n", version.GitCommit)
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
				Usage:       "target namespace where resource will be copied",
				Destination: targetNamespace,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
