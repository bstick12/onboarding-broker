package broker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRedisServiceBroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisServiceBroker Suite")
}
