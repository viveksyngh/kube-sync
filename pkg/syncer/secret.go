package syncer

import "k8s.io/client-go/kubernetes"

type SecretSyncer struct {
}

var _ Syncer = &SecretSyncer{}

func (cm *SecretSyncer) Sync(client *kubernetes.Clientset, name, namespace, targetNamespace string) error {

	return nil
}
