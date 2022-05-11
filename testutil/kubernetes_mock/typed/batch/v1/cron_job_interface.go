// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	batchv1 "k8s.io/api/batch/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/batch/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// CronJobInterface is an autogenerated mock type for the CronJobInterface type
type CronJobInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, cronJob, opts
func (_m *CronJobInterface) Apply(ctx context.Context, cronJob *v1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (*batchv1.CronJob, error) {
	ret := _m.Called(ctx, cronJob, opts)

	var r0 *batchv1.CronJob
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CronJobApplyConfiguration, metav1.ApplyOptions) *batchv1.CronJob); ok {
		r0 = rf(ctx, cronJob, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*batchv1.CronJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CronJobApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, cronJob, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplyStatus provides a mock function with given fields: ctx, cronJob, opts
func (_m *CronJobInterface) ApplyStatus(ctx context.Context, cronJob *v1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (*batchv1.CronJob, error) {
	ret := _m.Called(ctx, cronJob, opts)

	var r0 *batchv1.CronJob
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CronJobApplyConfiguration, metav1.ApplyOptions) *batchv1.CronJob); ok {
		r0 = rf(ctx, cronJob, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*batchv1.CronJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CronJobApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, cronJob, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, cronJob, opts
func (_m *CronJobInterface) Create(ctx context.Context, cronJob *batchv1.CronJob, opts metav1.CreateOptions) (*batchv1.CronJob, error) {
	ret := _m.Called(ctx, cronJob, opts)

	var r0 *batchv1.CronJob
	if rf, ok := ret.Get(0).(func(context.Context, *batchv1.CronJob, metav1.CreateOptions) *batchv1.CronJob); ok {
		r0 = rf(ctx, cronJob, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*batchv1.CronJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *batchv1.CronJob, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, cronJob, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *CronJobInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *CronJobInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *CronJobInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*batchv1.CronJob, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *batchv1.CronJob
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *batchv1.CronJob); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*batchv1.CronJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *CronJobInterface) List(ctx context.Context, opts metav1.ListOptions) (*batchv1.CronJobList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *batchv1.CronJobList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *batchv1.CronJobList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*batchv1.CronJobList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *CronJobInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*batchv1.CronJob, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *batchv1.CronJob
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *batchv1.CronJob); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*batchv1.CronJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, cronJob, opts
func (_m *CronJobInterface) Update(ctx context.Context, cronJob *batchv1.CronJob, opts metav1.UpdateOptions) (*batchv1.CronJob, error) {
	ret := _m.Called(ctx, cronJob, opts)

	var r0 *batchv1.CronJob
	if rf, ok := ret.Get(0).(func(context.Context, *batchv1.CronJob, metav1.UpdateOptions) *batchv1.CronJob); ok {
		r0 = rf(ctx, cronJob, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*batchv1.CronJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *batchv1.CronJob, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, cronJob, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: ctx, cronJob, opts
func (_m *CronJobInterface) UpdateStatus(ctx context.Context, cronJob *batchv1.CronJob, opts metav1.UpdateOptions) (*batchv1.CronJob, error) {
	ret := _m.Called(ctx, cronJob, opts)

	var r0 *batchv1.CronJob
	if rf, ok := ret.Get(0).(func(context.Context, *batchv1.CronJob, metav1.UpdateOptions) *batchv1.CronJob); ok {
		r0 = rf(ctx, cronJob, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*batchv1.CronJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *batchv1.CronJob, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, cronJob, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *CronJobInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCronJobInterface creates a new instance of CronJobInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCronJobInterface(t testing.TB) *CronJobInterface {
	mock := &CronJobInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}