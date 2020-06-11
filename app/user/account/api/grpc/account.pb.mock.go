// Code generated by MockGen. DO NOT EDIT.
// Source: app/user/account/api/grpc/account.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAccountClient is a mock of AccountClient interface
type MockAccountClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountClientMockRecorder
}

// MockAccountClientMockRecorder is the mock recorder for MockAccountClient
type MockAccountClientMockRecorder struct {
	mock *MockAccountClient
}

// NewMockAccountClient creates a new mock instance
func NewMockAccountClient(ctrl *gomock.Controller) *MockAccountClient {
	mock := &MockAccountClient{ctrl: ctrl}
	mock.recorder = &MockAccountClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountClient) EXPECT() *MockAccountClientMockRecorder {
	return m.recorder
}

// UID mocks base method
func (m *MockAccountClient) UID(ctx context.Context, in *UIDReq, opts ...grpc.CallOption) (*UIDResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UID", varargs...)
	ret0, _ := ret[0].(*UIDResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UID indicates an expected call of UID
func (mr *MockAccountClientMockRecorder) UID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UID", reflect.TypeOf((*MockAccountClient)(nil).UID), varargs...)
}

// Info mocks base method
func (m *MockAccountClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
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
func (mr *MockAccountClientMockRecorder) Info(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockAccountClient)(nil).Info), varargs...)
}

// Register mocks base method
func (m *MockAccountClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*RegisterResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockAccountClientMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAccountClient)(nil).Register), varargs...)
}

// ChangeName mocks base method
func (m *MockAccountClient) ChangeName(ctx context.Context, in *ChangeNameReq, opts ...grpc.CallOption) (*ChangeNameResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeName", varargs...)
	ret0, _ := ret[0].(*ChangeNameResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeName indicates an expected call of ChangeName
func (mr *MockAccountClientMockRecorder) ChangeName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeName", reflect.TypeOf((*MockAccountClient)(nil).ChangeName), varargs...)
}

// PlayerConnect mocks base method
func (m *MockAccountClient) PlayerConnect(ctx context.Context, in *PlayerConnectReq, opts ...grpc.CallOption) (*PlayerConnectResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PlayerConnect", varargs...)
	ret0, _ := ret[0].(*PlayerConnectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlayerConnect indicates an expected call of PlayerConnect
func (mr *MockAccountClientMockRecorder) PlayerConnect(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlayerConnect", reflect.TypeOf((*MockAccountClient)(nil).PlayerConnect), varargs...)
}

// MockAccountServer is a mock of AccountServer interface
type MockAccountServer struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServerMockRecorder
}

// MockAccountServerMockRecorder is the mock recorder for MockAccountServer
type MockAccountServerMockRecorder struct {
	mock *MockAccountServer
}

// NewMockAccountServer creates a new mock instance
func NewMockAccountServer(ctrl *gomock.Controller) *MockAccountServer {
	mock := &MockAccountServer{ctrl: ctrl}
	mock.recorder = &MockAccountServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountServer) EXPECT() *MockAccountServerMockRecorder {
	return m.recorder
}

// UID mocks base method
func (m *MockAccountServer) UID(arg0 context.Context, arg1 *UIDReq) (*UIDResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UID", arg0, arg1)
	ret0, _ := ret[0].(*UIDResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UID indicates an expected call of UID
func (mr *MockAccountServerMockRecorder) UID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UID", reflect.TypeOf((*MockAccountServer)(nil).UID), arg0, arg1)
}

// Info mocks base method
func (m *MockAccountServer) Info(arg0 context.Context, arg1 *InfoReq) (*InfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockAccountServerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockAccountServer)(nil).Info), arg0, arg1)
}

// Register mocks base method
func (m *MockAccountServer) Register(arg0 context.Context, arg1 *RegisterReq) (*RegisterResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*RegisterResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockAccountServerMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAccountServer)(nil).Register), arg0, arg1)
}

// ChangeName mocks base method
func (m *MockAccountServer) ChangeName(arg0 context.Context, arg1 *ChangeNameReq) (*ChangeNameResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeName", arg0, arg1)
	ret0, _ := ret[0].(*ChangeNameResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeName indicates an expected call of ChangeName
func (mr *MockAccountServerMockRecorder) ChangeName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeName", reflect.TypeOf((*MockAccountServer)(nil).ChangeName), arg0, arg1)
}

// PlayerConnect mocks base method
func (m *MockAccountServer) PlayerConnect(arg0 context.Context, arg1 *PlayerConnectReq) (*PlayerConnectResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlayerConnect", arg0, arg1)
	ret0, _ := ret[0].(*PlayerConnectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlayerConnect indicates an expected call of PlayerConnect
func (mr *MockAccountServerMockRecorder) PlayerConnect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlayerConnect", reflect.TypeOf((*MockAccountServer)(nil).PlayerConnect), arg0, arg1)
}
