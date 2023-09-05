package v1alpha1_jobservice

import (
	neosync_k8sclient "github.com/nucleuscloud/neosync/backend/internal/k8s/client"
	v1alpha1_connectionservice "github.com/nucleuscloud/neosync/backend/services/mgmt/v1alpha1/connection-service"
)

type Service struct {
	cfg               *Config
	k8sclient         *neosync_k8sclient.Client
	connectionService *v1alpha1_connectionservice.Service
}

type Config struct {
	JobConfigNamespace string
}

func New(
	cfg *Config,
	k8sclient *neosync_k8sclient.Client,
	connectionService *v1alpha1_connectionservice.Service,
) *Service {

	return &Service{
		cfg:               cfg,
		k8sclient:         k8sclient,
		connectionService: connectionService,
	}
}
