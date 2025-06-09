package main

import (
	"context"
	devopsconv1 "devopscon/controller/api/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (w *WebReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx).WithValues("webpage", req.NamespacedName)

	wp := devopsconv1.WebPage{}

	err := w.Client.Get(ctx, req.NamespacedName, &wp)
	if err != nil && kerr.IsNotFound(err) {
		return ctrl.Result{}, nil
	} else if err != nil {
		return ctrl.Result{}, err
	}

	// reconciliation logic
	// For example, you might want to update the status of the WebPage resource, or create/update a Deployment based on the WebPage spec.

	log.Info("webpage reconciled")

	return ctrl.Result{}, nil
}
