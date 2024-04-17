package exchange_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestExchange(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Exchange Suite")
}
