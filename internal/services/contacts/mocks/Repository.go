// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	contact "my/crm-golang/internal/models/contact"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Repository) Create(_a0 *contact.Contact) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*contact.Contact) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByName provides a mock function with given fields: name
func (_m *Repository) DeleteByName(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() ([]*contact.Contact, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*contact.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*contact.Contact, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*contact.Contact); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contact.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name
func (_m *Repository) GetByName(name string) (*contact.Contact, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *contact.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*contact.Contact, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *contact.Contact); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contact.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastContacts provides a mock function with given fields: count
func (_m *Repository) GetLastContacts(count uint) ([]*contact.Contact, error) {
	ret := _m.Called(count)

	if len(ret) == 0 {
		panic("no return value specified for GetLastContacts")
	}

	var r0 []*contact.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]*contact.Contact, error)); ok {
		return rf(count)
	}
	if rf, ok := ret.Get(0).(func(uint) []*contact.Contact); ok {
		r0 = rf(count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contact.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(count)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, contactUpdateData
func (_m *Repository) Update(_a0 *contact.Contact, contactUpdateData *contact.ContactUpdateData) error {
	ret := _m.Called(_a0, contactUpdateData)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*contact.Contact, *contact.ContactUpdateData) error); ok {
		r0 = rf(_a0, contactUpdateData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
