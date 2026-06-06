package testhelpers

import (
	"context"

	portainer "github.com/opendocking/opendocking/api"
	models "github.com/opendocking/opendocking/api/http/models/kubernetes"
)

type testKubeClient struct {
	portainer.KubeClient
}

func NewKubernetesClient() portainer.KubeClient {
	return &testKubeClient{}
}

// Event
func (kcl *testKubeClient) GetEvents(namespace string, resourceId string) ([]models.K8sEvent, error) {
	return nil, nil
}

// Pod
func (kcl *testKubeClient) DeletePod(namespace, name string) error             { return nil }
func (kcl *testKubeClient) RestartPod(namespace, name string) error            { return nil }
func (kcl *testKubeClient) SupportsPodRestart(_ context.Context) (bool, error) { return false, nil }
