// Code generated by MockGen. DO NOT EDIT.
// Source: resourcegroups.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	resourcegroups "github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	gomock "github.com/golang/mock/gomock"
)

// MockResourcegroupsClient is a mock of ResourcegroupsClient interface.
type MockResourcegroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourcegroupsClientMockRecorder
}

// MockResourcegroupsClientMockRecorder is the mock recorder for MockResourcegroupsClient.
type MockResourcegroupsClientMockRecorder struct {
	mock *MockResourcegroupsClient
}

// NewMockResourcegroupsClient creates a new mock instance.
func NewMockResourcegroupsClient(ctrl *gomock.Controller) *MockResourcegroupsClient {
	mock := &MockResourcegroupsClient{ctrl: ctrl}
	mock.recorder = &MockResourcegroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourcegroupsClient) EXPECT() *MockResourcegroupsClientMockRecorder {
	return m.recorder
}

// GetGroup mocks base method.
func (m *MockResourcegroupsClient) GetGroup(arg0 context.Context, arg1 *resourcegroups.GetGroupInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroup", varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockResourcegroupsClientMockRecorder) GetGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockResourcegroupsClient)(nil).GetGroup), varargs...)
}

// GetGroupConfiguration mocks base method.
func (m *MockResourcegroupsClient) GetGroupConfiguration(arg0 context.Context, arg1 *resourcegroups.GetGroupConfigurationInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupConfiguration", varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetGroupConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupConfiguration indicates an expected call of GetGroupConfiguration.
func (mr *MockResourcegroupsClientMockRecorder) GetGroupConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupConfiguration", reflect.TypeOf((*MockResourcegroupsClient)(nil).GetGroupConfiguration), varargs...)
}

// GetGroupQuery mocks base method.
func (m *MockResourcegroupsClient) GetGroupQuery(arg0 context.Context, arg1 *resourcegroups.GetGroupQueryInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupQuery", varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetGroupQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupQuery indicates an expected call of GetGroupQuery.
func (mr *MockResourcegroupsClientMockRecorder) GetGroupQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupQuery", reflect.TypeOf((*MockResourcegroupsClient)(nil).GetGroupQuery), varargs...)
}

// GetTags mocks base method.
func (m *MockResourcegroupsClient) GetTags(arg0 context.Context, arg1 *resourcegroups.GetTagsInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTags", varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags.
func (mr *MockResourcegroupsClientMockRecorder) GetTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockResourcegroupsClient)(nil).GetTags), varargs...)
}

// ListGroupResources mocks base method.
func (m *MockResourcegroupsClient) ListGroupResources(arg0 context.Context, arg1 *resourcegroups.ListGroupResourcesInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupResources", varargs...)
	ret0, _ := ret[0].(*resourcegroups.ListGroupResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupResources indicates an expected call of ListGroupResources.
func (mr *MockResourcegroupsClientMockRecorder) ListGroupResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupResources", reflect.TypeOf((*MockResourcegroupsClient)(nil).ListGroupResources), varargs...)
}

// ListGroups mocks base method.
func (m *MockResourcegroupsClient) ListGroups(arg0 context.Context, arg1 *resourcegroups.ListGroupsInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroups", varargs...)
	ret0, _ := ret[0].(*resourcegroups.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockResourcegroupsClientMockRecorder) ListGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockResourcegroupsClient)(nil).ListGroups), varargs...)
}

// SearchResources mocks base method.
func (m *MockResourcegroupsClient) SearchResources(arg0 context.Context, arg1 *resourcegroups.SearchResourcesInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.SearchResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchResources", varargs...)
	ret0, _ := ret[0].(*resourcegroups.SearchResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchResources indicates an expected call of SearchResources.
func (mr *MockResourcegroupsClientMockRecorder) SearchResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchResources", reflect.TypeOf((*MockResourcegroupsClient)(nil).SearchResources), varargs...)
}
