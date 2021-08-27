package syncer

import (
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

type testcase struct {
	objects        []runtime.Object
	err            error
	name           string
	namespace      string
	targetNamespce string
}

func TestConfigMapSyncer(t *testing.T) {

	testcases := map[string]testcase{
		"configmap missing from source namespace": {
			name:           "missing-configmap",
			err:            fmt.Errorf(`configmaps "%s" not found`, "missing-configmap"),
			namespace:      "source-namespace",
			targetNamespce: "target-namespace",
			objects: []runtime.Object{&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "source-namespace"}},
				&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "target-namespace"}}},
		},
		"configmap present in source namespace and missing in target namespace": {
			name:           "test-configmap",
			err:            nil,
			namespace:      "source-namespace",
			targetNamespce: "target-namespace",
			objects: []runtime.Object{&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "source-namespace"}},
				&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "target-namespace"}},
				&corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Namespace: "source-namespace", Name: "test-configmap"}, Data: map[string]string{"key": "value"}},
			},
		},
		"configmap present in both source and target namespace": {
			name:           "test-configmap",
			err:            fmt.Errorf(`configmap with name "%s" already exists in namepsace "%s"`, "test-configmap", "target-namespace"),
			namespace:      "source-namespace",
			targetNamespce: "target-namespace",
			objects: []runtime.Object{&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "source-namespace"}},
				&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "target-namespace"}},
				&corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Namespace: "source-namespace", Name: "test-configmap"}, Data: map[string]string{"key": "value"}},
				&corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Namespace: "target-namespace", Name: "test-configmap"}, Data: map[string]string{"key": "value"}},
			},
		},
	}

	for name, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			fakeClient := fake.NewSimpleClientset(testcase.objects...)
			cmSyncer := ConfigMapSyncer{}
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
