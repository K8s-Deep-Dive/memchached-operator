package controllers

import (
	"context"
	"fmt"
	"github.com/example/memcached-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

// TODO-mmc:  Need to finish setup of tests ...
//  Need to write stubs.  Currently this will call Reconcile
var _ = Describe("Namespecial controller", func() {

	// Define utility constants for object names and testing timeouts/durations and intervals.
	const (
		jobName      = "stuff"
		jobNamespace = "default"
		JobName      = "test-job"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When Deleting Namespace `stuff`", func() {
		It("Should recreate ", func() {
			By("By creating a new namespace")

			ctx := context.Background()

			cr := &v1alpha1.Memcached{}
			cr.Kind = "Memcached"
			cr.Namespace = "default"
			cr.Name = "memcached-sample"
			cr.Spec = v1alpha1.MemcachedSpec{}
			cr.Spec.Size = 3

			err := k8sClient.Create(ctx, cr)
			fmt.Printf("err: %v\n", err)
			cr.Spec.Size = 2
			k8sClient.Update(ctx, cr)

			// We'll need to retry getting this, given that creation may not immediately happen.
			Eventually(func() bool {
				err := k8sClient.Delete(ctx, cr)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue()) // TODO-mmc: put in the correct value and set to true

		})
	})
})
