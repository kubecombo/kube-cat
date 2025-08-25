/*
Copyright 2025 kubecat.

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

package v1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	catv1 "github.com/kubecombo/kube-cat/api/v1"
)

// nolint:unused
// log is for logging in this package.
var clustercheckerlog = logf.Log.WithName("clusterchecker-resource")

// SetupClusterCheckerWebhookWithManager registers the webhook for ClusterChecker in the manager.
func SetupClusterCheckerWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&catv1.ClusterChecker{}).
		WithValidator(&ClusterCheckerCustomValidator{}).
		WithDefaulter(&ClusterCheckerCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-cat-kubecat-com-v1-clusterchecker,mutating=true,failurePolicy=fail,sideEffects=None,groups=cat.kubecat.com,resources=clustercheckers,verbs=create;update,versions=v1,name=mclusterchecker-v1.kb.io,admissionReviewVersions=v1

// ClusterCheckerCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind ClusterChecker when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type ClusterCheckerCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &ClusterCheckerCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind ClusterChecker.
func (d *ClusterCheckerCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	clusterchecker, ok := obj.(*catv1.ClusterChecker)

	if !ok {
		return fmt.Errorf("expected an ClusterChecker object but got %T", obj)
	}
	clustercheckerlog.Info("Defaulting for ClusterChecker", "name", clusterchecker.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-cat-kubecat-com-v1-clusterchecker,mutating=false,failurePolicy=fail,sideEffects=None,groups=cat.kubecat.com,resources=clustercheckers,verbs=create;update,versions=v1,name=vclusterchecker-v1.kb.io,admissionReviewVersions=v1

// ClusterCheckerCustomValidator struct is responsible for validating the ClusterChecker resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type ClusterCheckerCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &ClusterCheckerCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type ClusterChecker.
func (v *ClusterCheckerCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	clusterchecker, ok := obj.(*catv1.ClusterChecker)
	if !ok {
		return nil, fmt.Errorf("expected a ClusterChecker object but got %T", obj)
	}
	clustercheckerlog.Info("Validation for ClusterChecker upon creation", "name", clusterchecker.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type ClusterChecker.
func (v *ClusterCheckerCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	clusterchecker, ok := newObj.(*catv1.ClusterChecker)
	if !ok {
		return nil, fmt.Errorf("expected a ClusterChecker object for the newObj but got %T", newObj)
	}
	clustercheckerlog.Info("Validation for ClusterChecker upon update", "name", clusterchecker.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type ClusterChecker.
func (v *ClusterCheckerCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	clusterchecker, ok := obj.(*catv1.ClusterChecker)
	if !ok {
		return nil, fmt.Errorf("expected a ClusterChecker object but got %T", obj)
	}
	clustercheckerlog.Info("Validation for ClusterChecker upon deletion", "name", clusterchecker.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
