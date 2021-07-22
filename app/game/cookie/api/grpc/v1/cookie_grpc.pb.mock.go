// Code generated by MockGen. DO NOT EDIT.
// Source: app/game/cookie/api/grpc/v1/cookie_grpc.pb.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockCookieClient is a mock of CookieClient interface.
type MockCookieClient struct {
	ctrl     *gomock.Controller
	recorder *MockCookieClientMockRecorder
}

// MockCookieClientMockRecorder is the mock recorder for MockCookieClient.
type MockCookieClientMockRecorder struct {
	mock *MockCookieClient
}

// NewMockCookieClient creates a new mock instance.
func NewMockCookieClient(ctrl *gomock.Controller) *MockCookieClient {
	mock := &MockCookieClient{ctrl: ctrl}
	mock.recorder = &MockCookieClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCookieClient) EXPECT() *MockCookieClientMockRecorder {
	return m.recorder
}

// GetAllCookie mocks base method.
func (m *MockCookieClient) GetAllCookie(ctx context.Context, in *GetAllCookieReq, opts ...grpc.CallOption) (*GetAllCookieResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllCookie", varargs...)
	ret0, _ := ret[0].(*GetAllCookieResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCookie indicates an expected call of GetAllCookie.
func (mr *MockCookieClientMockRecorder) GetAllCookie(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCookie", reflect.TypeOf((*MockCookieClient)(nil).GetAllCookie), varargs...)
}

// GetCookie mocks base method.
func (m *MockCookieClient) GetCookie(ctx context.Context, in *GetCookieReq, opts ...grpc.CallOption) (*GetCookieResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCookie", varargs...)
	ret0, _ := ret[0].(*GetCookieResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCookie indicates an expected call of GetCookie.
func (mr *MockCookieClientMockRecorder) GetCookie(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCookie", reflect.TypeOf((*MockCookieClient)(nil).GetCookie), varargs...)
}

// SetCookie mocks base method.
func (m *MockCookieClient) SetCookie(ctx context.Context, in *SetCookieReq, opts ...grpc.CallOption) (*SetCookieResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetCookie", varargs...)
	ret0, _ := ret[0].(*SetCookieResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCookie indicates an expected call of SetCookie.
func (mr *MockCookieClientMockRecorder) SetCookie(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCookie", reflect.TypeOf((*MockCookieClient)(nil).SetCookie), varargs...)
}

// MockCookieServer is a mock of CookieServer interface.
type MockCookieServer struct {
	ctrl     *gomock.Controller
	recorder *MockCookieServerMockRecorder
}

// MockCookieServerMockRecorder is the mock recorder for MockCookieServer.
type MockCookieServerMockRecorder struct {
	mock *MockCookieServer
}

// NewMockCookieServer creates a new mock instance.
func NewMockCookieServer(ctrl *gomock.Controller) *MockCookieServer {
	mock := &MockCookieServer{ctrl: ctrl}
	mock.recorder = &MockCookieServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCookieServer) EXPECT() *MockCookieServerMockRecorder {
	return m.recorder
}

// GetAllCookie mocks base method.
func (m *MockCookieServer) GetAllCookie(arg0 context.Context, arg1 *GetAllCookieReq) (*GetAllCookieResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCookie", arg0, arg1)
	ret0, _ := ret[0].(*GetAllCookieResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCookie indicates an expected call of GetAllCookie.
func (mr *MockCookieServerMockRecorder) GetAllCookie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCookie", reflect.TypeOf((*MockCookieServer)(nil).GetAllCookie), arg0, arg1)
}

// GetCookie mocks base method.
func (m *MockCookieServer) GetCookie(arg0 context.Context, arg1 *GetCookieReq) (*GetCookieResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCookie", arg0, arg1)
	ret0, _ := ret[0].(*GetCookieResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCookie indicates an expected call of GetCookie.
func (mr *MockCookieServerMockRecorder) GetCookie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCookie", reflect.TypeOf((*MockCookieServer)(nil).GetCookie), arg0, arg1)
}

// SetCookie mocks base method.
func (m *MockCookieServer) SetCookie(arg0 context.Context, arg1 *SetCookieReq) (*SetCookieResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCookie", arg0, arg1)
	ret0, _ := ret[0].(*SetCookieResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCookie indicates an expected call of SetCookie.
func (mr *MockCookieServerMockRecorder) SetCookie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCookie", reflect.TypeOf((*MockCookieServer)(nil).SetCookie), arg0, arg1)
}

// mustEmbedUnimplementedCookieServer mocks base method.
func (m *MockCookieServer) mustEmbedUnimplementedCookieServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedCookieServer")
}

// mustEmbedUnimplementedCookieServer indicates an expected call of mustEmbedUnimplementedCookieServer.
func (mr *MockCookieServerMockRecorder) mustEmbedUnimplementedCookieServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedCookieServer", reflect.TypeOf((*MockCookieServer)(nil).mustEmbedUnimplementedCookieServer))
}

// MockUnsafeCookieServer is a mock of UnsafeCookieServer interface.
type MockUnsafeCookieServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeCookieServerMockRecorder
}

// MockUnsafeCookieServerMockRecorder is the mock recorder for MockUnsafeCookieServer.
type MockUnsafeCookieServerMockRecorder struct {
	mock *MockUnsafeCookieServer
}

// NewMockUnsafeCookieServer creates a new mock instance.
func NewMockUnsafeCookieServer(ctrl *gomock.Controller) *MockUnsafeCookieServer {
	mock := &MockUnsafeCookieServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeCookieServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeCookieServer) EXPECT() *MockUnsafeCookieServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedCookieServer mocks base method.
func (m *MockUnsafeCookieServer) mustEmbedUnimplementedCookieServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedCookieServer")
}

// mustEmbedUnimplementedCookieServer indicates an expected call of mustEmbedUnimplementedCookieServer.
func (mr *MockUnsafeCookieServerMockRecorder) mustEmbedUnimplementedCookieServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedCookieServer", reflect.TypeOf((*MockUnsafeCookieServer)(nil).mustEmbedUnimplementedCookieServer))
}