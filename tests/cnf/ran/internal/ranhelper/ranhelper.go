package ranhelper

import (
	"github.com/golang/glog"
	"github.com/openshift-kni/eco-goinfra/pkg/clients"
	"github.com/openshift-kni/eco-goinfra/pkg/pod"
	"github.com/openshift-kni/eco-gotests/tests/cnf/ran/internal/ranparam"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// DoesContainerExistInPod checks if a given container exists in a given pod.
func DoesContainerExistInPod(pod *pod.Builder, containerName string) bool {
	containers := pod.Object.Status.ContainerStatuses

	for _, container := range containers {
		if container.Name == containerName {
			glog.V(ranparam.LogLevel).Infof("found %s container", containerName)

			return true
		}
	}

	return false
}

// AreClustersPresent checks all of the provided clusters and returns false if any are nil.
func AreClustersPresent(clusters []*clients.Settings) bool {
	for _, cluster := range clusters {
		if cluster == nil {
			return false
		}
	}

	return true
}

// UnmarshalRaw converts raw bytes for a K8s CR into the actual type.
func UnmarshalRaw[T any](raw []byte) (*T, error) {
	untyped := &unstructured.Unstructured{}
	err := untyped.UnmarshalJSON(raw)

	if err != nil {
		return nil, err
	}

	var typed T
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(untyped.UnstructuredContent(), &typed)

	if err != nil {
		return nil, err
	}

	return &typed, nil
}
