package syncer

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretSyncer struct {
}

var _ Syncer = &SecretSyncer{}

func (cm *SecretSyncer) Sync(client kubernetes.Interface, name, namespace, targetNamespace string) error {
	secret, err := client.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	_, err = client.CoreV1().Secrets(targetNamespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err == nil {
		return fmt.Errorf(`secret with name "%s" already exists in namepsace "%s"`, name, targetNamespace)
	}

	if apierrors.IsNotFound(err) {
		targetSecret := &corev1.Secret{
			TypeMeta: secret.TypeMeta,
			ObjectMeta: metav1.ObjectMeta{
				Name:        secret.Name,
				Namespace:   targetNamespace,
				Labels:      secret.Labels,
				Annotations: secret.Annotations,
			},
			Immutable:  secret.Immutable,
			StringData: secret.StringData,
			Data:       secret.Data,
			Type:       secret.Type,
		}
		_, err = client.CoreV1().Secrets(targetNamespace).Create(context.TODO(), targetSecret, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		return nil
	}

	return err
}
