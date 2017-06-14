package broker_test

import (
	"onboarding_broker/broker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-cf/brokerapi"
)

var _ = Describe("Broker", func() {

	Describe("Services()", func() {
		It("should return the redis service", func() {
			broker := broker.Broker{}
			plan := ServicePlan{ID: "base", Name: "base"}
			service := Service{ID: "redis", Name: "redis", Bindable: true, Plans: []ServicePlan{plan}}
			Expect(broker.Services(nil)).To(Equal([]Service{service}))
		})
	})

	Describe("Provision()", func() {
		It("should return the provisioned service spec", func() {
			broker := broker.Broker{}
			provisionedServiceSpec, err := broker.Provision(nil, "", ProvisionDetails{}, true)
			Expect(err).ToNot(HaveOccurred())
			Expect(provisionedServiceSpec).To(Equal(ProvisionedServiceSpec{}))
		})
	})

	Describe("Deprovision", func() {
		It("should return the deprovioned service spec", func() {
			broker := broker.Broker{}
			Expect(func() { _, _ = broker.Deprovision(nil, "", DeprovisionDetails{}, true) }).To(Panic())
		})
	})

	Describe("Bind", func() {
		It("should return the binding", func() {
			broker := broker.Broker{}
			Expect(func() { _, _ = broker.Bind(nil, "", "", BindDetails{}) }).To(Panic())
		})
	})

	Describe("Unbind", func() {
		It("should not result in an error", func() {
			broker := broker.Broker{}
			Expect(func() { _ = broker.Unbind(nil, "", "", UnbindDetails{}) }).To(Panic())
		})
	})

	Describe("Update", func() {
		It("should return an update service spec", func() {
			broker := broker.Broker{}
			Expect(func() { _, _ = broker.Update(nil, "", UpdateDetails{}, true) }).Should(Panic())
		})
	})

	Describe("Last Operation", func() {
		It("should return an update service spec", func() {
			broker := broker.Broker{}
			Expect(func() { _, _ = broker.LastOperation(nil, "", "") }).To(Panic())
		})
	})
})
