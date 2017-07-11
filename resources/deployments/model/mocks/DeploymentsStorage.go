// Copyright 2017 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package mocks

import context "context"
import deployments "github.com/mendersoftware/deployments/resources/deployments"
import mock "github.com/stretchr/testify/mock"
import model "github.com/mendersoftware/deployments/resources/deployments/model"
import time "time"

// DeploymentsStorage is an autogenerated mock type for the DeploymentsStorage type
type DeploymentsStorage struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *DeploymentsStorage) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExistByArtifactId provides a mock function with given fields: ctx, id
func (_m *DeploymentsStorage) ExistByArtifactId(ctx context.Context, id string) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistUnfinishedByArtifactId provides a mock function with given fields: ctx, id
func (_m *DeploymentsStorage) ExistUnfinishedByArtifactId(ctx context.Context, id string) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, query
func (_m *DeploymentsStorage) Find(ctx context.Context, query deployments.Query) ([]*deployments.Deployment, error) {
	ret := _m.Called(ctx, query)

	var r0 []*deployments.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, deployments.Query) []*deployments.Deployment); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*deployments.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, deployments.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *DeploymentsStorage) FindByID(ctx context.Context, id string) (*deployments.Deployment, error) {
	ret := _m.Called(ctx, id)

	var r0 *deployments.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, string) *deployments.Deployment); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deployments.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUnfinishedByID provides a mock function with given fields: ctx, id
func (_m *DeploymentsStorage) FindUnfinishedByID(ctx context.Context, id string) (*deployments.Deployment, error) {
	ret := _m.Called(ctx, id)

	var r0 *deployments.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, string) *deployments.Deployment); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deployments.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Finish provides a mock function with given fields: ctx, id, when
func (_m *DeploymentsStorage) Finish(ctx context.Context, id string, when time.Time) error {
	ret := _m.Called(ctx, id, when)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) error); ok {
		r0 = rf(ctx, id, when)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: ctx, deployment
func (_m *DeploymentsStorage) Insert(ctx context.Context, deployment *deployments.Deployment) error {
	ret := _m.Called(ctx, deployment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *deployments.Deployment) error); ok {
		r0 = rf(ctx, deployment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStats provides a mock function with given fields: ctx, id, state_from, state_to
func (_m *DeploymentsStorage) UpdateStats(ctx context.Context, id string, state_from string, state_to string) error {
	ret := _m.Called(ctx, id, state_from, state_to)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, id, state_from, state_to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatsAndFinishDeployment provides a mock function with given fields: ctx, id, stats
func (_m *DeploymentsStorage) UpdateStatsAndFinishDeployment(ctx context.Context, id string, stats deployments.Stats) error {
	ret := _m.Called(ctx, id, stats)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, deployments.Stats) error); ok {
		r0 = rf(ctx, id, stats)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ model.DeploymentsStorage = (*DeploymentsStorage)(nil)
