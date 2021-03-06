// Code generated by MockGen. DO NOT EDIT.
// Source: app/bot/qq/api/grpc/v1/qq.pb.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockQQClient is a mock of QQClient interface
type MockQQClient struct {
	ctrl     *gomock.Controller
	recorder *MockQQClientMockRecorder
}

// MockQQClientMockRecorder is the mock recorder for MockQQClient
type MockQQClientMockRecorder struct {
	mock *MockQQClient
}

// NewMockQQClient creates a new mock instance
func NewMockQQClient(ctrl *gomock.Controller) *MockQQClient {
	mock := &MockQQClient{ctrl: ctrl}
	mock.recorder = &MockQQClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQQClient) EXPECT() *MockQQClientMockRecorder {
	return m.recorder
}

// SendGroupMessage mocks base method
func (m *MockQQClient) SendGroupMessage(ctx context.Context, in *SendGroupMessageReq, opts ...grpc.CallOption) (*SendGroupMessageResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendGroupMessage", varargs...)
	ret0, _ := ret[0].(*SendGroupMessageResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendGroupMessage indicates an expected call of SendGroupMessage
func (mr *MockQQClientMockRecorder) SendGroupMessage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGroupMessage", reflect.TypeOf((*MockQQClient)(nil).SendGroupMessage), varargs...)
}

// MockQQServer is a mock of QQServer interface
type MockQQServer struct {
	ctrl     *gomock.Controller
	recorder *MockQQServerMockRecorder
}

// MockQQServerMockRecorder is the mock recorder for MockQQServer
type MockQQServerMockRecorder struct {
	mock *MockQQServer
}

// NewMockQQServer creates a new mock instance
func NewMockQQServer(ctrl *gomock.Controller) *MockQQServer {
	mock := &MockQQServer{ctrl: ctrl}
	mock.recorder = &MockQQServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQQServer) EXPECT() *MockQQServerMockRecorder {
	return m.recorder
}

// SendGroupMessage mocks base method
func (m *MockQQServer) SendGroupMessage(arg0 context.Context, arg1 *SendGroupMessageReq) (*SendGroupMessageResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendGroupMessage", arg0, arg1)
	ret0, _ := ret[0].(*SendGroupMessageResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendGroupMessage indicates an expected call of SendGroupMessage
func (mr *MockQQServerMockRecorder) SendGroupMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGroupMessage", reflect.TypeOf((*MockQQServer)(nil).SendGroupMessage), arg0, arg1)
}
