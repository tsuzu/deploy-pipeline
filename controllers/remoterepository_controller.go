/*
Copyright 2019 Tsuzu.

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

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	trialv1alpha1 "github.com/cs3238-tsuzu/deploy-pipeline/api/v1alpha1"
)

// RemoteRepositoryReconciler reconciles a RemoteRepository object
type RemoteRepositoryReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=trial.tsuzu.xyz,resources=remoterepositories,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=trial.tsuzu.xyz,resources=remoterepositories/status,verbs=get;update;patch

func (r *RemoteRepositoryReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("remoterepository", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *RemoteRepositoryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&trialv1alpha1.RemoteRepository{}).
		Complete(r)
}
