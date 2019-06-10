// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/blackhatbrigade/gomessagestore/repository (interfaces: Repository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	message "github.com/blackhatbrigade/gomessagestore/message"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetAllMessagesInCategory mocks base method
func (m *MockRepository) GetAllMessagesInCategory(arg0 context.Context, arg1 string) ([]*message.MessageEnvelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMessagesInCategory", arg0, arg1)
	ret0, _ := ret[0].([]*message.MessageEnvelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMessagesInCategory indicates an expected call of GetAllMessagesInCategory
func (mr *MockRepositoryMockRecorder) GetAllMessagesInCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMessagesInCategory", reflect.TypeOf((*MockRepository)(nil).GetAllMessagesInCategory), arg0, arg1)
}

// GetAllMessagesInCategorySince mocks base method
func (m *MockRepository) GetAllMessagesInCategorySince(arg0 context.Context, arg1 string, arg2 int64) ([]*message.MessageEnvelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMessagesInCategorySince", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*message.MessageEnvelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMessagesInCategorySince indicates an expected call of GetAllMessagesInCategorySince
func (mr *MockRepositoryMockRecorder) GetAllMessagesInCategorySince(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMessagesInCategorySince", reflect.TypeOf((*MockRepository)(nil).GetAllMessagesInCategorySince), arg0, arg1, arg2)
}

// GetAllMessagesInStream mocks base method
func (m *MockRepository) GetAllMessagesInStream(arg0 context.Context, arg1 string) ([]*message.MessageEnvelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMessagesInStream", arg0, arg1)
	ret0, _ := ret[0].([]*message.MessageEnvelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMessagesInStream indicates an expected call of GetAllMessagesInStream
func (mr *MockRepositoryMockRecorder) GetAllMessagesInStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMessagesInStream", reflect.TypeOf((*MockRepository)(nil).GetAllMessagesInStream), arg0, arg1)
}

// GetAllMessagesInStreamSince mocks base method
func (m *MockRepository) GetAllMessagesInStreamSince(arg0 context.Context, arg1 string, arg2 int64) ([]*message.MessageEnvelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMessagesInStreamSince", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*message.MessageEnvelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMessagesInStreamSince indicates an expected call of GetAllMessagesInStreamSince
func (mr *MockRepositoryMockRecorder) GetAllMessagesInStreamSince(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMessagesInStreamSince", reflect.TypeOf((*MockRepository)(nil).GetAllMessagesInStreamSince), arg0, arg1, arg2)
}

// GetLastMessageInStream mocks base method
func (m *MockRepository) GetLastMessageInStream(arg0 context.Context, arg1 string) (*message.MessageEnvelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastMessageInStream", arg0, arg1)
	ret0, _ := ret[0].(*message.MessageEnvelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastMessageInStream indicates an expected call of GetLastMessageInStream
func (mr *MockRepositoryMockRecorder) GetLastMessageInStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastMessageInStream", reflect.TypeOf((*MockRepository)(nil).GetLastMessageInStream), arg0, arg1)
}

// WriteMessage mocks base method
func (m *MockRepository) WriteMessage(arg0 context.Context, arg1 *message.MessageEnvelope) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage
func (mr *MockRepositoryMockRecorder) WriteMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockRepository)(nil).WriteMessage), arg0, arg1)
}

// WriteMessageWithExpectedPosition mocks base method
func (m *MockRepository) WriteMessageWithExpectedPosition(arg0 context.Context, arg1 *message.MessageEnvelope, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessageWithExpectedPosition", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessageWithExpectedPosition indicates an expected call of WriteMessageWithExpectedPosition
func (mr *MockRepositoryMockRecorder) WriteMessageWithExpectedPosition(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessageWithExpectedPosition", reflect.TypeOf((*MockRepository)(nil).WriteMessageWithExpectedPosition), arg0, arg1, arg2)
}
