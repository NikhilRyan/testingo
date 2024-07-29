package runner

import (
	"testing"
)

type TestSuite struct {
	Setup    func()
	Teardown func()
	Tests    []func(t *testing.T)
}

func RunTestSuite(t *testing.T, suite TestSuite) {
	if suite.Setup != nil {
		suite.Setup()
	}

	for _, test := range suite.Tests {
		test(t)
	}

	if suite.Teardown != nil {
		suite.Teardown()
	}
}
