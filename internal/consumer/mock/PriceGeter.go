// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	model "github.com/mmfshirokan/positionService/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// PriceGeter is an autogenerated mock type for the PriceGeter type
type PriceGeter struct {
	mock.Mock
}

type PriceGeter_Expecter struct {
	mock *mock.Mock
}

func (_m *PriceGeter) EXPECT() *PriceGeter_Expecter {
	return &PriceGeter_Expecter{mock: &_m.Mock}
}

// GetAllChanForSymb provides a mock function with given fields: symb
func (_m *PriceGeter) GetAllChanForSymb(symb string) ([]chan model.Price, bool) {
	ret := _m.Called(symb)

	if len(ret) == 0 {
		panic("no return value specified for GetAllChanForSymb")
	}

	var r0 []chan model.Price
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) ([]chan model.Price, bool)); ok {
		return rf(symb)
	}
	if rf, ok := ret.Get(0).(func(string) []chan model.Price); ok {
		r0 = rf(symb)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chan model.Price)
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(symb)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// PriceGeter_GetAllChanForSymb_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllChanForSymb'
type PriceGeter_GetAllChanForSymb_Call struct {
	*mock.Call
}

// GetAllChanForSymb is a helper method to define mock.On call
//   - symb string
func (_e *PriceGeter_Expecter) GetAllChanForSymb(symb interface{}) *PriceGeter_GetAllChanForSymb_Call {
	return &PriceGeter_GetAllChanForSymb_Call{Call: _e.mock.On("GetAllChanForSymb", symb)}
}

func (_c *PriceGeter_GetAllChanForSymb_Call) Run(run func(symb string)) *PriceGeter_GetAllChanForSymb_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *PriceGeter_GetAllChanForSymb_Call) Return(chanels []chan model.Price, isSuccessfull bool) *PriceGeter_GetAllChanForSymb_Call {
	_c.Call.Return(chanels, isSuccessfull)
	return _c
}

func (_c *PriceGeter_GetAllChanForSymb_Call) RunAndReturn(run func(string) ([]chan model.Price, bool)) *PriceGeter_GetAllChanForSymb_Call {
	_c.Call.Return(run)
	return _c
}

// NewPriceGeter creates a new instance of PriceGeter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPriceGeter(t interface {
	mock.TestingT
	Cleanup(func())
}) *PriceGeter {
	mock := &PriceGeter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
