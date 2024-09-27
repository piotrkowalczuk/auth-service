// Code generated by mockery v1.0.0. DO NOT EDIT.

package modelmock

import (
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// Rows is an autogenerated mock type for the Rows type
type Rows struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Rows) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ColumnTypes provides a mock function with given fields:
func (_m *Rows) ColumnTypes() ([]*sql.ColumnType, error) {
	ret := _m.Called()

	var r0 []*sql.ColumnType
	if rf, ok := ret.Get(0).(func() []*sql.ColumnType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*sql.ColumnType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Columns provides a mock function with given fields:
func (_m *Rows) Columns() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Err provides a mock function with given fields:
func (_m *Rows) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *Rows) Next() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NextResultSet provides a mock function with given fields:
func (_m *Rows) NextResultSet() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Scan provides a mock function with given fields: dst
func (_m *Rows) Scan(dst ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dst...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dst...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
