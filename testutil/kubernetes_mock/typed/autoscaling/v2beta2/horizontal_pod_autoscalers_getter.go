// Code generated by mockery v1.1.2. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
)

// HorizontalPodAutoscalersGetter is an autogenerated mock type for the HorizontalPodAutoscalersGetter type
type HorizontalPodAutoscalersGetter struct {
	mock.Mock
}

// HorizontalPodAutoscalers provides a mock function with given fields: namespace
func (_m *HorizontalPodAutoscalersGetter) HorizontalPodAutoscalers(namespace string) v2beta2.HorizontalPodAutoscalerInterface {
	ret := _m.Called(namespace)

	var r0 v2beta2.HorizontalPodAutoscalerInterface
	if rf, ok := ret.Get(0).(func(string) v2beta2.HorizontalPodAutoscalerInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2beta2.HorizontalPodAutoscalerInterface)
		}
	}

	return r0
}