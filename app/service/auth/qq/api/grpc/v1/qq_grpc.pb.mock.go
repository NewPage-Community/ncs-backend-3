// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/auth/qq/api/grpc/v1/qq_grpc.pb.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockWebClient is a mock of WebClient interface.
type MockWebClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebClientMockRecorder
}

// MockWebClientMockRecorder is the mock recorder for MockWebClient.
type MockWebClientMockRecorder struct {
	mock *MockWebClient
}

// NewMockWebClient creates a new mock instance.
func NewMockWebClient(ctrl *gomock.Controller) *MockWebClient {
	mock := &MockWebClient{ctrl: ctrl}
	mock.recorder = &MockWebClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebClient) EXPECT() *MockWebClientMockRecorder {
	return m.recorder
}

// AuthQQ mocks base method.
func (m *MockWebClient) AuthQQ(ctx context.Context, in *AuthQQReq, opts ...grpc.CallOption) (*AuthQQResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthQQ", varargs...)
	ret0, _ := ret[0].(*AuthQQResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthQQ indicates an expected call of AuthQQ.
func (mr *MockWebClientMockRecorder) AuthQQ(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthQQ", reflect.TypeOf((*MockWebClient)(nil).AuthQQ), varargs...)
}

// BindQQ mocks base method.
func (m *MockWebClient) BindQQ(ctx context.Context, in *BindQQReq, opts ...grpc.CallOption) (*BindQQResp, error) {
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

// BindQQ indicates an expected call of BindQQ.
func (mr *MockWebClientMockRecorder) BindQQ(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindQQ", reflect.TypeOf((*MockWebClient)(nil).BindQQ), varargs...)
}

// QQStatus mocks base method.
func (m *MockWebClient) QQStatus(ctx context.Context, in *QQStatusReq, opts ...grpc.CallOption) (*QQStatusResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QQStatus", varargs...)
	ret0, _ := ret[0].(*QQStatusResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QQStatus indicates an expected call of QQStatus.
func (mr *MockWebClientMockRecorder) QQStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QQStatus", reflect.TypeOf((*MockWebClient)(nil).QQStatus), varargs...)
}

// UnbindQQ mocks base method.
func (m *MockWebClient) UnbindQQ(ctx context.Context, in *UnbindQQReq, opts ...grpc.CallOption) (*UnbindQQResp, error) {
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

// UnbindQQ indicates an expected call of UnbindQQ.
func (mr *MockWebClientMockRecorder) UnbindQQ(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbindQQ", reflect.TypeOf((*MockWebClient)(nil).UnbindQQ), varargs...)
}

// MockWebServer is a mock of WebServer interface.
type MockWebServer struct {
	ctrl     *gomock.Controller
	recorder *MockWebServerMockRecorder
}

// MockWebServerMockRecorder is the mock recorder for MockWebServer.
type MockWebServerMockRecorder struct {
	mock *MockWebServer
}

// NewMockWebServer creates a new mock instance.
func NewMockWebServer(ctrl *gomock.Controller) *MockWebServer {
	mock := &MockWebServer{ctrl: ctrl}
	mock.recorder = &MockWebServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebServer) EXPECT() *MockWebServerMockRecorder {
	return m.recorder
}

// AuthQQ mocks base method.
func (m *MockWebServer) AuthQQ(arg0 context.Context, arg1 *AuthQQReq) (*AuthQQResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthQQ", arg0, arg1)
	ret0, _ := ret[0].(*AuthQQResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthQQ indicates an expected call of AuthQQ.
func (mr *MockWebServerMockRecorder) AuthQQ(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthQQ", reflect.TypeOf((*MockWebServer)(nil).AuthQQ), arg0, arg1)
}

// BindQQ mocks base method.
func (m *MockWebServer) BindQQ(arg0 context.Context, arg1 *BindQQReq) (*BindQQResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindQQ", arg0, arg1)
	ret0, _ := ret[0].(*BindQQResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindQQ indicates an expected call of BindQQ.
func (mr *MockWebServerMockRecorder) BindQQ(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindQQ", reflect.TypeOf((*MockWebServer)(nil).BindQQ), arg0, arg1)
}

// QQStatus mocks base method.
func (m *MockWebServer) QQStatus(arg0 context.Context, arg1 *QQStatusReq) (*QQStatusResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QQStatus", arg0, arg1)
	ret0, _ := ret[0].(*QQStatusResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QQStatus indicates an expected call of QQStatus.
func (mr *MockWebServerMockRecorder) QQStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QQStatus", reflect.TypeOf((*MockWebServer)(nil).QQStatus), arg0, arg1)
}

// UnbindQQ mocks base method.
func (m *MockWebServer) UnbindQQ(arg0 context.Context, arg1 *UnbindQQReq) (*UnbindQQResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbindQQ", arg0, arg1)
	ret0, _ := ret[0].(*UnbindQQResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnbindQQ indicates an expected call of UnbindQQ.
func (mr *MockWebServerMockRecorder) UnbindQQ(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbindQQ", reflect.TypeOf((*MockWebServer)(nil).UnbindQQ), arg0, arg1)
}

// mustEmbedUnimplementedWebServer mocks base method.
func (m *MockWebServer) mustEmbedUnimplementedWebServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedWebServer")
}

// mustEmbedUnimplementedWebServer indicates an expected call of mustEmbedUnimplementedWebServer.
func (mr *MockWebServerMockRecorder) mustEmbedUnimplementedWebServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedWebServer", reflect.TypeOf((*MockWebServer)(nil).mustEmbedUnimplementedWebServer))
}

// MockUnsafeWebServer is a mock of UnsafeWebServer interface.
type MockUnsafeWebServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeWebServerMockRecorder
}

// MockUnsafeWebServerMockRecorder is the mock recorder for MockUnsafeWebServer.
type MockUnsafeWebServerMockRecorder struct {
	mock *MockUnsafeWebServer
}

// NewMockUnsafeWebServer creates a new mock instance.
func NewMockUnsafeWebServer(ctrl *gomock.Controller) *MockUnsafeWebServer {
	mock := &MockUnsafeWebServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeWebServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeWebServer) EXPECT() *MockUnsafeWebServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedWebServer mocks base method.
func (m *MockUnsafeWebServer) mustEmbedUnimplementedWebServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedWebServer")
}

// mustEmbedUnimplementedWebServer indicates an expected call of mustEmbedUnimplementedWebServer.
func (mr *MockUnsafeWebServerMockRecorder) mustEmbedUnimplementedWebServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedWebServer", reflect.TypeOf((*MockUnsafeWebServer)(nil).mustEmbedUnimplementedWebServer))
}