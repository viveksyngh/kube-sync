package syncer

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapSyncer struct {
}

var _ Syncer = &ConfigMapSyncer{}

func (cm *ConfigMapSyncer) Sync(client *kubernetes.Clientset, name, namespace, targetNamespace string) error {
	configmap, err := client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	_, err = client.CoreV1().ConfigMaps(targetNamespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err == nil {
		return fmt.Errorf("configmap with name `%s` already exists in namepsace `%s`", name, targetNamespace)
	}

	if apierrors.IsNotFound(err) {
		targetConfigmap := &corev1.ConfigMap{
			TypeMeta: configmap.TypeMeta,
			ObjectMeta: metav1.ObjectMeta{
				Name:        name,
				Namespace:   targetNamespace,
				Labels:      configmap.Labels,
				Annotations: configmap.Annotations,
			},
			Data:       configmap.Data,
			BinaryData: configmap.BinaryData,
		}

		_, err = client.CoreV1().ConfigMaps(targetNamespace).Create(context.TODO(), targetConfigmap, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		return nil
	}

	return err
}
