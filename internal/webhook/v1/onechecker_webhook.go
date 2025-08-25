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
var onecheckerlog = logf.Log.WithName("onechecker-resource")

// SetupOneCheckerWebhookWithManager registers the webhook for OneChecker in the manager.
func SetupOneCheckerWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&catv1.OneChecker{}).
		WithValidator(&OneCheckerCustomValidator{}).
		WithDefaulter(&OneCheckerCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-cat-kubecat-com-v1-onechecker,mutating=true,failurePolicy=fail,sideEffects=None,groups=cat.kubecat.com,resources=onecheckers,verbs=create;update,versions=v1,name=monechecker-v1.kb.io,admissionReviewVersions=v1

// OneCheckerCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind OneChecker when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type OneCheckerCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &OneCheckerCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind OneChecker.
func (d *OneCheckerCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	onechecker, ok := obj.(*catv1.OneChecker)

	if !ok {
		return fmt.Errorf("expected an OneChecker object but got %T", obj)
	}
	onecheckerlog.Info("Defaulting for OneChecker", "name", onechecker.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-cat-kubecat-com-v1-onechecker,mutating=false,failurePolicy=fail,sideEffects=None,groups=cat.kubecat.com,resources=onecheckers,verbs=create;update,versions=v1,name=vonechecker-v1.kb.io,admissionReviewVersions=v1

// OneCheckerCustomValidator struct is responsible for validating the OneChecker resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type OneCheckerCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &OneCheckerCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type OneChecker.
func (v *OneCheckerCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	onechecker, ok := obj.(*catv1.OneChecker)
	if !ok {
		return nil, fmt.Errorf("expected a OneChecker object but got %T", obj)
	}
	onecheckerlog.Info("Validation for OneChecker upon creation", "name", onechecker.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type OneChecker.
func (v *OneCheckerCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	onechecker, ok := newObj.(*catv1.OneChecker)
	if !ok {
		return nil, fmt.Errorf("expected a OneChecker object for the newObj but got %T", newObj)
	}
	onecheckerlog.Info("Validation for OneChecker upon update", "name", onechecker.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type OneChecker.
func (v *OneCheckerCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	onechecker, ok := obj.(*catv1.OneChecker)
	if !ok {
		return nil, fmt.Errorf("expected a OneChecker object but got %T", obj)
	}
	onecheckerlog.Info("Validation for OneChecker upon deletion", "name", onechecker.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
