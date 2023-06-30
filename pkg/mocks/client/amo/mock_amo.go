// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/client/amo/amo.go

// Package mock_amo is a generated GoMock package.
package mock_amo

import (
	context "context"
	reflect "reflect"
	amo "week3_docker/internal/client/amo"

	gomock "github.com/golang/mock/gomock"
)

// MockIAmoClient is a mock of IAmoClient interface.
type MockIAmoClient struct {
	ctrl     *gomock.Controller
	recorder *MockIAmoClientMockRecorder
}

// MockIAmoClientMockRecorder is the mock recorder for MockIAmoClient.
type MockIAmoClientMockRecorder struct {
	mock *MockIAmoClient
}

// NewMockIAmoClient creates a new mock instance.
func NewMockIAmoClient(ctrl *gomock.Controller) *MockIAmoClient {
	mock := &MockIAmoClient{ctrl: ctrl}
	mock.recorder = &MockIAmoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAmoClient) EXPECT() *MockIAmoClientMockRecorder {
	return m.recorder
}

// AccessToken mocks base method.
func (m *MockIAmoClient) AccessToken(ctx context.Context, subdomain string, req amo.AuthRequest) (amo.AuthTokenPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessToken", ctx, subdomain, req)
	ret0, _ := ret[0].(amo.AuthTokenPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessToken indicates an expected call of AccessToken.
func (mr *MockIAmoClientMockRecorder) AccessToken(ctx, subdomain, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessToken", reflect.TypeOf((*MockIAmoClient)(nil).AccessToken), ctx, subdomain, req)
}

// Account mocks base method.
func (m *MockIAmoClient) Account(ctx context.Context, req amo.AccountRequest) (amo.AccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Account", ctx, req)
	ret0, _ := ret[0].(amo.AccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Account indicates an expected call of Account.
func (mr *MockIAmoClientMockRecorder) Account(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Account", reflect.TypeOf((*MockIAmoClient)(nil).Account), ctx, req)
}

// ListContacts mocks base method.
func (m *MockIAmoClient) ListContacts(ctx context.Context, request amo.AccountRequest, params amo.ContactQueryParams) (amo.ContactResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContacts", ctx, request, params)
	ret0, _ := ret[0].(amo.ContactResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContacts indicates an expected call of ListContacts.
func (mr *MockIAmoClientMockRecorder) ListContacts(ctx, request, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContacts", reflect.TypeOf((*MockIAmoClient)(nil).ListContacts), ctx, request, params)
}

// WebHookContactsSubscribe mocks base method.
func (m *MockIAmoClient) WebHookContactsSubscribe(ctx context.Context, request amo.AccountRequest, accountID uint64) (amo.WebhookSubscribeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WebHookContactsSubscribe", ctx, request, accountID)
	ret0, _ := ret[0].(amo.WebhookSubscribeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WebHookContactsSubscribe indicates an expected call of WebHookContactsSubscribe.
func (mr *MockIAmoClientMockRecorder) WebHookContactsSubscribe(ctx, request, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebHookContactsSubscribe", reflect.TypeOf((*MockIAmoClient)(nil).WebHookContactsSubscribe), ctx, request, accountID)
}
