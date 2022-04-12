// Code generated by mockery v1.0.0. DO NOT EDIT.

package http

import (
	context "context"

	models "github.com/onflow/flow-go/engine/access/rest/models"
	mock "github.com/stretchr/testify/mock"
)

// mockHandler is an autogenerated mock type for the handler type
type mockHandler struct {
	mock.Mock
}

// executeScriptAtBlockHeight provides a mock function with given fields: ctx, height, script, arguments, opts
func (_m *mockHandler) executeScriptAtBlockHeight(ctx context.Context, height string, script string, arguments []string, opts ...queryOpts) (string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, height, script, arguments)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []string, ...queryOpts) string); ok {
		r0 = rf(ctx, height, script, arguments, opts...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, []string, ...queryOpts) error); ok {
		r1 = rf(ctx, height, script, arguments, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// executeScriptAtBlockID provides a mock function with given fields: ctx, ID, script, arguments, opts
func (_m *mockHandler) executeScriptAtBlockID(ctx context.Context, ID string, script string, arguments []string, opts ...queryOpts) (string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ID, script, arguments)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []string, ...queryOpts) string); ok {
		r0 = rf(ctx, ID, script, arguments, opts...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, []string, ...queryOpts) error); ok {
		r1 = rf(ctx, ID, script, arguments, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getAccount provides a mock function with given fields: ctx, address, height, opts
func (_m *mockHandler) getAccount(ctx context.Context, address string, height string, opts ...queryOpts) (*models.Account, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, address, height)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Account
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...queryOpts) *models.Account); ok {
		r0 = rf(ctx, address, height, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...queryOpts) error); ok {
		r1 = rf(ctx, address, height, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getBlockByID provides a mock function with given fields: ctx, ID, opts
func (_m *mockHandler) getBlockByID(ctx context.Context, ID string, opts ...queryOpts) (*models.Block, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Block
	if rf, ok := ret.Get(0).(func(context.Context, string, ...queryOpts) *models.Block); ok {
		r0 = rf(ctx, ID, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...queryOpts) error); ok {
		r1 = rf(ctx, ID, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getBlocksByHeights provides a mock function with given fields: ctx, heights, startHeight, endHeight, opts
func (_m *mockHandler) getBlocksByHeights(ctx context.Context, heights string, startHeight string, endHeight string, opts ...queryOpts) ([]*models.Block, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, heights, startHeight, endHeight)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*models.Block
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, ...queryOpts) []*models.Block); ok {
		r0 = rf(ctx, heights, startHeight, endHeight, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, ...queryOpts) error); ok {
		r1 = rf(ctx, heights, startHeight, endHeight, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getCollection provides a mock function with given fields: ctx, ID, opts
func (_m *mockHandler) getCollection(ctx context.Context, ID string, opts ...queryOpts) (*models.Collection, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Collection
	if rf, ok := ret.Get(0).(func(context.Context, string, ...queryOpts) *models.Collection); ok {
		r0 = rf(ctx, ID, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...queryOpts) error); ok {
		r1 = rf(ctx, ID, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getEvents provides a mock function with given fields: ctx, eventType, start, end, blockIDs, opts
func (_m *mockHandler) getEvents(ctx context.Context, eventType string, start string, end string, blockIDs []string, opts ...queryOpts) ([]models.BlockEvents, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, eventType, start, end, blockIDs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []models.BlockEvents
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []string, ...queryOpts) []models.BlockEvents); ok {
		r0 = rf(ctx, eventType, start, end, blockIDs, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.BlockEvents)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, []string, ...queryOpts) error); ok {
		r1 = rf(ctx, eventType, start, end, blockIDs, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getExecutionResultByID provides a mock function with given fields: ctx, id, opts
func (_m *mockHandler) getExecutionResultByID(ctx context.Context, id string, opts ...queryOpts) (*models.ExecutionResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.ExecutionResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...queryOpts) *models.ExecutionResult); ok {
		r0 = rf(ctx, id, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ExecutionResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...queryOpts) error); ok {
		r1 = rf(ctx, id, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getExecutionResults provides a mock function with given fields: ctx, blockIDs, opts
func (_m *mockHandler) getExecutionResults(ctx context.Context, blockIDs []string, opts ...queryOpts) ([]models.ExecutionResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, blockIDs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []models.ExecutionResult
	if rf, ok := ret.Get(0).(func(context.Context, []string, ...queryOpts) []models.ExecutionResult); ok {
		r0 = rf(ctx, blockIDs, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ExecutionResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, ...queryOpts) error); ok {
		r1 = rf(ctx, blockIDs, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getTransaction provides a mock function with given fields: ctx, ID, includeResult, opts
func (_m *mockHandler) getTransaction(ctx context.Context, ID string, includeResult bool, opts ...queryOpts) (*models.Transaction, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ID, includeResult)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, ...queryOpts) *models.Transaction); ok {
		r0 = rf(ctx, ID, includeResult, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool, ...queryOpts) error); ok {
		r1 = rf(ctx, ID, includeResult, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// sendTransaction provides a mock function with given fields: ctx, transaction, opts
func (_m *mockHandler) sendTransaction(ctx context.Context, transaction []byte, opts ...queryOpts) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, transaction)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, ...queryOpts) error); ok {
		r0 = rf(ctx, transaction, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
