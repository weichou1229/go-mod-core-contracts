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

// NotificationClient is an autogenerated mock type for the NotificationClient type
type NotificationClient struct {
	mock.Mock
}

// CleanupNotifications provides a mock function with given fields: ctx
func (_m *NotificationClient) CleanupNotifications(ctx context.Context) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context) common.BaseResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context) errors.EdgeX); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// CleanupNotificationsByAge provides a mock function with given fields: ctx, age
func (_m *NotificationClient) CleanupNotificationsByAge(ctx context.Context, age int) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, age)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, int) common.BaseResponse); ok {
		r0 = rf(ctx, age)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, int) errors.EdgeX); ok {
		r1 = rf(ctx, age)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteNotificationById provides a mock function with given fields: ctx, id
func (_m *NotificationClient) DeleteNotificationById(ctx context.Context, id string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, id)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteProcessedNotificationsByAge provides a mock function with given fields: ctx, age
func (_m *NotificationClient) DeleteProcessedNotificationsByAge(ctx context.Context, age int) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, age)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, int) common.BaseResponse); ok {
		r0 = rf(ctx, age)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, int) errors.EdgeX); ok {
		r1 = rf(ctx, age)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationById provides a mock function with given fields: ctx, id
func (_m *NotificationClient) NotificationById(ctx context.Context, id string) (responses.NotificationResponse, errors.EdgeX) {
	ret := _m.Called(ctx, id)

	var r0 responses.NotificationResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) responses.NotificationResponse); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(responses.NotificationResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationsByCategory provides a mock function with given fields: ctx, category, offset, limit
func (_m *NotificationClient) NotificationsByCategory(ctx context.Context, category string, offset int, limit int) (responses.MultiNotificationsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, category, offset, limit)

	var r0 responses.MultiNotificationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiNotificationsResponse); ok {
		r0 = rf(ctx, category, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiNotificationsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, category, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationsByLabel provides a mock function with given fields: ctx, label, offset, limit
func (_m *NotificationClient) NotificationsByLabel(ctx context.Context, label string, offset int, limit int) (responses.MultiNotificationsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, label, offset, limit)

	var r0 responses.MultiNotificationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiNotificationsResponse); ok {
		r0 = rf(ctx, label, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiNotificationsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, label, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationsByStatus provides a mock function with given fields: ctx, status, offset, limit
func (_m *NotificationClient) NotificationsByStatus(ctx context.Context, status string, offset int, limit int) (responses.MultiNotificationsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, status, offset, limit)

	var r0 responses.MultiNotificationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiNotificationsResponse); ok {
		r0 = rf(ctx, status, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiNotificationsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, status, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationsBySubscriptionName provides a mock function with given fields: ctx, subscriptionName, offset, limit
func (_m *NotificationClient) NotificationsBySubscriptionName(ctx context.Context, subscriptionName string, offset int, limit int) (responses.MultiNotificationsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, subscriptionName, offset, limit)

	var r0 responses.MultiNotificationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiNotificationsResponse); ok {
		r0 = rf(ctx, subscriptionName, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiNotificationsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, subscriptionName, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationsByTimeRange provides a mock function with given fields: ctx, start, end, offset, limit
func (_m *NotificationClient) NotificationsByTimeRange(ctx context.Context, start int, end int, offset int, limit int) (responses.MultiNotificationsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, start, end, offset, limit)

	var r0 responses.MultiNotificationsResponse
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int, int) responses.MultiNotificationsResponse); ok {
		r0 = rf(ctx, start, end, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiNotificationsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, int, int, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SendNotification provides a mock function with given fields: ctx, reqs
func (_m *NotificationClient) SendNotification(ctx context.Context, reqs []requests.AddNotificationRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseWithIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.AddNotificationRequest) []common.BaseWithIdResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseWithIdResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.AddNotificationRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewNotificationClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewNotificationClient creates a new instance of NotificationClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotificationClient(t mockConstructorTestingTNewNotificationClient) *NotificationClient {
	mock := &NotificationClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
