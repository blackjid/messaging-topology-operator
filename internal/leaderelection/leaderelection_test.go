package leaderelection_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rabbitmq/messaging-topology-operator/internal/leaderelection"
	corev1 "k8s.io/api/core/v1"
)

var _ = Describe("Leaderelection", func() {

	var pods []*corev1.Pod
	var numberOfPartitions int
	var numberOfConsumerSets int

	JustBeforeEach(func() {
		pods = generateTestPodSet(numberOfPartitions, numberOfConsumerSets)
	})

	When("passed as many partitions as consumer sets", func() {
		BeforeEach(func() {
			numberOfPartitions = 3
			numberOfConsumerSets = 3
		})
		It("distributes evenly", func() {
			leaderelection.Elect(pods)
			Expect(pods).To(BeBalanced())
		})
	})

	When("passed more partitions than consumer sets", func() {
		BeforeEach(func() {
			numberOfPartitions = 5
			numberOfConsumerSets = 3
		})
		It("distributes evenly", func() {
			leaderelection.Elect(pods)
			Expect(pods).To(BeBalanced())
		})

	})
	When("passed fewer partitions than consumer sets", func() {
		BeforeEach(func() {
			numberOfPartitions = 2
			numberOfConsumerSets = 7
		})
		It("distributes evenly", func() {
			leaderelection.Elect(pods)
			Expect(pods).To(BeBalanced())
		})
	})
})
