package v1beta1

import (
	"fmt"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (s *SuperStreamConsumer) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(s).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-rabbitmq-com-v1beta1-superstreamconsumer,mutating=false,failurePolicy=fail,groups=rabbitmq.com,resources=superstreamconsumer,versions=v1beta1,name=vsuperstreamconsumer.kb.io,sideEffects=none,admissionReviewVersions=v1

var _ webhook.Validator = &SuperStreamConsumer{}

// no validation on create
func (s *SuperStreamConsumer) ValidateCreate() error {
	return nil
}

// returns error type 'forbidden' for updates on superstream name and rabbitmqClusterReference
func (s *SuperStreamConsumer) ValidateUpdate(old runtime.Object) error {
	oldSuperStreamConsumer, ok := old.(*SuperStreamConsumer)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("expected a superstream but got a %T", old))
	}

	detailMsg := "updates on superStreamReference and consumerPodSpec are forbidden"

	if s.Spec.SuperStreamReference != oldSuperStreamConsumer.Spec.SuperStreamReference {
		return apierrors.NewForbidden(s.GroupResource(), s.Name,
			field.Forbidden(field.NewPath("spec", "superStreamReference"), detailMsg))
	}
	return nil
}

// ValidateDelete no validation on delete
func (s *SuperStreamConsumer) ValidateDelete() error {
	return nil
}