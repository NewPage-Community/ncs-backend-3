// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/auth/qq/api/grpc/v1/qq.pb.go

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

// SignInWithQQ mocks base method
func (m *MockQQClient) SignInWithQQ(ctx context.Context, in *SignInWithQQReq, opts ...grpc.CallOption) (*SignInWithQQResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignInWithQQ", varargs...)
	ret0, _ := ret[0].(*SignInWithQQResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInWithQQ indicates an expected call of SignInWithQQ
func (mr *MockQQClientMockRecorder) SignInWithQQ(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInWithQQ", reflect.TypeOf((*MockQQClient)(nil).SignInWithQQ), varargs...)
}

// GetQQConnectStatus mocks base method
func (m *MockQQClient) GetQQConnectStatus(ctx context.Context, in *GetQQConnectStatusReq, opts ...grpc.CallOption) (*GetQQConnectStatusResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQQConnectStatus", varargs...)
	ret0, _ := ret[0].(*GetQQConnectStatusResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQQConnectStatus indicates an expected call of GetQQConnectStatus
func (mr *MockQQClientMockRecorder) GetQQConnectStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQQConnectStatus", reflect.TypeOf((*MockQQClient)(nil).GetQQConnectStatus), varargs...)
}

// GetUID mocks base method
func (m *MockQQClient) GetUID(ctx context.Context, in *GetUIDReq, opts ...grpc.CallOption) (*GetUIDResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUID", varargs...)
	ret0, _ := ret[0].(*GetUIDResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUID indicates an expected call of GetUID
func (mr *MockQQClientMockRecorder) GetUID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUID", reflect.TypeOf((*MockQQClient)(nil).GetUID), varargs...)
}

// BindQQ mocks base method
func (m *MockQQClient) BindQQ(ctx context.Context, in *BindQQReq, opts ...grpc.CallOption) (*BindQQResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BindQQ", varargs...)
	ret0, _ := ret[0].(*BindQQResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindQQ indicates an expected call of BindQQ
func (mr *MockQQClientMockRecorder) BindQQ(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindQQ", reflect.TypeOf((*MockQQClient)(nil).BindQQ), varargs...)
}

// UnbindQQ mocks base method
func (m *MockQQClient) UnbindQQ(ctx context.Context, in *UnbindQQReq, opts ...grpc.CallOption) (*UnbindQQResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnbindQQ", varargs...)
	ret0, _ := ret[0].(*UnbindQQResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnbindQQ indicates an expected call of UnbindQQ
func (mr *MockQQClientMockRecorder) UnbindQQ(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbindQQ", reflect.TypeOf((*MockQQClient)(nil).UnbindQQ), varargs...)
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

// SignInWithQQ mocks base method
func (m *MockQQServer) SignInWithQQ(arg0 context.Context, arg1 *SignInWithQQReq) (*SignInWithQQResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignInWithQQ", arg0, arg1)
	ret0, _ := ret[0].(*SignInWithQQResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInWithQQ indicates an expected call of SignInWithQQ
func (mr *MockQQServerMockRecorder) SignInWithQQ(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInWithQQ", reflect.TypeOf((*MockQQServer)(nil).SignInWithQQ), arg0, arg1)
}

// GetQQConnectStatus mocks base method
func (m *MockQQServer) GetQQConnectStatus(arg0 context.Context, arg1 *GetQQConnectStatusReq) (*GetQQConnectStatusResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQQConnectStatus", arg0, arg1)
	ret0, _ := ret[0].(*GetQQConnectStatusResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQQConnectStatus indicates an expected call of GetQQConnectStatus
func (mr *MockQQServerMockRecorder) GetQQConnectStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQQConnectStatus", reflect.TypeOf((*MockQQServer)(nil).GetQQConnectStatus), arg0, arg1)
}

// GetUID mocks base method
func (m *MockQQServer) GetUID(arg0 context.Context, arg1 *GetUIDReq) (*GetUIDResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUID", arg0, arg1)
	ret0, _ := ret[0].(*GetUIDResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUID indicates an expected call of GetUID
func (mr *MockQQServerMockRecorder) GetUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUID", reflect.TypeOf((*MockQQServer)(nil).GetUID), arg0, arg1)
}

// BindQQ mocks base method
func (m *MockQQServer) BindQQ(arg0 context.Context, arg1 *BindQQReq) (*BindQQResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindQQ", arg0, arg1)
	ret0, _ := ret[0].(*BindQQResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindQQ indicates an expected call of BindQQ
func (mr *MockQQServerMockRecorder) BindQQ(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindQQ", reflect.TypeOf((*MockQQServer)(nil).BindQQ), arg0, arg1)
}

// UnbindQQ mocks base method
func (m *MockQQServer) UnbindQQ(arg0 context.Context, arg1 *UnbindQQReq) (*UnbindQQResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbindQQ", arg0, arg1)
	ret0, _ := ret[0].(*UnbindQQResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnbindQQ indicates an expected call of UnbindQQ
func (mr *MockQQServerMockRecorder) UnbindQQ(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbindQQ", reflect.TypeOf((*MockQQServer)(nil).UnbindQQ), arg0, arg1)
}
