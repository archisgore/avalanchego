// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	ids "github.com/ava-labs/avalanchego/ids"
	common "github.com/ava-labs/avalanchego/snow/engine/common"

	manager "github.com/ava-labs/avalanchego/database/manager"

	mock "github.com/stretchr/testify/mock"

	snow "github.com/ava-labs/avalanchego/snow"

	snowman "github.com/ava-labs/avalanchego/snow/consensus/snowman"

	time "time"

	version "github.com/ava-labs/avalanchego/version"
)

// ChainVM is an autogenerated mock type for the ChainVM type
type ChainVM struct {
	mock.Mock
}

// AppGossip provides a mock function with given fields: nodeID, msg
func (_m *ChainVM) AppGossip(nodeID ids.ShortID, msg []byte) error {
	ret := _m.Called(nodeID, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, []byte) error); ok {
		r0 = rf(nodeID, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppRequest provides a mock function with given fields: nodeID, requestID, deadline, request
func (_m *ChainVM) AppRequest(nodeID ids.ShortID, requestID uint32, deadline time.Time, request []byte) error {
	ret := _m.Called(nodeID, requestID, deadline, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, time.Time, []byte) error); ok {
		r0 = rf(nodeID, requestID, deadline, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppRequestFailed provides a mock function with given fields: nodeID, requestID
func (_m *ChainVM) AppRequestFailed(nodeID ids.ShortID, requestID uint32) error {
	ret := _m.Called(nodeID, requestID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32) error); ok {
		r0 = rf(nodeID, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppResponse provides a mock function with given fields: nodeID, requestID, response
func (_m *ChainVM) AppResponse(nodeID ids.ShortID, requestID uint32, response []byte) error {
	ret := _m.Called(nodeID, requestID, response)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, []byte) error); ok {
		r0 = rf(nodeID, requestID, response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Bootstrapped provides a mock function with given fields:
func (_m *ChainVM) Bootstrapped() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Bootstrapping provides a mock function with given fields:
func (_m *ChainVM) Bootstrapping() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BuildBlock provides a mock function with given fields:
func (_m *ChainVM) BuildBlock() (snowman.Block, error) {
	ret := _m.Called()

	var r0 snowman.Block
	if rf, ok := ret.Get(0).(func() snowman.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(snowman.Block)
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

// Connected provides a mock function with given fields: id, nodeVersion
func (_m *ChainVM) Connected(id ids.ShortID, nodeVersion version.Application) error {
	ret := _m.Called(id, nodeVersion)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, version.Application) error); ok {
		r0 = rf(id, nodeVersion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateHandlers provides a mock function with given fields:
func (_m *ChainVM) CreateHandlers() (map[string]*common.HTTPHandler, error) {
	ret := _m.Called()

	var r0 map[string]*common.HTTPHandler
	if rf, ok := ret.Get(0).(func() map[string]*common.HTTPHandler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*common.HTTPHandler)
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

// CreateStaticHandlers provides a mock function with given fields:
func (_m *ChainVM) CreateStaticHandlers() (map[string]*common.HTTPHandler, error) {
	ret := _m.Called()

	var r0 map[string]*common.HTTPHandler
	if rf, ok := ret.Get(0).(func() map[string]*common.HTTPHandler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*common.HTTPHandler)
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

// Disconnected provides a mock function with given fields: id
func (_m *ChainVM) Disconnected(id ids.ShortID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBlock provides a mock function with given fields: _a0
func (_m *ChainVM) GetBlock(_a0 ids.ID) (snowman.Block, error) {
	ret := _m.Called(_a0)

	var r0 snowman.Block
	if rf, ok := ret.Get(0).(func(ids.ID) snowman.Block); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(snowman.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ids.ID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthCheck provides a mock function with given fields:
func (_m *ChainVM) HealthCheck() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
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

// Initialize provides a mock function with given fields: ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender
func (_m *ChainVM) Initialize(ctx *snow.Context, dbManager manager.Manager, genesisBytes []byte, upgradeBytes []byte, configBytes []byte, toEngine chan<- common.Message, fxs []*common.Fx, appSender common.AppSender) error {
	ret := _m.Called(ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender)

	var r0 error
	if rf, ok := ret.Get(0).(func(*snow.Context, manager.Manager, []byte, []byte, []byte, chan<- common.Message, []*common.Fx, common.AppSender) error); ok {
		r0 = rf(ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LastAccepted provides a mock function with given fields:
func (_m *ChainVM) LastAccepted() (ids.ID, error) {
	ret := _m.Called()

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func() ids.ID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
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

// ParseBlock provides a mock function with given fields: _a0
func (_m *ChainVM) ParseBlock(_a0 []byte) (snowman.Block, error) {
	ret := _m.Called(_a0)

	var r0 snowman.Block
	if rf, ok := ret.Get(0).(func([]byte) snowman.Block); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(snowman.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetPreference provides a mock function with given fields: _a0
func (_m *ChainVM) SetPreference(_a0 ids.ID) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ID) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields:
func (_m *ChainVM) Shutdown() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Version provides a mock function with given fields:
func (_m *ChainVM) Version() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
