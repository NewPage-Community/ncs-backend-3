// Code generated by MockGen. DO NOT EDIT.
// Source: app/game/skin/api/grpc/v1/skin_grpc.pb.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockSkinClient is a mock of SkinClient interface.
type MockSkinClient struct {
	ctrl     *gomock.Controller
	recorder *MockSkinClientMockRecorder
}

// MockSkinClientMockRecorder is the mock recorder for MockSkinClient.
type MockSkinClientMockRecorder struct {
	mock *MockSkinClient
}

// NewMockSkinClient creates a new mock instance.
func NewMockSkinClient(ctrl *gomock.Controller) *MockSkinClient {
	mock := &MockSkinClient{ctrl: ctrl}
	mock.recorder = &MockSkinClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSkinClient) EXPECT() *MockSkinClientMockRecorder {
	return m.recorder
}

// GetInfo mocks base method.
func (m *MockSkinClient) GetInfo(ctx context.Context, in *GetInfoReq, opts ...grpc.CallOption) (*GetInfoResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInfo", varargs...)
	ret0, _ := ret[0].(*GetInfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockSkinClientMockRecorder) GetInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockSkinClient)(nil).GetInfo), varargs...)
}

// GetSkins mocks base method.
func (m *MockSkinClient) GetSkins(ctx context.Context, in *GetSkinsReq, opts ...grpc.CallOption) (*GetSkinsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSkins", varargs...)
	ret0, _ := ret[0].(*GetSkinsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSkins indicates an expected call of GetSkins.
func (mr *MockSkinClientMockRecorder) GetSkins(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSkins", reflect.TypeOf((*MockSkinClient)(nil).GetSkins), varargs...)
}

// SetUsedSkin mocks base method.
func (m *MockSkinClient) SetUsedSkin(ctx context.Context, in *SetUsedSkinReq, opts ...grpc.CallOption) (*SetUsedSkinResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetUsedSkin", varargs...)
	ret0, _ := ret[0].(*SetUsedSkinResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUsedSkin indicates an expected call of SetUsedSkin.
func (mr *MockSkinClientMockRecorder) SetUsedSkin(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUsedSkin", reflect.TypeOf((*MockSkinClient)(nil).SetUsedSkin), varargs...)
}

// MockSkinServer is a mock of SkinServer interface.
type MockSkinServer struct {
	ctrl     *gomock.Controller
	recorder *MockSkinServerMockRecorder
}

// MockSkinServerMockRecorder is the mock recorder for MockSkinServer.
type MockSkinServerMockRecorder struct {
	mock *MockSkinServer
}

// NewMockSkinServer creates a new mock instance.
func NewMockSkinServer(ctrl *gomock.Controller) *MockSkinServer {
	mock := &MockSkinServer{ctrl: ctrl}
	mock.recorder = &MockSkinServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSkinServer) EXPECT() *MockSkinServerMockRecorder {
	return m.recorder
}

// GetInfo mocks base method.
func (m *MockSkinServer) GetInfo(arg0 context.Context, arg1 *GetInfoReq) (*GetInfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", arg0, arg1)
	ret0, _ := ret[0].(*GetInfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockSkinServerMockRecorder) GetInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockSkinServer)(nil).GetInfo), arg0, arg1)
}

// GetSkins mocks base method.
func (m *MockSkinServer) GetSkins(arg0 context.Context, arg1 *GetSkinsReq) (*GetSkinsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSkins", arg0, arg1)
	ret0, _ := ret[0].(*GetSkinsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSkins indicates an expected call of GetSkins.
func (mr *MockSkinServerMockRecorder) GetSkins(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSkins", reflect.TypeOf((*MockSkinServer)(nil).GetSkins), arg0, arg1)
}

// SetUsedSkin mocks base method.
func (m *MockSkinServer) SetUsedSkin(arg0 context.Context, arg1 *SetUsedSkinReq) (*SetUsedSkinResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUsedSkin", arg0, arg1)
	ret0, _ := ret[0].(*SetUsedSkinResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUsedSkin indicates an expected call of SetUsedSkin.
func (mr *MockSkinServerMockRecorder) SetUsedSkin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUsedSkin", reflect.TypeOf((*MockSkinServer)(nil).SetUsedSkin), arg0, arg1)
}

// mustEmbedUnimplementedSkinServer mocks base method.
func (m *MockSkinServer) mustEmbedUnimplementedSkinServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSkinServer")
}

// mustEmbedUnimplementedSkinServer indicates an expected call of mustEmbedUnimplementedSkinServer.
func (mr *MockSkinServerMockRecorder) mustEmbedUnimplementedSkinServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSkinServer", reflect.TypeOf((*MockSkinServer)(nil).mustEmbedUnimplementedSkinServer))
}

// MockUnsafeSkinServer is a mock of UnsafeSkinServer interface.
type MockUnsafeSkinServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSkinServerMockRecorder
}

// MockUnsafeSkinServerMockRecorder is the mock recorder for MockUnsafeSkinServer.
type MockUnsafeSkinServerMockRecorder struct {
	mock *MockUnsafeSkinServer
}

// NewMockUnsafeSkinServer creates a new mock instance.
func NewMockUnsafeSkinServer(ctrl *gomock.Controller) *MockUnsafeSkinServer {
	mock := &MockUnsafeSkinServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSkinServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSkinServer) EXPECT() *MockUnsafeSkinServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSkinServer mocks base method.
func (m *MockUnsafeSkinServer) mustEmbedUnimplementedSkinServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSkinServer")
}

// mustEmbedUnimplementedSkinServer indicates an expected call of mustEmbedUnimplementedSkinServer.
func (mr *MockUnsafeSkinServerMockRecorder) mustEmbedUnimplementedSkinServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSkinServer", reflect.TypeOf((*MockUnsafeSkinServer)(nil).mustEmbedUnimplementedSkinServer))
}
