package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	environment "github.com/upbound/provider-gcp/internal/controller/composer/environment"
)

// Setup_composer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_composer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		environment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
