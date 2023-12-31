// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	reflect "reflect"

	models "github.com/Shemistan/Lesson_5/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIStorage is a mock of IStorage interface.
type MockIStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageMockRecorder
}

// MockIStorageMockRecorder is the mock recorder for MockIStorage.
type MockIStorageMockRecorder struct {
	mock *MockIStorage
}

// NewMockIStorage creates a new mock instance.
func NewMockIStorage(ctrl *gomock.Controller) *MockIStorage {
	mock := &MockIStorage{ctrl: ctrl}
	mock.recorder = &MockIStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStorage) EXPECT() *MockIStorageMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIStorage) Add(user *models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockIStorageMockRecorder) Add(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIStorage)(nil).Add), user)
}

// Delete mocks base method.
func (m *MockIStorage) Delete(userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIStorageMockRecorder) Delete(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIStorage)(nil).Delete), userID)
}

// Get mocks base method.
func (m *MockIStorage) Get(userId int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userId)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIStorageMockRecorder) Get(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIStorage)(nil).Get), userId)
}

// Update mocks base method.
func (m *MockIStorage) Update(userId int, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIStorageMockRecorder) Update(userId, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIStorage)(nil).Update), userId, user)
}
