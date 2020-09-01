// Code generated by MockGen. DO NOT EDIT.
// Source: app/game/cvar/api/grpc/cvar.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockCVarClient is a mock of CVarClient interface
type MockCVarClient struct {
	ctrl     *gomock.Controller
	recorder *MockCVarClientMockRecorder
}

// MockCVarClientMockRecorder is the mock recorder for MockCVarClient
type MockCVarClientMockRecorder struct {
	mock *MockCVarClient
}

// NewMockCVarClient creates a new mock instance
func NewMockCVarClient(ctrl *gomock.Controller) *MockCVarClient {
	mock := &MockCVarClient{ctrl: ctrl}
	mock.recorder = &MockCVarClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCVarClient) EXPECT() *MockCVarClientMockRecorder {
	return m.recorder
}

// UpdateCVars mocks base method
func (m *MockCVarClient) UpdateCVars(ctx context.Context, in *UpdateCVarsReq, opts ...grpc.CallOption) (*UpdateCVarsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCVars", varargs...)
	ret0, _ := ret[0].(*UpdateCVarsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCVars indicates an expected call of UpdateCVars
func (mr *MockCVarClientMockRecorder) UpdateCVars(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCVars", reflect.TypeOf((*MockCVarClient)(nil).UpdateCVars), varargs...)
}

// MockCVarServer is a mock of CVarServer interface
type MockCVarServer struct {
	ctrl     *gomock.Controller
	recorder *MockCVarServerMockRecorder
}

// MockCVarServerMockRecorder is the mock recorder for MockCVarServer
type MockCVarServerMockRecorder struct {
	mock *MockCVarServer
}

// NewMockCVarServer creates a new mock instance
func NewMockCVarServer(ctrl *gomock.Controller) *MockCVarServer {
	mock := &MockCVarServer{ctrl: ctrl}
	mock.recorder = &MockCVarServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCVarServer) EXPECT() *MockCVarServerMockRecorder {
	return m.recorder
}

// UpdateCVars mocks base method
func (m *MockCVarServer) UpdateCVars(arg0 context.Context, arg1 *UpdateCVarsReq) (*UpdateCVarsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCVars", arg0, arg1)
	ret0, _ := ret[0].(*UpdateCVarsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCVars indicates an expected call of UpdateCVars
func (mr *MockCVarServerMockRecorder) UpdateCVars(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCVars", reflect.TypeOf((*MockCVarServer)(nil).UpdateCVars), arg0, arg1)
}