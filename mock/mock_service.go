// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\Lenovo\Desktop\todolist_backend\service\service.go

// Package mock is a generated GoMock package.
package mock

import (
	model "bootcamp/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIService) Get() []model.ToDoModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].([]model.ToDoModel)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockIServiceMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIService)(nil).Get))
}

// Save mocks base method.
func (m *MockIService) Save(todo string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", todo)
	ret0, _ := ret[0].(int)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIServiceMockRecorder) Save(todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIService)(nil).Save), todo)
}
