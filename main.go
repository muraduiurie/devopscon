package main

import (
	"context"
	devopsconv1 "devopscon/controller/api/v1"
	"github.com/go-logr/logr"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"k8s.io/apimachinery/pkg/runtime"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type WebReconciler struct {
	client.Client
	scheme *runtime.Scheme
	log    logr.Logger
}

var (
	scheme = runtime.NewScheme()
)

// initiate the program by creating the scheme
func init() {
	utilruntime.Must(devopsconv1.AddToScheme(scheme))
}

func main() {
	// create main logger
	logger := zap.New()
	ctrl.SetLogger(logger)
	log := ctrl.Log.WithName("main")
	log.Info("set up manager")

	// create manager
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
	})

	wr := WebReconciler{
		Client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
		log:    log.WithName("web-reconciler"),
	}

	// create new controller
	err = wr.SetupWithManager(mgr)
	if err != nil {
		log.Error(err, "unable to create controller")
		os.Exit(1)
	}

	// start manager
	ctx := ctrl.SetupSignalHandler()
	if err = mgr.Start(ctx); err != nil {
		log.Error(err, "problem running manager")
		os.Exit(1)
	}
}

// SetupWithManager should specify your resource explicitly
func (wr *WebReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devopsconv1.WebPage{}).
		Complete(wr)
}

func (wr *WebReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	wr.log.Info("reconciling webpage", "name", req.Name, "namespace", req.Namespace)

	wp := devopsconv1.WebPage{}

	err := wr.Client.Get(ctx, req.NamespacedName, &wp)
	if err != nil && kerr.IsNotFound(err) {
		return ctrl.Result{}, nil
	} else if err != nil {
		return ctrl.Result{}, err
	}

	// reconciliation logic
	// For example, you might want to update the status of the WebPage resource, or create/update a Deployment based on the WebPage spec.

	wr.log.Info("webpage reconciled")

	return ctrl.Result{}, nil
}
