// Code generated by mockery v1.0.0
package mocks

import (
	"io"

	"github.com/stretchr/testify/mock"
)

// RequesterVariadic is an autogenerated mock type for the RequesterVariadic type
type RequesterVariadic struct {
	mock.Mock
}

// Get provides a mock function with given fields: values
func (_m *RequesterVariadic) Get(values ...string) bool {
	_va := make([]interface{}, len(values))
	for _i := range values {
		_va[_i] = values[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(...string) bool); ok {
		r0 = rf(values...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MultiWriteToFile provides a mock function with given fields: filename, w
func (_m *RequesterVariadic) MultiWriteToFile(filename string, w ...io.Writer) string {
	_va := make([]interface{}, len(w))
	for _i := range w {
		_va[_i] = w[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, filename)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...io.Writer) string); ok {
		r0 = rf(filename, w...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OneInterface provides a mock function with given fields: a
func (_m *RequesterVariadic) OneInterface(a ...interface{}) bool {
	var _ca []interface{}
	_ca = append(_ca, a...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(...interface{}) bool); ok {
		r0 = rf(a...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Sprintf provides a mock function with given fields: format, a
func (_m *RequesterVariadic) Sprintf(format string, a ...interface{}) string {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...interface{}) string); ok {
		r0 = rf(format, a...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
