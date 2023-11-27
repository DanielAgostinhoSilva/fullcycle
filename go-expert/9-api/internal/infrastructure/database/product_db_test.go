package database

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductDbSuiteTest struct {
	suite.Suite
	db      *sql.DB
	configs configs.EnvConfig
}

func (suite *ProductDbSuiteTest) SetupSuite() {
	suite.configs = configs.LoadConfig("../../../cmd/server/test.env")
	configs.MigrationUP(suite.configs)
	suite.db = configs.DatabaseInitialize(suite.configs)
}

func (suite *ProductDbSuiteTest) TearDownTest() {
	stmt, err := suite.db.Prepare("DELETE FROM product WHERE id IS NOT NULL")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

func (suite *ProductDbSuiteTest) TearDownSuite() {
	configs.MigrationDown(suite.configs)
}

func (suite ProductDbSuiteTest) Test_deve_salva_um_produto_no_banco_de_dados() {
	repository := NewProductRepository(suite.db)
	product, _ := entity.NewProduct("product test a", 10.00)
	err := repository.Create(product)
	suite.Nil(err)
	productFound, err := repository.FindByID(product.ID)
	suite.Nil(err)
	suite.Equal(productFound.ID, product.ID)
	suite.Equal(productFound.Name, product.Name)
	suite.Equal(productFound.Price, product.Price)
	suite.Equal(productFound.CreatedAt, product.CreatedAt)
}

func (suite ProductDbSuiteTest) Test_deve_buscar_todos_os_produtos_no_banco_de_dados() {
	repository := NewProductRepository(suite.db)
	productA, _ := entity.NewProduct("product test a", 10.00)
	productB, _ := entity.NewProduct("product test B", 22.00)
	err := repository.Create(productA)
	suite.Nil(err)
	err = repository.Create(productB)
	suite.Nil(err)
	products, err := repository.FindAll()
	suite.Nil(err)
	suite.Len(products, 2)
}

func (suite ProductDbSuiteTest) Test_deve_deletar_um_produto_pelo_id() {
	repository := NewProductRepository(suite.db)
	product, _ := entity.NewProduct("product test", 10.00)
	err := repository.Create(product)
	suite.Nil(err)

	err = repository.DeleteById(product.ID)
	suite.Nil(err)

	products, err := repository.FindAll()
	suite.Nil(err)
	suite.Len(products, 0)

}

func TestProductDbSuiteTest(t *testing.T) {
	suite.Run(t, new(ProductDbSuiteTest))
}
