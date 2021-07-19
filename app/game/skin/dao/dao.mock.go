// Code generated by MockGen. DO NOT EDIT.
// Source: app/game/skin/dao/dao.go

// Package dao is a generated GoMock package.
package dao

import (
	model "backend/app/game/skin/model"
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

// GetInfo mocks base method
func (m *MockDao) GetInfo(uid int64) (*model.SkinUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", uid)
	ret0, _ := ret[0].(*model.SkinUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo
func (mr *MockDaoMockRecorder) GetInfo(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockDao)(nil).GetInfo), uid)
}

// SetUsedSkin mocks base method
func (m *MockDao) SetUsedSkin(uid int64, usedSkin int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUsedSkin", uid, usedSkin)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUsedSkin indicates an expected call of SetUsedSkin
func (mr *MockDaoMockRecorder) SetUsedSkin(uid, usedSkin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUsedSkin", reflect.TypeOf((*MockDao)(nil).SetUsedSkin), uid, usedSkin)
}

// Create mocks base method
func (m *MockDao) Create(uid int64) (*model.SkinUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", uid)
	ret0, _ := ret[0].(*model.SkinUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockDaoMockRecorder) Create(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDao)(nil).Create), uid)
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