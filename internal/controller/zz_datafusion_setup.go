package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	instance "github.com/upbound/provider-gcp/internal/controller/datafusion/instance"
)

// Setup_datafusion creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datafusion(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
