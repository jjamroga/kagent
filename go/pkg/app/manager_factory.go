package app

import (
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type ManagerFactory interface {
	CreateManager(controllerConfig *rest.Config, options manager.Options) (ctrl.Manager, error)
}

type managerFactory struct{}

func NewManagerFactory() ManagerFactory {
	return &managerFactory{}
}

func (m *managerFactory) CreateManager(controllerConfig *rest.Config, options manager.Options) (ctrl.Manager, error) {
	return ctrl.NewManager(controllerConfig, options)
}
