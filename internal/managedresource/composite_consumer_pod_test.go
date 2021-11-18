package managedresource_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	topology "github.com/rabbitmq/messaging-topology-operator/api/v1beta1"
	"github.com/rabbitmq/messaging-topology-operator/internal/managedresource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var _ = Describe("SuperstreamExchange", func() {
	var (
		builder                     *managedresource.Builder
		compositeConsumerSet *topology.CompositeConsumerSet
		compositeConsumerPodBuilder *managedresource.CompositeConsumerPodBuilder
		pod                     *corev1.Pod
		podSpec                     corev1.PodSpec
		scheme                      *runtime.Scheme
	)

	BeforeEach(func() {
		scheme = runtime.NewScheme()
		Expect(topology.AddToScheme(scheme)).To(Succeed())
		compositeConsumerSet = &topology.CompositeConsumerSet{}
		compositeConsumerSet.Name = "parent-set"
		compositeConsumerSet.Namespace = "parent-namespace"
		compositeConsumerSet.Spec.SuperStreamReference = topology.SuperStreamReference{
			Name:      "super-stream-1",
			Namespace: "parent-namespace",
		}

		podSpec = corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "example-container",
					Image: "example-image",
				},
			},
		}

		builder = &managedresource.Builder{
			ObjectOwner: compositeConsumerSet,
			Scheme:      scheme,
		}
		compositeConsumerPodBuilder = builder.CompositeConsumerPod(podSpec, "sample-partition", 4)
		obj, _ := compositeConsumerPodBuilder.Build()
		pod = obj.(*corev1.Pod)
	})

	Context("Build", func() {
		It("generates an exchange object with the correct name", func() {
			Expect(pod.Name).To(Equal("parent-set-sample-partition-4"))
		})

		It("generates an pod object with the correct namespace", func() {
			Expect(pod.Namespace).To(Equal(compositeConsumerSet.Namespace))
		})
	})

	Context("Update", func() {
		BeforeEach(func() {
			Expect(compositeConsumerPodBuilder.Update(pod)).To(Succeed())
		})
		It("sets owner reference", func() {
			Expect(pod.OwnerReferences[0].Name).To(Equal(compositeConsumerSet.Name))
		})
		It("sets expected labels on the Pod", func() {
			Expect(pod.ObjectMeta.Labels).To(HaveKeyWithValue("rabbitmq.com/super-stream-partition", "sample-partition"))
			Expect(pod.ObjectMeta.Labels).To(HaveKeyWithValue("rabbitmq.com/composite-consumer-replica", "4"))
		})

	})
})
