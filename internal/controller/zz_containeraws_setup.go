package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/upbound/provider-gcp/internal/controller/containeraws/cluster"
	nodepool "github.com/upbound/provider-gcp/internal/controller/containeraws/nodepool"
)

// Setup_containeraws creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containeraws(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		nodepool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
