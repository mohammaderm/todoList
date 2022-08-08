// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/job/job.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/mohammaderm/todoList/internal/models"
)

// MockJobRepository is a mock of JobRepository interface.
type MockJobRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJobRepositoryMockRecorder
}

// MockJobRepositoryMockRecorder is the mock recorder for MockJobRepository.
type MockJobRepositoryMockRecorder struct {
	mock *MockJobRepository
}

// NewMockJobRepository creates a new mock instance.
func NewMockJobRepository(ctrl *gomock.Controller) *MockJobRepository {
	mock := &MockJobRepository{ctrl: ctrl}
	mock.recorder = &MockJobRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobRepository) EXPECT() *MockJobRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockJobRepository) Create(ctx context.Context, job models.CreateJob) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockJobRepositoryMockRecorder) Create(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockJobRepository)(nil).Create), ctx, job)
}

// Delete mocks base method.
func (m *MockJobRepository) Delete(ctx context.Context, jobid, accountid uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, jobid, accountid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockJobRepositoryMockRecorder) Delete(ctx, jobid, accountid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockJobRepository)(nil).Delete), ctx, jobid, accountid)
}

// GetAll mocks base method.
func (m *MockJobRepository) GetAll(ctx context.Context, accountid uint, offset int) (*[]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, accountid, offset)
	ret0, _ := ret[0].(*[]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockJobRepositoryMockRecorder) GetAll(ctx, accountid, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockJobRepository)(nil).GetAll), ctx, accountid, offset)
}

// Update mocks base method.
func (m *MockJobRepository) Update(ctx context.Context, jobid, accountid uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, jobid, accountid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockJobRepositoryMockRecorder) Update(ctx, jobid, accountid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockJobRepository)(nil).Update), ctx, jobid, accountid)
}