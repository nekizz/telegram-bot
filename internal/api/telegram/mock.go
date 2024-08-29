// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nekizz/telegram-bot/internal/api/telegram (interfaces: Service)
//
// Generated by this command:
//
//	mockgen -destination internal/api/telegram/mock.go -package=telegram github.com/nekizz/telegram-bot/internal/api/telegram Service
//

// Package telegram is a generated GoMock package.
package telegram

import (
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// alertJob mocks base method.
func (m *MockService) alertJob(arg0 echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "alertJob", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// alertJob indicates an expected call of alertJob.
func (mr *MockServiceMockRecorder) alertJob(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "alertJob", reflect.TypeOf((*MockService)(nil).alertJob), arg0)
}

// checkBirthdays mocks base method.
func (m *MockService) checkBirthdays(arg0 echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "checkBirthdays", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// checkBirthdays indicates an expected call of checkBirthdays.
func (mr *MockServiceMockRecorder) checkBirthdays(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "checkBirthdays", reflect.TypeOf((*MockService)(nil).checkBirthdays), arg0)
}

// handleCommand mocks base method.
func (m *MockService) handleCommand() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleCommand")
	ret0, _ := ret[0].(error)
	return ret0
}

// handleCommand indicates an expected call of handleCommand.
func (mr *MockServiceMockRecorder) handleCommand() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleCommand", reflect.TypeOf((*MockService)(nil).handleCommand))
}
