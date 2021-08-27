package syncer

import (
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func TestSecretSyncer(t *testing.T) {

	testcases := map[string]testcase{
		"secret missing from source namespace": {
			name:           "missing-secret",
			err:            fmt.Errorf(`secrets "%s" not found`, "missing-secret"),
			namespace:      "source-namespace",
			targetNamespce: "target-namespace",
			objects: []runtime.Object{&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "source-namespace"}},
				&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "target-namespace"}}},
		},
		"secret present in source namespace and missing in target namespace": {
			name:           "test-secret",
			err:            nil,
			namespace:      "source-namespace",
			targetNamespce: "target-namespace",
			objects: []runtime.Object{&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "source-namespace"}},
				&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "target-namespace"}},
				&corev1.Secret{ObjectMeta: metav1.ObjectMeta{Namespace: "source-namespace", Name: "test-secret"}, Data: map[string][]byte{"key": []byte("value")}},
			},
		},
		"secret present in both source and target namespace": {
			name:           "test-secret",
			err:            fmt.Errorf(`secret with name "%s" already exists in namepsace "%s"`, "test-secret", "target-namespace"),
			namespace:      "source-namespace",
			targetNamespce: "target-namespace",
			objects: []runtime.Object{&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "source-namespace"}},
				&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "target-namespace"}},
				&corev1.Secret{ObjectMeta: metav1.ObjectMeta{Namespace: "source-namespace", Name: "test-secret"}, Data: map[string][]byte{"key": []byte("value")}},
				&corev1.Secret{ObjectMeta: metav1.ObjectMeta{Namespace: "target-namespace", Name: "test-secret"}, Data: map[string][]byte{"key": []byte("value")}},
			},
		},
	}

	for name, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			fakeClient := fake.NewSimpleClientset(testcase.objects...)
			cmSyncer := SecretSyncer{}
			err := cmSyncer.Sync(fakeClient, testcase.name, testcase.namespace, testcase.targetNamespce)
			if testcase.err != nil {
				if testcase.err.Error() != err.Error() {
					t.Fatalf("testcase %s failed, got: %v, expected: %v", name, err, testcase.err)
				}
			} else {
				if err != nil {
					t.Fatalf("testcase %s failed, got: %v, expected: %v", name, err, testcase.err)
				}
			}
		})

	}

}
