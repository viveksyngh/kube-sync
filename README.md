# kube-sync
Kube Sync is a CLI application to copy/sync configmaps and secrets from one namespace to another.

# Motivation

While working with kubernetes, all of us might have encountered a situation where applications deployed in a new namespace were failing because of missing image pull secret, which we forgot to create in the new namespace. This tool help us in copying such secrets or configmaps from already existing namespace to a new namespace.

# Installation

Stable binaries are also available on the [releases](https://github.com/viveksyngh/kube-sync/releases) page. Stable releases are generally recommended for CI usage in particular. To install, download the binary for your platform from "Assets" and place this into your $PATH:

## For Linux:
```
curl -Lo ./kube-sync https://github.com/viveksyngh/kube-sync/releases/download/v0.1.0/kube-sync-$(uname)-amd64
```
```
chmod +x ./kube-sync
```
```
mv ./kube-sync /usr/local/bin/kube-sync
```

## On macOS via Bash:
```
curl -Lo ./kube-sync https://github.com/viveksyngh/kube-sync/releases/download/v0.1.0/kube-sync-darwin-amd64
```
```
chmod +x ./kube-sync
```
```
mv ./kube-sync /usr/local/bin/kube-sync
```
## On Windows:
```
curl.exe -Lo kube-sync-windows-amd64.exe https://github.com/viveksyngh/kube-sync/releases/download/v0.1.0/kube-sync-windows-amd64
```
```
Move-Item .\kube-sync-windows-amd64.exe  c:\some-dir-in-your-PATH\kube-sync-windows-amd64.exe 
```
# Usage

```
NAME:
   kube-sync - A CLI to sync kubernetes resources

USAGE:
   kube-sync [global options] command [command options] [arguments...]

DESCRIPTION:
   A CLI to sync resources in a kubernetes cluster

COMMANDS:
   configmap, cm  Copy a configmap from one namespace to another
   secret, sec    Copy a secret from one namespace to another
   version, v     print version
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --kubeconfig value                    path to the kubeconfig file to use for CLI requests (default: "/Users/svivekkumar/.kube/config") [$KUBECONFIG]
   --namespace value, -n value           namespace of the source resource (default: "default")
   --target-namespace value, --tn value  target namespace in which resource will be copied (default: "default")
   --help, -h                            show help (default: false)
```

### Sync/Copy a configmap

```
kube-sync configmap --help
NAME:
   kube-sync configmap - Copy a configmap from one namespace to another

USAGE:
   kube-sync configmap [command options] NAME

OPTIONS:
   --help, -h  show help (default: false)
```

```
kube-sync --namespace source-ns --target-namespace target-ns configmap configmap-name

kube-sync --namespace source-ns --target-namespace target-ns cm configmap-name
```

### Sync/Copy a secret
```
kube-sync secret --help
NAME:
   kube-sync secret - Copy a secret from one namespace to another

USAGE:
   kube-sync secret [command options] NAME

OPTIONS:
   --help, -h  show help (default: false)
```

```
kube-sync --namespace source-ns --target-namespace target-ns secret secret-name

kube-sync --namespace source-ns --target-namespace target-ns sec secret-name
```
