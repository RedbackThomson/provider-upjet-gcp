package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	node "github.com/upbound/provider-gcp/internal/controller/tpu/node"
)

// Setup_tpu creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_tpu(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		node.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
