// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/donate/api/grpc/donate.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockDonateClient is a mock of DonateClient interface
type MockDonateClient struct {
	ctrl     *gomock.Controller
	recorder *MockDonateClientMockRecorder
}

// MockDonateClientMockRecorder is the mock recorder for MockDonateClient
type MockDonateClientMockRecorder struct {
	mock *MockDonateClient
}

// NewMockDonateClient creates a new mock instance
func NewMockDonateClient(ctrl *gomock.Controller) *MockDonateClient {
	mock := &MockDonateClient{ctrl: ctrl}
	mock.recorder = &MockDonateClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDonateClient) EXPECT() *MockDonateClientMockRecorder {
	return m.recorder
}

// CreateDonate mocks base method
func (m *MockDonateClient) CreateDonate(ctx context.Context, in *CreateDonateReq, opts ...grpc.CallOption) (*CreateDonateResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDonate", varargs...)
	ret0, _ := ret[0].(*CreateDonateResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDonate indicates an expected call of CreateDonate
func (mr *MockDonateClientMockRecorder) CreateDonate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDonate", reflect.TypeOf((*MockDonateClient)(nil).CreateDonate), varargs...)
}

// QueryDonate mocks base method
func (m *MockDonateClient) QueryDonate(ctx context.Context, in *QueryDonateReq, opts ...grpc.CallOption) (*QueryDonateResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryDonate", varargs...)
	ret0, _ := ret[0].(*QueryDonateResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryDonate indicates an expected call of QueryDonate
func (mr *MockDonateClientMockRecorder) QueryDonate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryDonate", reflect.TypeOf((*MockDonateClient)(nil).QueryDonate), varargs...)
}

// GetDonateList mocks base method
func (m *MockDonateClient) GetDonateList(ctx context.Context, in *GetDonateListReq, opts ...grpc.CallOption) (*GetDonateListResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDonateList", varargs...)
	ret0, _ := ret[0].(*GetDonateListResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDonateList indicates an expected call of GetDonateList
func (mr *MockDonateClientMockRecorder) GetDonateList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDonateList", reflect.TypeOf((*MockDonateClient)(nil).GetDonateList), varargs...)
}

// AddDonate mocks base method
func (m *MockDonateClient) AddDonate(ctx context.Context, in *AddDonateReq, opts ...grpc.CallOption) (*AddDonateResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddDonate", varargs...)
	ret0, _ := ret[0].(*AddDonateResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDonate indicates an expected call of AddDonate
func (mr *MockDonateClientMockRecorder) AddDonate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDonate", reflect.TypeOf((*MockDonateClient)(nil).AddDonate), varargs...)
}

// MockDonateServer is a mock of DonateServer interface
type MockDonateServer struct {
	ctrl     *gomock.Controller
	recorder *MockDonateServerMockRecorder
}

// MockDonateServerMockRecorder is the mock recorder for MockDonateServer
type MockDonateServerMockRecorder struct {
	mock *MockDonateServer
}

// NewMockDonateServer creates a new mock instance
func NewMockDonateServer(ctrl *gomock.Controller) *MockDonateServer {
	mock := &MockDonateServer{ctrl: ctrl}
	mock.recorder = &MockDonateServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDonateServer) EXPECT() *MockDonateServerMockRecorder {
	return m.recorder
}

// CreateDonate mocks base method
func (m *MockDonateServer) CreateDonate(arg0 context.Context, arg1 *CreateDonateReq) (*CreateDonateResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDonate", arg0, arg1)
	ret0, _ := ret[0].(*CreateDonateResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDonate indicates an expected call of CreateDonate
func (mr *MockDonateServerMockRecorder) CreateDonate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDonate", reflect.TypeOf((*MockDonateServer)(nil).CreateDonate), arg0, arg1)
}

// QueryDonate mocks base method
func (m *MockDonateServer) QueryDonate(arg0 context.Context, arg1 *QueryDonateReq) (*QueryDonateResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryDonate", arg0, arg1)
	ret0, _ := ret[0].(*QueryDonateResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryDonate indicates an expected call of QueryDonate
func (mr *MockDonateServerMockRecorder) QueryDonate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryDonate", reflect.TypeOf((*MockDonateServer)(nil).QueryDonate), arg0, arg1)
}

// GetDonateList mocks base method
func (m *MockDonateServer) GetDonateList(arg0 context.Context, arg1 *GetDonateListReq) (*GetDonateListResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDonateList", arg0, arg1)
	ret0, _ := ret[0].(*GetDonateListResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDonateList indicates an expected call of GetDonateList
func (mr *MockDonateServerMockRecorder) GetDonateList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDonateList", reflect.TypeOf((*MockDonateServer)(nil).GetDonateList), arg0, arg1)
}

// AddDonate mocks base method
func (m *MockDonateServer) AddDonate(arg0 context.Context, arg1 *AddDonateReq) (*AddDonateResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDonate", arg0, arg1)
	ret0, _ := ret[0].(*AddDonateResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDonate indicates an expected call of AddDonate
func (mr *MockDonateServerMockRecorder) AddDonate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDonate", reflect.TypeOf((*MockDonateServer)(nil).AddDonate), arg0, arg1)
}
