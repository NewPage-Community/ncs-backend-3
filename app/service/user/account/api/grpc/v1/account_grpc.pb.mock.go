// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/user/account/api/grpc/v1/account_grpc.pb.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAccountClient is a mock of AccountClient interface.
type MockAccountClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountClientMockRecorder
}

// MockAccountClientMockRecorder is the mock recorder for MockAccountClient.
type MockAccountClientMockRecorder struct {
	mock *MockAccountClient
}

// NewMockAccountClient creates a new mock instance.
func NewMockAccountClient(ctrl *gomock.Controller) *MockAccountClient {
	mock := &MockAccountClient{ctrl: ctrl}
	mock.recorder = &MockAccountClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountClient) EXPECT() *MockAccountClientMockRecorder {
	return m.recorder
}

// ChangeName mocks base method.
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

// ChangeName indicates an expected call of ChangeName.
func (mr *MockAccountClientMockRecorder) ChangeName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeName", reflect.TypeOf((*MockAccountClient)(nil).ChangeName), varargs...)
}

// GetAllUID mocks base method.
func (m *MockAccountClient) GetAllUID(ctx context.Context, in *GetAllUIDReq, opts ...grpc.CallOption) (*GetAllUIDResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllUID", varargs...)
	ret0, _ := ret[0].(*GetAllUIDResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUID indicates an expected call of GetAllUID.
func (mr *MockAccountClientMockRecorder) GetAllUID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUID", reflect.TypeOf((*MockAccountClient)(nil).GetAllUID), varargs...)
}

// Info mocks base method.
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

// Info indicates an expected call of Info.
func (mr *MockAccountClientMockRecorder) Info(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockAccountClient)(nil).Info), varargs...)
}

// Register mocks base method.
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

// Register indicates an expected call of Register.
func (mr *MockAccountClientMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAccountClient)(nil).Register), varargs...)
}

// UID mocks base method.
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

// UID indicates an expected call of UID.
func (mr *MockAccountClientMockRecorder) UID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UID", reflect.TypeOf((*MockAccountClient)(nil).UID), varargs...)
}

// MockAccountServer is a mock of AccountServer interface.
type MockAccountServer struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServerMockRecorder
}

// MockAccountServerMockRecorder is the mock recorder for MockAccountServer.
type MockAccountServerMockRecorder struct {
	mock *MockAccountServer
}

// NewMockAccountServer creates a new mock instance.
func NewMockAccountServer(ctrl *gomock.Controller) *MockAccountServer {
	mock := &MockAccountServer{ctrl: ctrl}
	mock.recorder = &MockAccountServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountServer) EXPECT() *MockAccountServerMockRecorder {
	return m.recorder
}

// ChangeName mocks base method.
func (m *MockAccountServer) ChangeName(arg0 context.Context, arg1 *ChangeNameReq) (*ChangeNameResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeName", arg0, arg1)
	ret0, _ := ret[0].(*ChangeNameResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeName indicates an expected call of ChangeName.
func (mr *MockAccountServerMockRecorder) ChangeName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeName", reflect.TypeOf((*MockAccountServer)(nil).ChangeName), arg0, arg1)
}

// GetAllUID mocks base method.
func (m *MockAccountServer) GetAllUID(arg0 context.Context, arg1 *GetAllUIDReq) (*GetAllUIDResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUID", arg0, arg1)
	ret0, _ := ret[0].(*GetAllUIDResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUID indicates an expected call of GetAllUID.
func (mr *MockAccountServerMockRecorder) GetAllUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUID", reflect.TypeOf((*MockAccountServer)(nil).GetAllUID), arg0, arg1)
}

// Info mocks base method.
func (m *MockAccountServer) Info(arg0 context.Context, arg1 *InfoReq) (*InfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockAccountServerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockAccountServer)(nil).Info), arg0, arg1)
}

// Register mocks base method.
func (m *MockAccountServer) Register(arg0 context.Context, arg1 *RegisterReq) (*RegisterResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*RegisterResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAccountServerMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAccountServer)(nil).Register), arg0, arg1)
}

// UID mocks base method.
func (m *MockAccountServer) UID(arg0 context.Context, arg1 *UIDReq) (*UIDResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UID", arg0, arg1)
	ret0, _ := ret[0].(*UIDResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UID indicates an expected call of UID.
func (mr *MockAccountServerMockRecorder) UID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UID", reflect.TypeOf((*MockAccountServer)(nil).UID), arg0, arg1)
}

// mustEmbedUnimplementedAccountServer mocks base method.
func (m *MockAccountServer) mustEmbedUnimplementedAccountServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAccountServer")
}

// mustEmbedUnimplementedAccountServer indicates an expected call of mustEmbedUnimplementedAccountServer.
func (mr *MockAccountServerMockRecorder) mustEmbedUnimplementedAccountServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAccountServer", reflect.TypeOf((*MockAccountServer)(nil).mustEmbedUnimplementedAccountServer))
}

// MockUnsafeAccountServer is a mock of UnsafeAccountServer interface.
type MockUnsafeAccountServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAccountServerMockRecorder
}

// MockUnsafeAccountServerMockRecorder is the mock recorder for MockUnsafeAccountServer.
type MockUnsafeAccountServerMockRecorder struct {
	mock *MockUnsafeAccountServer
}

// NewMockUnsafeAccountServer creates a new mock instance.
func NewMockUnsafeAccountServer(ctrl *gomock.Controller) *MockUnsafeAccountServer {
	mock := &MockUnsafeAccountServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAccountServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAccountServer) EXPECT() *MockUnsafeAccountServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAccountServer mocks base method.
func (m *MockUnsafeAccountServer) mustEmbedUnimplementedAccountServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAccountServer")
}

// mustEmbedUnimplementedAccountServer indicates an expected call of mustEmbedUnimplementedAccountServer.
func (mr *MockUnsafeAccountServerMockRecorder) mustEmbedUnimplementedAccountServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAccountServer", reflect.TypeOf((*MockUnsafeAccountServer)(nil).mustEmbedUnimplementedAccountServer))
}

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

// Info mocks base method.
func (m *MockWebClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
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

// Info indicates an expected call of Info.
func (mr *MockWebClientMockRecorder) Info(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockWebClient)(nil).Info), varargs...)
}

// UID mocks base method.
func (m *MockWebClient) UID(ctx context.Context, in *UIDReq, opts ...grpc.CallOption) (*UIDResp, error) {
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

// UID indicates an expected call of UID.
func (mr *MockWebClientMockRecorder) UID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UID", reflect.TypeOf((*MockWebClient)(nil).UID), varargs...)
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

// Info mocks base method.
func (m *MockWebServer) Info(arg0 context.Context, arg1 *InfoReq) (*InfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockWebServerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockWebServer)(nil).Info), arg0, arg1)
}

// UID mocks base method.
func (m *MockWebServer) UID(arg0 context.Context, arg1 *UIDReq) (*UIDResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UID", arg0, arg1)
	ret0, _ := ret[0].(*UIDResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UID indicates an expected call of UID.
func (mr *MockWebServerMockRecorder) UID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UID", reflect.TypeOf((*MockWebServer)(nil).UID), arg0, arg1)
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