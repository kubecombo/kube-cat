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
var custompinglog = logf.Log.WithName("customping-resource")

// SetupCustomPingWebhookWithManager registers the webhook for CustomPing in the manager.
func SetupCustomPingWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&catv1.CustomPing{}).
		WithValidator(&CustomPingCustomValidator{}).
		WithDefaulter(&CustomPingCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-cat-kubecat-com-v1-customping,mutating=true,failurePolicy=fail,sideEffects=None,groups=cat.kubecat.com,resources=custompings,verbs=create;update,versions=v1,name=mcustomping-v1.kb.io,admissionReviewVersions=v1

// CustomPingCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind CustomPing when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type CustomPingCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &CustomPingCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind CustomPing.
func (d *CustomPingCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	customping, ok := obj.(*catv1.CustomPing)

	if !ok {
		return fmt.Errorf("expected an CustomPing object but got %T", obj)
	}
	custompinglog.Info("Defaulting for CustomPing", "name", customping.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-cat-kubecat-com-v1-customping,mutating=false,failurePolicy=fail,sideEffects=None,groups=cat.kubecat.com,resources=custompings,verbs=create;update,versions=v1,name=vcustomping-v1.kb.io,admissionReviewVersions=v1

// CustomPingCustomValidator struct is responsible for validating the CustomPing resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type CustomPingCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &CustomPingCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type CustomPing.
func (v *CustomPingCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	customping, ok := obj.(*catv1.CustomPing)
	if !ok {
		return nil, fmt.Errorf("expected a CustomPing object but got %T", obj)
	}
	custompinglog.Info("Validation for CustomPing upon creation", "name", customping.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type CustomPing.
func (v *CustomPingCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	customping, ok := newObj.(*catv1.CustomPing)
	if !ok {
		return nil, fmt.Errorf("expected a CustomPing object for the newObj but got %T", newObj)
	}
	custompinglog.Info("Validation for CustomPing upon update", "name", customping.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type CustomPing.
func (v *CustomPingCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	customping, ok := obj.(*catv1.CustomPing)
	if !ok {
		return nil, fmt.Errorf("expected a CustomPing object but got %T", obj)
	}
	custompinglog.Info("Validation for CustomPing upon deletion", "name", customping.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
