package syncer

import (
	"k8s.io/client-go/kubernetes"
)

type Syncer interface {
	Sync(kubernetes.Interface, string, string, string) error
}
