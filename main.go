package main

import (
	devopsconv1 "devopscon/controller/api/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"k8s.io/apimachinery/pkg/runtime"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type WebReconciler struct {
	client.Client
	scheme *runtime.Scheme
}

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

// initiate the program by creating the scheme
func init() {
	utilruntime.Must(devopsconv1.AddToScheme(scheme))
}

func main() {
	// create manager
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
	})

	wr := WebReconciler{
		Client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
	}

	// create new controller
	err = wr.SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller")
		os.Exit(1)
	}

	// start manager
	ctx := ctrl.SetupSignalHandler()
	if err = mgr.Start(ctx); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func (w *WebReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devopsconv1.WebPage{}).
		Complete(w)
}
