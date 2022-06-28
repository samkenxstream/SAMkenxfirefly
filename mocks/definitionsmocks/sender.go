// Code generated by mockery v1.0.0. DO NOT EDIT.

package definitionsmocks

import (
	context "context"

	definitions "github.com/hyperledger/firefly/internal/definitions"
	core "github.com/hyperledger/firefly/pkg/core"

	mock "github.com/stretchr/testify/mock"
)

// Sender is an autogenerated mock type for the Sender type
type Sender struct {
	mock.Mock
}

// CreateDefinitionWithIdentity provides a mock function with given fields: ctx, def, signingIdentity, tag, waitConfirm
func (_m *Sender) CreateDefinitionWithIdentity(ctx context.Context, def core.Definition, signingIdentity *core.SignerRef, tag string, waitConfirm bool) (*core.Message, error) {
	ret := _m.Called(ctx, def, signingIdentity, tag, waitConfirm)

	var r0 *core.Message
	if rf, ok := ret.Get(0).(func(context.Context, core.Definition, *core.SignerRef, string, bool) *core.Message); ok {
		r0 = rf(ctx, def, signingIdentity, tag, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, core.Definition, *core.SignerRef, string, bool) error); ok {
		r1 = rf(ctx, def, signingIdentity, tag, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateIdentityClaim provides a mock function with given fields: ctx, def, signingIdentity, tag, waitConfirm
func (_m *Sender) CreateIdentityClaim(ctx context.Context, def *core.IdentityClaim, signingIdentity *core.SignerRef, tag string, waitConfirm bool) (*core.Message, error) {
	ret := _m.Called(ctx, def, signingIdentity, tag, waitConfirm)

	var r0 *core.Message
	if rf, ok := ret.Get(0).(func(context.Context, *core.IdentityClaim, *core.SignerRef, string, bool) *core.Message); ok {
		r0 = rf(ctx, def, signingIdentity, tag, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.IdentityClaim, *core.SignerRef, string, bool) error); ok {
		r1 = rf(ctx, def, signingIdentity, tag, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DefineContractAPI provides a mock function with given fields: ctx, httpServerURL, api, waitConfirm
func (_m *Sender) DefineContractAPI(ctx context.Context, httpServerURL string, api *core.ContractAPI, waitConfirm bool) error {
	ret := _m.Called(ctx, httpServerURL, api, waitConfirm)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *core.ContractAPI, bool) error); ok {
		r0 = rf(ctx, httpServerURL, api, waitConfirm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DefineDatatype provides a mock function with given fields: ctx, datatype, waitConfirm
func (_m *Sender) DefineDatatype(ctx context.Context, datatype *core.Datatype, waitConfirm bool) error {
	ret := _m.Called(ctx, datatype, waitConfirm)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Datatype, bool) error); ok {
		r0 = rf(ctx, datatype, waitConfirm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DefineFFI provides a mock function with given fields: ctx, ffi, waitConfirm
func (_m *Sender) DefineFFI(ctx context.Context, ffi *core.FFI, waitConfirm bool) error {
	ret := _m.Called(ctx, ffi, waitConfirm)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.FFI, bool) error); ok {
		r0 = rf(ctx, ffi, waitConfirm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DefineTokenPool provides a mock function with given fields: ctx, pool, waitConfirm
func (_m *Sender) DefineTokenPool(ctx context.Context, pool *core.TokenPoolAnnouncement, waitConfirm bool) error {
	ret := _m.Called(ctx, pool, waitConfirm)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.TokenPoolAnnouncement, bool) error); ok {
		r0 = rf(ctx, pool, waitConfirm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Init provides a mock function with given fields: handler
func (_m *Sender) Init(handler definitions.Handler) {
	_m.Called(handler)
}

// Name provides a mock function with given fields:
func (_m *Sender) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
