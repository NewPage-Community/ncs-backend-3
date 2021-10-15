// Code generated by MockGen. DO NOT EDIT.
// Source: app/game/stats/dao/dao.go

// Package dao is a generated GoMock package.
package dao

import (
	model "backend/app/game/stats/model"
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

// Get mocks base method.
func (m *MockDao) Get(stats *model.Stats) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", stats)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDaoMockRecorder) Get(stats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDao)(nil).Get), stats)
}

// GetAll mocks base method.
func (m *MockDao) GetAll(stats *model.Stats) ([]*model.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", stats)
	ret0, _ := ret[0].([]*model.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockDaoMockRecorder) GetAll(stats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDao)(nil).GetAll), stats)
}

// GetPartly mocks base method.
func (m *MockDao) GetPartly(stats *model.Stats, start, end int64) ([]*model.Stats, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartly", stats, start, end)
	ret0, _ := ret[0].([]*model.Stats)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPartly indicates an expected call of GetPartly.
func (mr *MockDaoMockRecorder) GetPartly(stats, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartly", reflect.TypeOf((*MockDao)(nil).GetPartly), stats, start, end)
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

// Incr mocks base method.
func (m *MockDao) Incr(stats *model.Stats) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", stats)
	ret0, _ := ret[0].(error)
	return ret0
}

// Incr indicates an expected call of Incr.
func (mr *MockDaoMockRecorder) Incr(stats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockDao)(nil).Incr), stats)
}

// Set mocks base method.
func (m *MockDao) Set(stats *model.Stats) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", stats)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockDaoMockRecorder) Set(stats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockDao)(nil).Set), stats)
}
