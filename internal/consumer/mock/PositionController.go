// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/mmfshirokan/positionService/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// PositionController is an autogenerated mock type for the PositionController type
type PositionController struct {
	mock.Mock
}

type PositionController_Expecter struct {
	mock *mock.Mock
}

func (_m *PositionController) EXPECT() *PositionController_Expecter {
	return &PositionController_Expecter{mock: &_m.Mock}
}

// GetAllOpened provides a mock function with given fields: ctx
func (_m *PositionController) GetAllOpened(ctx context.Context) ([]model.Position, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllOpened")
	}

	var r0 []model.Position
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Position, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Position); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Position)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PositionController_GetAllOpened_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllOpened'
type PositionController_GetAllOpened_Call struct {
	*mock.Call
}

// GetAllOpened is a helper method to define mock.On call
//   - ctx context.Context
func (_e *PositionController_Expecter) GetAllOpened(ctx interface{}) *PositionController_GetAllOpened_Call {
	return &PositionController_GetAllOpened_Call{Call: _e.mock.On("GetAllOpened", ctx)}
}

func (_c *PositionController_GetAllOpened_Call) Run(run func(ctx context.Context)) *PositionController_GetAllOpened_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *PositionController_GetAllOpened_Call) Return(_a0 []model.Position, _a1 error) *PositionController_GetAllOpened_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PositionController_GetAllOpened_Call) RunAndReturn(run func(context.Context) ([]model.Position, error)) *PositionController_GetAllOpened_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, position
func (_m *PositionController) Update(ctx context.Context, position model.Position) error {
	ret := _m.Called(ctx, position)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Position) error); ok {
		r0 = rf(ctx, position)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PositionController_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type PositionController_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - position model.Position
func (_e *PositionController_Expecter) Update(ctx interface{}, position interface{}) *PositionController_Update_Call {
	return &PositionController_Update_Call{Call: _e.mock.On("Update", ctx, position)}
}

func (_c *PositionController_Update_Call) Run(run func(ctx context.Context, position model.Position)) *PositionController_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Position))
	})
	return _c
}

func (_c *PositionController_Update_Call) Return(_a0 error) *PositionController_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PositionController_Update_Call) RunAndReturn(run func(context.Context, model.Position) error) *PositionController_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewPositionController creates a new instance of PositionController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPositionController(t interface {
	mock.TestingT
	Cleanup(func())
}) *PositionController {
	mock := &PositionController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
