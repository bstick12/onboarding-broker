package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOnboardingBroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OnboardingBroker Suite")
}
