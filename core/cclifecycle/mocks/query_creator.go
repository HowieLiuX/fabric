/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by mockery v1.0.0
package mocks

import cc "github.com/hyperledger/fabric/core/cclifecycle"
import mock "github.com/stretchr/testify/mock"

// QueryCreator is an autogenerated mock type for the QueryCreator type
type QueryCreator struct {
	mock.Mock
}

// NewQuery provides a mock function with given fields:
func (_m *QueryCreator) NewQuery() (cc.Query, error) {
	ret := _m.Called()

	var r0 cc.Query
	if rf, ok := ret.Get(0).(func() cc.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cc.Query)
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
