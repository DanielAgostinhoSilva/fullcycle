package database

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/database"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/enviroment"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/configs/migration"
	"github.com/DanielAgostinhoSilva/fullcycle/9-api/internal/domain/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductDbSuiteTest struct {
	suite.Suite
	db      *sql.DB
	configs *enviroment.EnvConfig
}

func (suite *ProductDbSuiteTest) SetupSuite() {
	suite.configs = enviroment.LoadConfig("../../../cmd/server/test.env")
	migration.MigrationUP(suite.configs)
	suite.db = database.InitializeDatabase(suite.configs)
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
	migration.MigrationDown(suite.configs)
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

	productF, _ := entity.NewProduct("product test F", 22.00)
	repository.Create(productF)
	productB, _ := entity.NewProduct("product test B", 22.00)
	repository.Create(productB)
	productC, _ := entity.NewProduct("product test C", 22.00)
	repository.Create(productC)
	productD, _ := entity.NewProduct("product test D", 22.00)
	repository.Create(productD)
	productE, _ := entity.NewProduct("product test E", 22.00)
	repository.Create(productE)
	productG, _ := entity.NewProduct("product test G", 22.00)
	repository.Create(productG)
	productH, _ := entity.NewProduct("product test H", 22.00)
	repository.Create(productH)
	productA, _ := entity.NewProduct("product test A", 10.00)
	repository.Create(productA)

	products, pageable, err := repository.FindAll(1, 4, "ASC")
	suite.Nil(err)
	suite.Len(products, 4)
	suite.Equal(1, pageable.PageNumber)
	suite.Equal(4, pageable.PageSize)
	suite.Equal(8, pageable.TotalElements)
	suite.Equal(2, pageable.TotalPages)
	suite.Equal("product test A", products[0].Name)
	suite.Equal("product test B", products[1].Name)
	suite.Equal("product test C", products[2].Name)
	suite.Equal("product test D", products[3].Name)

	products, pageable, err = repository.FindAll(2, 4, "ASC")
	suite.Nil(err)
	suite.Len(products, 4)
	suite.Equal(2, pageable.PageNumber)
	suite.Equal(4, pageable.PageSize)
	suite.Equal(8, pageable.TotalElements)
	suite.Equal(2, pageable.TotalPages)
	suite.Equal("product test E", products[0].Name)
	suite.Equal("product test F", products[1].Name)
	suite.Equal("product test G", products[2].Name)
	suite.Equal("product test H", products[3].Name)
}

func (suite ProductDbSuiteTest) Test_deve_deletar_um_produto_pelo_id() {
	repository := NewProductRepository(suite.db)
	product, _ := entity.NewProduct("Z1", 10.00)
	err := repository.Create(product)
	suite.Nil(err)

	err = repository.DeleteById(product.ID)
	suite.Nil(err)

}

func (suite ProductDbSuiteTest) Test_deve_atualizar_o_nome_e_o_price_de_um_produto() {
	repository := NewProductRepository(suite.db)
	product, _ := entity.NewProduct("product test", 10.00)
	err := repository.Create(product)
	suite.Nil(err)

	productFound, err := repository.FindByID(product.ID)
	suite.Nil(err)
	suite.Equal(product, productFound)

	product.Name = "product test 2"
	product.Price = 20.00
	err = repository.Update(product)
	suite.Nil(err)

	productUpdated, err := repository.FindByID(product.ID)
	suite.Nil(err)
	suite.Equal(product.Name, productUpdated.Name)
	suite.Equal(product.Price, productUpdated.Price)
}

func TestProductDbSuiteTest(t *testing.T) {
	suite.Run(t, new(ProductDbSuiteTest))
}
