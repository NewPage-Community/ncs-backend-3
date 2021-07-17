// Code generated by MockGen. DO NOT EDIT.
// Source: app/game/a2s/api/grpc/v1/a2s.pb.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockA2SClient is a mock of A2SClient interface
type MockA2SClient struct {
	ctrl     *gomock.Controller
	recorder *MockA2SClientMockRecorder
}

// MockA2SClientMockRecorder is the mock recorder for MockA2SClient
type MockA2SClientMockRecorder struct {
	mock *MockA2SClient
}

// NewMockA2SClient creates a new mock instance
func NewMockA2SClient(ctrl *gomock.Controller) *MockA2SClient {
	mock := &MockA2SClient{ctrl: ctrl}
	mock.recorder = &MockA2SClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockA2SClient) EXPECT() *MockA2SClientMockRecorder {
	return m.recorder
}

// A2SQuery mocks base method
func (m *MockA2SClient) A2SQuery(ctx context.Context, in *A2SQueryReq, opts ...grpc.CallOption) (*A2SQueryResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "A2SQuery", varargs...)
	ret0, _ := ret[0].(*A2SQueryResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// A2SQuery indicates an expected call of A2SQuery
func (mr *MockA2SClientMockRecorder) A2SQuery(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "A2SQuery", reflect.TypeOf((*MockA2SClient)(nil).A2SQuery), varargs...)
}

// MockA2SServer is a mock of A2SServer interface
type MockA2SServer struct {
	ctrl     *gomock.Controller
	recorder *MockA2SServerMockRecorder
}

// MockA2SServerMockRecorder is the mock recorder for MockA2SServer
type MockA2SServerMockRecorder struct {
	mock *MockA2SServer
}

// NewMockA2SServer creates a new mock instance
func NewMockA2SServer(ctrl *gomock.Controller) *MockA2SServer {
	mock := &MockA2SServer{ctrl: ctrl}
	mock.recorder = &MockA2SServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockA2SServer) EXPECT() *MockA2SServerMockRecorder {
	return m.recorder
}

// A2SQuery mocks base method
func (m *MockA2SServer) A2SQuery(arg0 context.Context, arg1 *A2SQueryReq) (*A2SQueryResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "A2SQuery", arg0, arg1)
	ret0, _ := ret[0].(*A2SQueryResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// A2SQuery indicates an expected call of A2SQuery
func (mr *MockA2SServerMockRecorder) A2SQuery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "A2SQuery", reflect.TypeOf((*MockA2SServer)(nil).A2SQuery), arg0, arg1)
}
