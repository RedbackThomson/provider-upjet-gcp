package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	folderbucketconfig "github.com/upbound/provider-gcp/internal/controller/logging/folderbucketconfig"
	folderexclusion "github.com/upbound/provider-gcp/internal/controller/logging/folderexclusion"
	foldersink "github.com/upbound/provider-gcp/internal/controller/logging/foldersink"
	logview "github.com/upbound/provider-gcp/internal/controller/logging/logview"
	metric "github.com/upbound/provider-gcp/internal/controller/logging/metric"
	projectbucketconfig "github.com/upbound/provider-gcp/internal/controller/logging/projectbucketconfig"
	projectexclusion "github.com/upbound/provider-gcp/internal/controller/logging/projectexclusion"
	projectsink "github.com/upbound/provider-gcp/internal/controller/logging/projectsink"
)

// Setup_logging creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_logging(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		folderbucketconfig.Setup,
		folderexclusion.Setup,
		foldersink.Setup,
		logview.Setup,
		metric.Setup,
		projectbucketconfig.Setup,
		projectexclusion.Setup,
		projectsink.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
