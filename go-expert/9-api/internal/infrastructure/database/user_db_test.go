package database

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserDbTestSuite struct {
	suite.Suite
	db      *sql.DB
	configs configs.EnvConfig
}

func (suite *UserDbTestSuite) SetupSuite() {
	suite.configs = configs.LoadConfig("../../../cmd/server/test.env")
	configs.MigrationUP(suite.configs)
	suite.db = configs.DatabaseInitialize(suite.configs)
}

func (suite *UserDbTestSuite) TearDownTest() {
	stmt, err := suite.db.Prepare("DELETE FROM user WHERE id IS NOT NULL")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

func (suite *UserDbTestSuite) TearDownSuite() {
	configs.MigrationDown(suite.configs)
}

func (suite *UserDbTestSuite) Test_deve_salvar_um_usuario_no_banco_de_dados() {
	repository := NewUserRepository(suite.db)
	user, err := entity.NewUser("John Doh", "test@test.com", "123456")
	suite.Nil(err)
	err = repository.Create(user)
	suite.Nil(err)
	userFound, err := repository.FindByEmail(user.Email)
	suite.Nil(err)
	suite.Equal(user, userFound)
}

func TestUserDbTestSuite(t *testing.T) {
	suite.Run(t, new(UserDbTestSuite))
}
