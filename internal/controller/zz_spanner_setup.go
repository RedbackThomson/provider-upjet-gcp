package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	database "github.com/upbound/provider-gcp/internal/controller/spanner/database"
	databaseiammember "github.com/upbound/provider-gcp/internal/controller/spanner/databaseiammember"
	instance "github.com/upbound/provider-gcp/internal/controller/spanner/instance"
	instanceiammember "github.com/upbound/provider-gcp/internal/controller/spanner/instanceiammember"
)

// Setup_spanner creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_spanner(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		databaseiammember.Setup,
		instance.Setup,
		instanceiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
