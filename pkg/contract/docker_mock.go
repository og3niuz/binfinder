// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/contract/dockerclient.go

// Package docker is a generated GoMock package.
package contract

import (
	context "context"
	types "github.com/docker/docker/api/types"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockDockerContract is a mock of DockerContract interface
type MockDockerContract struct {
	ctrl     *gomock.Controller
	recorder *MockDockerContractMockRecorder
}

// MockDockerContractMockRecorder is the mock recorder for MockDockerContract
type MockDockerContractMockRecorder struct {
	mock *MockDockerContract
}

// NewMockDockerContract creates a new mock instance
func NewMockDockerContract(ctrl *gomock.Controller) *MockDockerContract {
	mock := &MockDockerContract{ctrl: ctrl}
	mock.recorder = &MockDockerContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDockerContract) EXPECT() *MockDockerContractMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockDockerContract) Info(ctx context.Context) (types.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", ctx)
	ret0, _ := ret[0].(types.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockDockerContractMockRecorder) Info(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockDockerContract)(nil).Info), ctx)
}

// ImagePull mocks base method
func (m *MockDockerContract) ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagePull", ctx, ref, options)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImagePull indicates an expected call of ImagePull
func (mr *MockDockerContractMockRecorder) ImagePull(ctx, ref, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagePull", reflect.TypeOf((*MockDockerContract)(nil).ImagePull), ctx, ref, options)
}