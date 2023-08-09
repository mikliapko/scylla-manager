// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scylladb/scylla-manager/v3/pkg/restapi (interfaces: RepairService)

// Package restapi is a generated GoMock package.
package restapi

import (
	context "context"
	json "encoding/json"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repair "github.com/scylladb/scylla-manager/v3/pkg/service/repair"
	uuid "github.com/scylladb/scylla-manager/v3/pkg/util/uuid"
)

// MockRepairService is a mock of RepairService interface.
type MockRepairService struct {
	ctrl     *gomock.Controller
	recorder *MockRepairServiceMockRecorder
}

// MockRepairServiceMockRecorder is the mock recorder for MockRepairService.
type MockRepairServiceMockRecorder struct {
	mock *MockRepairService
}

// NewMockRepairService creates a new mock instance.
func NewMockRepairService(ctrl *gomock.Controller) *MockRepairService {
	mock := &MockRepairService{ctrl: ctrl}
	mock.recorder = &MockRepairServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepairService) EXPECT() *MockRepairServiceMockRecorder {
	return m.recorder
}

// GetProgress mocks base method.
func (m *MockRepairService) GetProgress(arg0 context.Context, arg1, arg2, arg3 uuid.UUID) (repair.Progress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProgress", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(repair.Progress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProgress indicates an expected call of GetProgress.
func (mr *MockRepairServiceMockRecorder) GetProgress(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProgress", reflect.TypeOf((*MockRepairService)(nil).GetProgress), arg0, arg1, arg2, arg3)
}

// GetRun mocks base method.
func (m *MockRepairService) GetRun(arg0 context.Context, arg1, arg2, arg3 uuid.UUID) (*repair.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRun", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*repair.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRun indicates an expected call of GetRun.
func (mr *MockRepairServiceMockRecorder) GetRun(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRun", reflect.TypeOf((*MockRepairService)(nil).GetRun), arg0, arg1, arg2, arg3)
}

// GetTarget mocks base method.
func (m *MockRepairService) GetTarget(arg0 context.Context, arg1 uuid.UUID, arg2 json.RawMessage) (repair.Target, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(repair.Target)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTarget indicates an expected call of GetTarget.
func (mr *MockRepairServiceMockRecorder) GetTarget(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTarget", reflect.TypeOf((*MockRepairService)(nil).GetTarget), arg0, arg1, arg2)
}

// SetIntensity mocks base method.
func (m *MockRepairService) SetIntensity(arg0 context.Context, arg1 uuid.UUID, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIntensity", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetIntensity indicates an expected call of SetIntensity.
func (mr *MockRepairServiceMockRecorder) SetIntensity(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIntensity", reflect.TypeOf((*MockRepairService)(nil).SetIntensity), arg0, arg1, arg2)
}

// SetParallel mocks base method.
func (m *MockRepairService) SetParallel(arg0 context.Context, arg1 uuid.UUID, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetParallel", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetParallel indicates an expected call of SetParallel.
func (mr *MockRepairServiceMockRecorder) SetParallel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParallel", reflect.TypeOf((*MockRepairService)(nil).SetParallel), arg0, arg1, arg2)
}
