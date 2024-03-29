// Code generated by MockGen. DO NOT EDIT.
// Source: app/game/stats/api/grpc/v1/public_grpc.pb.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockStatsPublicClient is a mock of StatsPublicClient interface.
type MockStatsPublicClient struct {
	ctrl     *gomock.Controller
	recorder *MockStatsPublicClientMockRecorder
}

// MockStatsPublicClientMockRecorder is the mock recorder for MockStatsPublicClient.
type MockStatsPublicClientMockRecorder struct {
	mock *MockStatsPublicClient
}

// NewMockStatsPublicClient creates a new mock instance.
func NewMockStatsPublicClient(ctrl *gomock.Controller) *MockStatsPublicClient {
	mock := &MockStatsPublicClient{ctrl: ctrl}
	mock.recorder = &MockStatsPublicClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatsPublicClient) EXPECT() *MockStatsPublicClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockStatsPublicClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*GetResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStatsPublicClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStatsPublicClient)(nil).Get), varargs...)
}

// GetPartly mocks base method.
func (m *MockStatsPublicClient) GetPartly(ctx context.Context, in *GetPartlyReq, opts ...grpc.CallOption) (*GetPartlyResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPartly", varargs...)
	ret0, _ := ret[0].(*GetPartlyResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartly indicates an expected call of GetPartly.
func (mr *MockStatsPublicClientMockRecorder) GetPartly(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartly", reflect.TypeOf((*MockStatsPublicClient)(nil).GetPartly), varargs...)
}

// MockStatsPublicServer is a mock of StatsPublicServer interface.
type MockStatsPublicServer struct {
	ctrl     *gomock.Controller
	recorder *MockStatsPublicServerMockRecorder
}

// MockStatsPublicServerMockRecorder is the mock recorder for MockStatsPublicServer.
type MockStatsPublicServerMockRecorder struct {
	mock *MockStatsPublicServer
}

// NewMockStatsPublicServer creates a new mock instance.
func NewMockStatsPublicServer(ctrl *gomock.Controller) *MockStatsPublicServer {
	mock := &MockStatsPublicServer{ctrl: ctrl}
	mock.recorder = &MockStatsPublicServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatsPublicServer) EXPECT() *MockStatsPublicServerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockStatsPublicServer) Get(arg0 context.Context, arg1 *GetReq) (*GetResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*GetResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStatsPublicServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStatsPublicServer)(nil).Get), arg0, arg1)
}

// GetPartly mocks base method.
func (m *MockStatsPublicServer) GetPartly(arg0 context.Context, arg1 *GetPartlyReq) (*GetPartlyResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartly", arg0, arg1)
	ret0, _ := ret[0].(*GetPartlyResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartly indicates an expected call of GetPartly.
func (mr *MockStatsPublicServerMockRecorder) GetPartly(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartly", reflect.TypeOf((*MockStatsPublicServer)(nil).GetPartly), arg0, arg1)
}

// mustEmbedUnimplementedStatsPublicServer mocks base method.
func (m *MockStatsPublicServer) mustEmbedUnimplementedStatsPublicServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedStatsPublicServer")
}

// mustEmbedUnimplementedStatsPublicServer indicates an expected call of mustEmbedUnimplementedStatsPublicServer.
func (mr *MockStatsPublicServerMockRecorder) mustEmbedUnimplementedStatsPublicServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedStatsPublicServer", reflect.TypeOf((*MockStatsPublicServer)(nil).mustEmbedUnimplementedStatsPublicServer))
}

// MockUnsafeStatsPublicServer is a mock of UnsafeStatsPublicServer interface.
type MockUnsafeStatsPublicServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeStatsPublicServerMockRecorder
}

// MockUnsafeStatsPublicServerMockRecorder is the mock recorder for MockUnsafeStatsPublicServer.
type MockUnsafeStatsPublicServerMockRecorder struct {
	mock *MockUnsafeStatsPublicServer
}

// NewMockUnsafeStatsPublicServer creates a new mock instance.
func NewMockUnsafeStatsPublicServer(ctrl *gomock.Controller) *MockUnsafeStatsPublicServer {
	mock := &MockUnsafeStatsPublicServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeStatsPublicServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeStatsPublicServer) EXPECT() *MockUnsafeStatsPublicServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedStatsPublicServer mocks base method.
func (m *MockUnsafeStatsPublicServer) mustEmbedUnimplementedStatsPublicServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedStatsPublicServer")
}

// mustEmbedUnimplementedStatsPublicServer indicates an expected call of mustEmbedUnimplementedStatsPublicServer.
func (mr *MockUnsafeStatsPublicServerMockRecorder) mustEmbedUnimplementedStatsPublicServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedStatsPublicServer", reflect.TypeOf((*MockUnsafeStatsPublicServer)(nil).mustEmbedUnimplementedStatsPublicServer))
}
