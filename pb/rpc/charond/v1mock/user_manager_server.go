// Code generated by mockery v1.0.0. DO NOT EDIT.

package charondmock

import (
	context "context"

	charond "github.com/piotrkowalczuk/charon/pb/rpc/charond/v1"

	mock "github.com/stretchr/testify/mock"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// UserManagerServer is an autogenerated mock type for the UserManagerServer type
type UserManagerServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) Create(_a0 context.Context, _a1 *charond.CreateUserRequest) (*charond.CreateUserResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charond.CreateUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.CreateUserRequest) *charond.CreateUserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.CreateUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.CreateUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) Delete(_a0 context.Context, _a1 *charond.DeleteUserRequest) (*wrappers.BoolValue, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *wrappers.BoolValue
	if rf, ok := ret.Get(0).(func(context.Context, *charond.DeleteUserRequest) *wrappers.BoolValue); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wrappers.BoolValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.DeleteUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) Get(_a0 context.Context, _a1 *charond.GetUserRequest) (*charond.GetUserResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charond.GetUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.GetUserRequest) *charond.GetUserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.GetUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.GetUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) List(_a0 context.Context, _a1 *charond.ListUsersRequest) (*charond.ListUsersResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charond.ListUsersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.ListUsersRequest) *charond.ListUsersResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.ListUsersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.ListUsersRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroups provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) ListGroups(_a0 context.Context, _a1 *charond.ListUserGroupsRequest) (*charond.ListUserGroupsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charond.ListUserGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.ListUserGroupsRequest) *charond.ListUserGroupsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.ListUserGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.ListUserGroupsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPermissions provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) ListPermissions(_a0 context.Context, _a1 *charond.ListUserPermissionsRequest) (*charond.ListUserPermissionsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charond.ListUserPermissionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.ListUserPermissionsRequest) *charond.ListUserPermissionsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.ListUserPermissionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.ListUserPermissionsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modify provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) Modify(_a0 context.Context, _a1 *charond.ModifyUserRequest) (*charond.ModifyUserResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charond.ModifyUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.ModifyUserRequest) *charond.ModifyUserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.ModifyUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.ModifyUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetGroups provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) SetGroups(_a0 context.Context, _a1 *charond.SetUserGroupsRequest) (*charond.SetUserGroupsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charond.SetUserGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.SetUserGroupsRequest) *charond.SetUserGroupsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.SetUserGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.SetUserGroupsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetPermissions provides a mock function with given fields: _a0, _a1
func (_m *UserManagerServer) SetPermissions(_a0 context.Context, _a1 *charond.SetUserPermissionsRequest) (*charond.SetUserPermissionsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *charond.SetUserPermissionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.SetUserPermissionsRequest) *charond.SetUserPermissionsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.SetUserPermissionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.SetUserPermissionsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
