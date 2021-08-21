package syncer

import (
	"k8s.io/client-go/kubernetes"
)

type Syncer interface {
	Sync(*kubernetes.Clientset, string, string, string) error
}
