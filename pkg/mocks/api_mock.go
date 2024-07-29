package mocks

import (
	"github.com/jarcoal/httpmock"
)

type APIMock struct{}

func NewAPIMock() *APIMock {
	httpmock.Activate()
	return &APIMock{}
}

func (m *APIMock) RegisterResponder(method, url string, responder httpmock.Responder) {
	httpmock.RegisterResponder(method, url, responder)
}

func (m *APIMock) DeactivateAndReset() {
	httpmock.DeactivateAndReset()
}
