// Code generated by MockGen. DO NOT EDIT.
// Source: app/game/server/api/grpc/server.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockServerClient is a mock of ServerClient interface
type MockServerClient struct {
	ctrl     *gomock.Controller
	recorder *MockServerClientMockRecorder
}

// MockServerClientMockRecorder is the mock recorder for MockServerClient
type MockServerClientMockRecorder struct {
	mock *MockServerClient
}

// NewMockServerClient creates a new mock instance
func NewMockServerClient(ctrl *gomock.Controller) *MockServerClient {
	mock := &MockServerClient{ctrl: ctrl}
	mock.recorder = &MockServerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerClient) EXPECT() *MockServerClientMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockServerClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Info", varargs...)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockServerClientMockRecorder) Info(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockServerClient)(nil).Info), varargs...)
}

// MockServerServer is a mock of ServerServer interface
type MockServerServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerServerMockRecorder
}

// MockServerServerMockRecorder is the mock recorder for MockServerServer
type MockServerServerMockRecorder struct {
	mock *MockServerServer
}

// NewMockServerServer creates a new mock instance
func NewMockServerServer(ctrl *gomock.Controller) *MockServerServer {
	mock := &MockServerServer{ctrl: ctrl}
	mock.recorder = &MockServerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerServer) EXPECT() *MockServerServerMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockServerServer) Info(arg0 context.Context, arg1 *InfoReq) (*InfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockServerServerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockServerServer)(nil).Info), arg0, arg1)
}
