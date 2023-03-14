/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	servicesv1 "github.com/SAP/sap-btp-service-operator/api/v1"
	"github.com/google/uuid"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ctrl "sigs.k8s.io/controller-runtime"
)

// SharedServiceInstanceReconciler reconciles a SharedServiceInstance object
type SharedServiceInstanceReconciler struct {
	*BaseReconciler
}

//+kubebuilder:rbac:groups=services.cloud.sap.com,resources=sharedserviceinstances,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=services.cloud.sap.com,resources=sharedserviceinstances/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=services.cloud.sap.com,resources=sharedserviceinstances/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SharedServiceInstance object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *SharedServiceInstanceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("sharedserviceinstance", req.NamespacedName).WithValues("correlation_id", uuid.New().String())
	ctx = context.WithValue(ctx, LogKey{}, log)

	serviceInstance := &servicesv1.ServiceInstance{}
	if err := r.Get(ctx, req.NamespacedName, serviceInstance); err != nil {
		if !apierrors.IsNotFound(err) {
			log.Error(err, "unable to fetch ServiceInstance")
		}
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SharedServiceInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&servicesv1.SharedServiceInstance{}).
		Complete(r)
}
