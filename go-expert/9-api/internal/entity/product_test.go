package entity

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductSuiteTest struct {
	suite.Suite
}

func (suite ProductSuiteTest) Test_deve_inicializar_um_produto() {
	product, err := NewProduct("product test a", 10.0)
	suite.Nil(err)
	suite.NotNil(product)
	suite.NotEmpty(product.ID)
	suite.Equal("product test a", product.Name)
	suite.Equal(10.0, product.Price)
	suite.NotNil(product.CreatedAt)
}

func (suite ProductSuiteTest) Test_deve_retornar_um_erro_nome_e_obrigatorio() {
	product, err := NewProduct("", 10.0)
	suite.Nil(product)
	suite.ErrorContains(err, "name is required")
}

func (suite ProductSuiteTest) Test_deve_retonar_um_erro_preco_e_obrigatorio() {
	product, err := NewProduct("Teste A", 0)
	suite.Nil(product)
	suite.ErrorContains(err, "price is required")
}

func (suite ProductSuiteTest) Test_deve_retonar_um_erro_preco_invalido() {
	product, err := NewProduct("Teste A", -10)
	suite.Nil(product)
	suite.ErrorContains(err, "invalid price")
}

func TestProductSuiteTest(t *testing.T) {
	suite.Run(t, new(ProductSuiteTest))
}
