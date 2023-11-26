package entity

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserSuitTest struct {
	suite.Suite
}

func (suite *UserSuitTest) Test_deve_inicializar_um_usuario() {
	user, err := NewUser("John Doe", "john.doe@test.com", "123456")
	suite.Nil(err)
	suite.NotNil(user)
	suite.NotEmpty(user.Id)
	suite.NotEmpty(user.Password)
	suite.Equal("John Doe", user.Name)
	suite.Equal("john.doe@test.com", user.Email)
}

func (suite *UserSuitTest) Test_deve_validar_a_senha() {
	user, err := NewUser("John Doe", "john.doe@test.com", "123456")
	suite.Nil(err)
	suite.True(user.ValidatePassword("123456"))
	suite.False(user.ValidatePassword("123456789"))
	suite.NotEqual("123456", user.Password)
}

func TestUserSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(UserSuitTest))
}
