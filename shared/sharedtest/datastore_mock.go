// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/shared (interfaces: Datastore)

// Package sharedtest is a generated GoMock package.
package sharedtest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
	reflect "reflect"
)

// MockDatastore is a mock of Datastore interface
type MockDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockDatastoreMockRecorder
}

// MockDatastoreMockRecorder is the mock recorder for MockDatastore
type MockDatastoreMockRecorder struct {
	mock *MockDatastore
}

// NewMockDatastore creates a new mock instance
func NewMockDatastore(ctrl *gomock.Controller) *MockDatastore {
	mock := &MockDatastore{ctrl: ctrl}
	mock.recorder = &MockDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatastore) EXPECT() *MockDatastoreMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockDatastore) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockDatastoreMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDatastore)(nil).Context))
}

// Delete mocks base method
func (m *MockDatastore) Delete(arg0 shared.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDatastoreMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDatastore)(nil).Delete), arg0)
}

// Done mocks base method
func (m *MockDatastore) Done() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Done indicates an expected call of Done
func (mr *MockDatastoreMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockDatastore)(nil).Done))
}

// Get mocks base method
func (m *MockDatastore) Get(arg0 shared.Key, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockDatastoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDatastore)(nil).Get), arg0, arg1)
}

// GetAll mocks base method
func (m *MockDatastore) GetAll(arg0 shared.Query, arg1 interface{}) ([]shared.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].([]shared.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockDatastoreMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDatastore)(nil).GetAll), arg0, arg1)
}

// GetMulti mocks base method
func (m *MockDatastore) GetMulti(arg0 []shared.Key, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMulti", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetMulti indicates an expected call of GetMulti
func (mr *MockDatastoreMockRecorder) GetMulti(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMulti", reflect.TypeOf((*MockDatastore)(nil).GetMulti), arg0, arg1)
}

// Insert mocks base method
func (m *MockDatastore) Insert(arg0 shared.Key, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockDatastoreMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDatastore)(nil).Insert), arg0, arg1)
}

// NewIDKey mocks base method
func (m *MockDatastore) NewIDKey(arg0 string, arg1 int64) shared.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewIDKey", arg0, arg1)
	ret0, _ := ret[0].(shared.Key)
	return ret0
}

// NewIDKey indicates an expected call of NewIDKey
func (mr *MockDatastoreMockRecorder) NewIDKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewIDKey", reflect.TypeOf((*MockDatastore)(nil).NewIDKey), arg0, arg1)
}

// NewNameKey mocks base method
func (m *MockDatastore) NewNameKey(arg0, arg1 string) shared.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNameKey", arg0, arg1)
	ret0, _ := ret[0].(shared.Key)
	return ret0
}

// NewNameKey indicates an expected call of NewNameKey
func (mr *MockDatastoreMockRecorder) NewNameKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNameKey", reflect.TypeOf((*MockDatastore)(nil).NewNameKey), arg0, arg1)
}

// NewQuery mocks base method
func (m *MockDatastore) NewQuery(arg0 string) shared.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewQuery", arg0)
	ret0, _ := ret[0].(shared.Query)
	return ret0
}

// NewQuery indicates an expected call of NewQuery
func (mr *MockDatastoreMockRecorder) NewQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewQuery", reflect.TypeOf((*MockDatastore)(nil).NewQuery), arg0)
}

// Put mocks base method
func (m *MockDatastore) Put(arg0 shared.Key, arg1 interface{}) (shared.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(shared.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put
func (mr *MockDatastoreMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockDatastore)(nil).Put), arg0, arg1)
}

// ReserveID mocks base method
func (m *MockDatastore) ReserveID(arg0 string) (shared.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveID", arg0)
	ret0, _ := ret[0].(shared.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReserveID indicates an expected call of ReserveID
func (mr *MockDatastoreMockRecorder) ReserveID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveID", reflect.TypeOf((*MockDatastore)(nil).ReserveID), arg0)
}

// TestRunQuery mocks base method
func (m *MockDatastore) TestRunQuery() shared.TestRunQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestRunQuery")
	ret0, _ := ret[0].(shared.TestRunQuery)
	return ret0
}

// TestRunQuery indicates an expected call of TestRunQuery
func (mr *MockDatastoreMockRecorder) TestRunQuery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestRunQuery", reflect.TypeOf((*MockDatastore)(nil).TestRunQuery))
}
