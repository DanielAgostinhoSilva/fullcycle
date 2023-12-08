package mock

import "github.com/stretchr/testify/mock"

type UseCaseMock struct {
	mock.Mock
}

func (u *UseCaseMock) Execute(input interface{}) (*interface{}, error) {
	args := u.Called(input)
	return args.Get(0).(*interface{}), args.Error(1)
}

type NullaryUseCaseMock struct {
	mock.Mock
}

func (u *NullaryUseCaseMock) Execute() (interface{}, error) {
	args := u.Called()
	return args.Get(0).(*interface{}), args.Error(1)
}

type UnitUseCaseMock struct {
	mock.Mock
}

func (u *UnitUseCaseMock) Execute(input interface{}) error {
	args := u.Called(input)
	return args.Error(0)
}
