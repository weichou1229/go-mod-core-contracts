// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v4/errors"

	mock "github.com/stretchr/testify/mock"

	requests "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"

	responses "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
)

// IntervalClient is an autogenerated mock type for the IntervalClient type
type IntervalClient struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, reqs
func (_m *IntervalClient) Add(ctx context.Context, reqs []requests.AddIntervalRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseWithIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.AddIntervalRequest) []common.BaseWithIdResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseWithIdResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.AddIntervalRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllIntervals provides a mock function with given fields: ctx, offset, limit
func (_m *IntervalClient) AllIntervals(ctx context.Context, offset int, limit int) (responses.MultiIntervalsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, offset, limit)

	var r0 responses.MultiIntervalsResponse
	if rf, ok := ret.Get(0).(func(context.Context, int, int) responses.MultiIntervalsResponse); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiIntervalsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteIntervalByName provides a mock function with given fields: ctx, name
func (_m *IntervalClient) DeleteIntervalByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// IntervalByName provides a mock function with given fields: ctx, name
func (_m *IntervalClient) IntervalByName(ctx context.Context, name string) (responses.IntervalResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 responses.IntervalResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) responses.IntervalResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(responses.IntervalResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, reqs
func (_m *IntervalClient) Update(ctx context.Context, reqs []requests.UpdateIntervalRequest) ([]common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.UpdateIntervalRequest) []common.BaseResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.UpdateIntervalRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewIntervalClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewIntervalClient creates a new instance of IntervalClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIntervalClient(t mockConstructorTestingTNewIntervalClient) *IntervalClient {
	mock := &IntervalClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
