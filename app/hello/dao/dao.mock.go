// Code generated by MockGen. DO NOT EDIT.
// Source: app/hello/dao/dao.go

// Package dao is a generated GoMock package.
package dao

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDao is a mock of Dao interface.
type MockDao struct {
	ctrl     *gomock.Controller
	recorder *MockDaoMockRecorder
}

// MockDaoMockRecorder is the mock recorder for MockDao.
type MockDaoMockRecorder struct {
	mock *MockDao
}

// NewMockDao creates a new mock instance.
func NewMockDao(ctrl *gomock.Controller) *MockDao {
	mock := &MockDao{ctrl: ctrl}
	mock.recorder = &MockDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDao) EXPECT() *MockDaoMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDao) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockDaoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDao)(nil).Close))
}

// Healthy mocks base method.
func (m *MockDao) Healthy() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Healthy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Healthy indicates an expected call of Healthy.
func (mr *MockDaoMockRecorder) Healthy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthy", reflect.TypeOf((*MockDao)(nil).Healthy))
}

// Hello mocks base method.
func (m *MockDao) Hello() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Hello")
}

// Hello indicates an expected call of Hello.
func (mr *MockDaoMockRecorder) Hello() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockDao)(nil).Hello))
}
