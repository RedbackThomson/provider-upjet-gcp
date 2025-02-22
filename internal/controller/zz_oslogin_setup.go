package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	sshpublickey "github.com/upbound/provider-gcp/internal/controller/oslogin/sshpublickey"
)

// Setup_oslogin creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_oslogin(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		sshpublickey.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
