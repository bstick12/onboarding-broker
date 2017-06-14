package main_test

import (
	"fmt"
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var pathToMain string

var _ = Describe("Main", func() {

	BeforeSuite(func() {
		var err error
		pathToMain, err = gexec.Build("onboarding_broker")
		fmt.Println(pathToMain)
		Expect(err).ShouldNot(HaveOccurred())
	})

	BeforeEach(func() {
		command := exec.Command(pathToMain)
		_, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ShouldNot(HaveOccurred())
		Eventually(func() error { _, err := http.Get("http://localhost:8080"); return err }).ShouldNot(HaveOccurred())
	})

	It("should have broker routes", func() {
		response, err := http.Get("http://localhost:8080/v2/catalog")
		Expect(err).ToNot(HaveOccurred())
		Expect(response.StatusCode).To(Equal(200))
	})

	AfterEach(func() {
		gexec.Terminate()
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

})
