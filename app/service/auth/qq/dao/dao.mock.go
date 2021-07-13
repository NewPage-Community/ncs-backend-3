// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/auth/qq/dao/dao.go

// Package dao is a generated GoMock package.
package dao

import (
	model "backend/app/service/auth/qq/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDao is a mock of Dao interface
type MockDao struct {
	ctrl     *gomock.Controller
	recorder *MockDaoMockRecorder
}

// MockDaoMockRecorder is the mock recorder for MockDao
type MockDaoMockRecorder struct {
	mock *MockDao
}

// NewMockDao creates a new mock instance
func NewMockDao(ctrl *gomock.Controller) *MockDao {
	mock := &MockDao{ctrl: ctrl}
	mock.recorder = &MockDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDao) EXPECT() *MockDaoMockRecorder {
	return m.recorder
}

// GetUID mocks base method
func (m *MockDao) GetUID(openID string) (*model.QQConnect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUID", openID)
	ret0, _ := ret[0].(*model.QQConnect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUID indicates an expected call of GetUID
func (mr *MockDaoMockRecorder) GetUID(openID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUID", reflect.TypeOf((*MockDao)(nil).GetUID), openID)
}

// GetStatus mocks base method
func (m *MockDao) GetStatus(uid int64) (*model.QQConnect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", uid)
	ret0, _ := ret[0].(*model.QQConnect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockDaoMockRecorder) GetStatus(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockDao)(nil).GetStatus), uid)
}

// BindQQ mocks base method
func (m *MockDao) BindQQ(info model.QQConnect) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindQQ", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindQQ indicates an expected call of BindQQ
func (mr *MockDaoMockRecorder) BindQQ(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindQQ", reflect.TypeOf((*MockDao)(nil).BindQQ), info)
}

// UnbindQQ mocks base method
func (m *MockDao) UnbindQQ(info model.QQConnect) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbindQQ", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnbindQQ indicates an expected call of UnbindQQ
func (mr *MockDaoMockRecorder) UnbindQQ(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbindQQ", reflect.TypeOf((*MockDao)(nil).UnbindQQ), info)
}

// Healthy mocks base method
func (m *MockDao) Healthy() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Healthy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Healthy indicates an expected call of Healthy
func (mr *MockDaoMockRecorder) Healthy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthy", reflect.TypeOf((*MockDao)(nil).Healthy))
}

// Close mocks base method
func (m *MockDao) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockDaoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDao)(nil).Close))
}
