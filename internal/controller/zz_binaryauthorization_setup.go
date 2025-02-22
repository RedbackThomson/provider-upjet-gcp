package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	attestor "github.com/upbound/provider-gcp/internal/controller/binaryauthorization/attestor"
	policy "github.com/upbound/provider-gcp/internal/controller/binaryauthorization/policy"
)

// Setup_binaryauthorization creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_binaryauthorization(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attestor.Setup,
		policy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
