// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	categoryItems "github.com/daffashafwan/pointcuan/business/categoryItems"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, domain
func (_m *Repository) Create(ctx context.Context, domain *categoryItems.Domain) (categoryItems.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 categoryItems.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *categoryItems.Domain) categoryItems.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(categoryItems.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *categoryItems.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]categoryItems.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []categoryItems.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []categoryItems.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]categoryItems.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *Repository) GetById(ctx context.Context, id int) (categoryItems.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 categoryItems.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) categoryItems.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(categoryItems.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, domain
func (_m *Repository) Update(ctx context.Context, domain categoryItems.Domain) (categoryItems.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 categoryItems.Domain
	if rf, ok := ret.Get(0).(func(context.Context, categoryItems.Domain) categoryItems.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(categoryItems.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, categoryItems.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
