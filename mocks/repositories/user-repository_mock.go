// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/user-repository.go

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// GetLimitTaskPerDay mocks base method.
func (m *MockIUserRepository) GetLimitTaskPerDay(userId uint64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLimitTaskPerDay", userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLimitTaskPerDay indicates an expected call of GetLimitTaskPerDay.
func (mr *MockIUserRepositoryMockRecorder) GetLimitTaskPerDay(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLimitTaskPerDay", reflect.TypeOf((*MockIUserRepository)(nil).GetLimitTaskPerDay), userId)
}
