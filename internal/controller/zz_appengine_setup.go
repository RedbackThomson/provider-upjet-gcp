package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/upbound/provider-gcp/internal/controller/appengine/application"
	applicationurldispatchrules "github.com/upbound/provider-gcp/internal/controller/appengine/applicationurldispatchrules"
	firewallrule "github.com/upbound/provider-gcp/internal/controller/appengine/firewallrule"
	servicenetworksettings "github.com/upbound/provider-gcp/internal/controller/appengine/servicenetworksettings"
	standardappversion "github.com/upbound/provider-gcp/internal/controller/appengine/standardappversion"
)

// Setup_appengine creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appengine(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		applicationurldispatchrules.Setup,
		firewallrule.Setup,
		servicenetworksettings.Setup,
		standardappversion.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
