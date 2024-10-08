// Code generated by MockGen. DO NOT EDIT.
// Source: internal/agent/device/hook/manager.go
//
// Generated by this command:
//
//	mockgen -source=internal/agent/device/hook/manager.go -destination=internal/agent/device/hook/mock_manager.go -package=hook
//

// Package hook is a generated GoMock package.
package hook

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/flightctl/flightctl/api/v1alpha1"
	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockManager)(nil).Close))
}

// Errors mocks base method.
func (m *MockManager) Errors() []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Errors")
	ret0, _ := ret[0].([]error)
	return ret0
}

// Errors indicates an expected call of Errors.
func (mr *MockManagerMockRecorder) Errors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errors", reflect.TypeOf((*MockManager)(nil).Errors))
}

// OnAfterCreate mocks base method.
func (m *MockManager) OnAfterCreate(ctx context.Context, path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnAfterCreate", ctx, path)
}

// OnAfterCreate indicates an expected call of OnAfterCreate.
func (mr *MockManagerMockRecorder) OnAfterCreate(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnAfterCreate", reflect.TypeOf((*MockManager)(nil).OnAfterCreate), ctx, path)
}

// OnAfterRemove mocks base method.
func (m *MockManager) OnAfterRemove(ctx context.Context, path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnAfterRemove", ctx, path)
}

// OnAfterRemove indicates an expected call of OnAfterRemove.
func (mr *MockManagerMockRecorder) OnAfterRemove(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnAfterRemove", reflect.TypeOf((*MockManager)(nil).OnAfterRemove), ctx, path)
}

// OnAfterUpdate mocks base method.
func (m *MockManager) OnAfterUpdate(ctx context.Context, path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnAfterUpdate", ctx, path)
}

// OnAfterUpdate indicates an expected call of OnAfterUpdate.
func (mr *MockManagerMockRecorder) OnAfterUpdate(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnAfterUpdate", reflect.TypeOf((*MockManager)(nil).OnAfterUpdate), ctx, path)
}

// OnBeforeCreate mocks base method.
func (m *MockManager) OnBeforeCreate(ctx context.Context, path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnBeforeCreate", ctx, path)
}

// OnBeforeCreate indicates an expected call of OnBeforeCreate.
func (mr *MockManagerMockRecorder) OnBeforeCreate(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnBeforeCreate", reflect.TypeOf((*MockManager)(nil).OnBeforeCreate), ctx, path)
}

// OnBeforeRemove mocks base method.
func (m *MockManager) OnBeforeRemove(ctx context.Context, path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnBeforeRemove", ctx, path)
}

// OnBeforeRemove indicates an expected call of OnBeforeRemove.
func (mr *MockManagerMockRecorder) OnBeforeRemove(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnBeforeRemove", reflect.TypeOf((*MockManager)(nil).OnBeforeRemove), ctx, path)
}

// OnBeforeUpdate mocks base method.
func (m *MockManager) OnBeforeUpdate(ctx context.Context, path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnBeforeUpdate", ctx, path)
}

// OnBeforeUpdate indicates an expected call of OnBeforeUpdate.
func (mr *MockManagerMockRecorder) OnBeforeUpdate(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnBeforeUpdate", reflect.TypeOf((*MockManager)(nil).OnBeforeUpdate), ctx, path)
}

// Run mocks base method.
func (m *MockManager) Run(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", ctx)
}

// Run indicates an expected call of Run.
func (mr *MockManagerMockRecorder) Run(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockManager)(nil).Run), ctx)
}

// Sync mocks base method.
func (m *MockManager) Sync(current, desired *v1alpha1.RenderedDeviceSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", current, desired)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockManagerMockRecorder) Sync(current, desired any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockManager)(nil).Sync), current, desired)
}

// MockActionHook is a mock of ActionHook interface.
type MockActionHook struct {
	ctrl     *gomock.Controller
	recorder *MockActionHookMockRecorder
}

// MockActionHookMockRecorder is the mock recorder for MockActionHook.
type MockActionHookMockRecorder struct {
	mock *MockActionHook
}

// NewMockActionHook creates a new mock instance.
func NewMockActionHook(ctrl *gomock.Controller) *MockActionHook {
	mock := &MockActionHook{ctrl: ctrl}
	mock.recorder = &MockActionHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActionHook) EXPECT() *MockActionHookMockRecorder {
	return m.recorder
}

// OnChange mocks base method.
func (m *MockActionHook) OnChange(ctx context.Context, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnChange", ctx, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnChange indicates an expected call of OnChange.
func (mr *MockActionHookMockRecorder) OnChange(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChange", reflect.TypeOf((*MockActionHook)(nil).OnChange), ctx, path)
}
