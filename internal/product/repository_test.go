package product_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db = InitDatabase()

var expectedProductResult = domain.Product{
	Description:    "milk",
	ExpirationRate: 1,
	FreezingRate:   2,
	Height:         6.4,
	Length:         4.5,
	Netweight:      3.4,
	ProductCode:    "PROD03",
	RecomFreezTemp: 1.3,
	Width:          1.2,
	ProductTypeID:  1,
	SellerID:       1,
}

func TestProductsGetAll(t *testing.T) {
	t.Run("Should get all products", func(t *testing.T) {

		repository := product.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		products, err := repository.GetAll(ctx)
		fmt.Println("products get all", products)
		assert.NoError(t, err)
		assert.True(t, len(products) > 1)
	})
}

func TestProductGet(t *testing.T) {
	t.Run("It should get one product by it's id", func(t *testing.T) {
		id := 3
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		product, err := repository.Get(ctx, id)
		fmt.Println("product get", product)
		assert.NoError(t, err)
		assert.Equal(t, "PROD08", product.ProductCode)
	})
	t.Run("It should return an error when there is no product in the database.", func(t *testing.T) {
		expectedMessage := product.ErrNotFound.Error()
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		_, err := repository.Get(ctx, 50000000)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestProductSave(t *testing.T) {
	t.Run("should create a product", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		// Salva um produto e espera não obter erro
		result, err := repository.Save(ctx, expectedProductResult)
		assert.NoError(t, err)

		// Busca o produto que foi salvo e verifica o campo product_code
		product, err := repository.Get(ctx, result)
		assert.NoError(t, err)
		assert.NotNil(t, product)
		expectedProductResult.ProductCode = "PROD03"
		assert.Equal(t, expectedProductResult.ProductCode, product.ProductCode)
	})
}

func TestProductDelete(t *testing.T) {
	t.Run("It should delete a product", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		//Deleta um produto
		expectedProductResult.ID = 3
		err := repository.Delete(ctx, expectedProductResult.ID)
		assert.NoError(t, err)

		//Busca o produto que foi deletado, e verifica se não o encontra,
		_, err = repository.Get(ctx, expectedProductResult.ID)
		assert.Error(t, err)
		expectedErrorMessage := product.ErrNotFound.Error()
		assert.Equal(t, expectedErrorMessage, err.Error())
	})
	t.Run("It should return an error when delete a product that does not exist", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		// Deleta um produto que não existe e espera erro not found
		err := repository.Delete(ctx, 50000000)
		assert.Error(t, err)
		expectedErrorMessage := product.ErrNotFound.Error()
		assert.Equal(t, expectedErrorMessage, err.Error())
	})
}

func TestProductUpdate(t *testing.T) {
	expectedProducts := []domain.Product{
		{
			ID:             1,
			Description:    "milk",
			ExpirationRate: 1,
			FreezingRate:   2,
			Height:         6.4,
			Length:         4.5,
			Netweight:      3.4,
			ProductCode:    "PROD01",
			RecomFreezTemp: 1.3,
			Width:          1.2,
			ProductTypeID:  1,
			SellerID:       1,
		},
		{
			ID:             2,
			Description:    "milk",
			ExpirationRate: 1,
			FreezingRate:   2,
			Height:         6.4,
			Length:         4.5,
			Netweight:      3.4,
			ProductCode:    "PROD02",
			RecomFreezTemp: 1.3,
			Width:          1.2,
			ProductTypeID:  2,
			SellerID:       2,
		},
	}
	t.Run("should update a product", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		// update a product and expect no error
		err := repository.Update(ctx, expectedProducts[0])

		assert.NoError(t, err)

		// busca o produto atualizado e verifica o campo product_code
		product, err := repository.Get(ctx, expectedProducts[0].ID)
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, expectedProducts[0].ID, product.ID)
	})
	t.Run("It should return an error when update a product that does not exist", func(t *testing.T) {

		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		expectedProducts[0].ID = 50000000

		// atualiaza um  produto cujo id não existe e espera receber erro not found
		expectedErrorMessage := product.ErrNotFound.Error()
		err := repository.Update(ctx, expectedProductResult)
		assert.Error(t, err)
		assert.Equal(t, expectedErrorMessage, err.Error())
	})
}

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
