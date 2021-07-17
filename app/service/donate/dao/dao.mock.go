// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/donate/dao/dao.go

// Package dao is a generated GoMock package.
package dao

import (
	model "backend/app/service/donate/model"
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

// CreateDonate mocks base method
func (m *MockDao) CreateDonate(uid, steamID int64, amount int32, payment model.DonatePayment) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDonate", uid, steamID, amount, payment)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDonate indicates an expected call of CreateDonate
func (mr *MockDaoMockRecorder) CreateDonate(uid, steamID, amount, payment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDonate", reflect.TypeOf((*MockDao)(nil).CreateDonate), uid, steamID, amount, payment)
}

// GetTradeInfo mocks base method
func (m *MockDao) GetTradeInfo(outTradeNo string) (*model.Donate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTradeInfo", outTradeNo)
	ret0, _ := ret[0].(*model.Donate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTradeInfo indicates an expected call of GetTradeInfo
func (mr *MockDaoMockRecorder) GetTradeInfo(outTradeNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTradeInfo", reflect.TypeOf((*MockDao)(nil).GetTradeInfo), outTradeNo)
}

// FinishTrade mocks base method
func (m *MockDao) FinishTrade(outTradeNo string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishTrade", outTradeNo)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishTrade indicates an expected call of FinishTrade
func (mr *MockDaoMockRecorder) FinishTrade(outTradeNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishTrade", reflect.TypeOf((*MockDao)(nil).FinishTrade), outTradeNo)
}

// CancelTrade mocks base method
func (m *MockDao) CancelTrade(outTradeNo string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelTrade", outTradeNo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelTrade indicates an expected call of CancelTrade
func (mr *MockDaoMockRecorder) CancelTrade(outTradeNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelTrade", reflect.TypeOf((*MockDao)(nil).CancelTrade), outTradeNo)
}

// GetDonateList mocks base method
func (m *MockDao) GetDonateList(startTime, endTime int64) ([]*model.Donate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDonateList", startTime, endTime)
	ret0, _ := ret[0].([]*model.Donate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDonateList indicates an expected call of GetDonateList
func (mr *MockDaoMockRecorder) GetDonateList(startTime, endTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDonateList", reflect.TypeOf((*MockDao)(nil).GetDonateList), startTime, endTime)
}

// GetCheckTradeList mocks base method
func (m *MockDao) GetCheckTradeList() ([]*model.Donate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckTradeList")
	ret0, _ := ret[0].([]*model.Donate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckTradeList indicates an expected call of GetCheckTradeList
func (mr *MockDaoMockRecorder) GetCheckTradeList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckTradeList", reflect.TypeOf((*MockDao)(nil).GetCheckTradeList))
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
