package syncer

import "k8s.io/client-go/kubernetes"

type ConfigMapSyncer struct {
}

var _ Syncer = &ConfigMapSyncer{}

func (cm *ConfigMapSyncer) Sync(client *kubernetes.Clientset, name, namespace, targetNamespace string) error {
	return nil
}
