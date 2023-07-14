package product_test

import (
	"context"
	"database/sql"

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

		// busca um array de produtos e espera que seu tamanho seja maior que 1
		products, err := repository.GetAll(ctx)

		assert.NoError(t, err)
		assert.True(t, len(products) > 1)
	})

}

func TestProductGet(t *testing.T) {
	t.Run("It should get one product by it's id", func(t *testing.T) {

		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		// Salva um produto
		productId, err := repository.Save(ctx, expectedProductResult)
		assert.NoError(t, err)

		// Busca o produto que foi salvo pelo seu Id e verifica seu productCode
		product, err := repository.Get(ctx, productId)
		assert.NoError(t, err)
		assert.Equal(t, "PROD03", product.ProductCode)
	})
	t.Run("It should return an error when there is no product in the database.", func(t *testing.T) {
		expectedMessage := product.ErrNotFound.Error()
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		// busca um produto de Id que inexistente
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

		// Salva um produto
		productId, err := repository.Save(ctx, expectedProductResult)
		assert.NoError(t, err)

		//Deleta o produto que foi criado
		err = repository.Delete(ctx, productId)
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

	t.Run("should update a product", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		// Salva um produto
		id, err := repository.Save(ctx, expectedProductResult)
		assert.NoError(t, err)
		expectedProductResult.ID = id
		expectedProductResult.Description = "teste"

		// update a product and expect no error
		err = repository.Update(ctx, expectedProductResult)
		assert.NoError(t, err)

		// busca o produto atualizado e verifica o campo product_code
		product, err := repository.Get(ctx, expectedProductResult.ID)
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, expectedProductResult.ProductCode, product.ProductCode)
	})
	t.Run("It should return an error when update a product that does not exist", func(t *testing.T) {

		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		expectedProductResult.ID = 50000000

		// atualiaza um  produto cujo id não existe e espera receber erro not found
		expectedErrorMessage := product.ErrNotFound.Error()
		err := repository.Update(ctx, expectedProductResult)
		assert.Error(t, err)
		assert.Equal(t, expectedErrorMessage, err.Error())
	})
}

func TestExistsProduct(t *testing.T) {
	t.Run("Is should return true if a product exist by it's productCode", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		// Salva um produto
		_, err := repository.Save(ctx, expectedProductResult)
		assert.NoError(t, err)

		// verifica se o produto salvo existe pelo seu productCode, espera receber true
		existsResult := repository.Exists(ctx, "PROD03")
		assert.True(t, existsResult)
	})

	t.Run("Should return error if the product does not exist", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		// verifica se o produto  existe, passa o productCode vazio e espera receber falso
		existsResult := repository.Exists(ctx, "")
		assert.False(t, existsResult)
	})
}

func TestExistsByIdProduct(t *testing.T) {
	t.Run("Is should return true if a product exist by it's Id", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		// Salva um produto
		productId, err := repository.Save(ctx, expectedProductResult)
		assert.NoError(t, err)

		// verifica se o produto criado existe pelo seu Id, espera receber true
		exists := repository.ExistsById(productId)
		assert.True(t, exists)
	})
	t.Run("Should return false if a product does not exists", func(t *testing.T) {
		repository := product.NewRepository(db)

		// verifica se um produto de Id 0 existe, espera receber false
		exists := repository.ExistsById(0)
		assert.False(t, exists)
	})
}

func TestAllEndpointsWithClosedDataBase(t *testing.T) {
	db.Close()

	t.Run("Should return error when there is an GetAll database error", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.GetAll(ctx)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Get database error", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Get(ctx, 200)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Save database error", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Save(ctx, expectedProduct)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Update database error", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Update(ctx, expectedProductResult)

		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Delete database error", func(t *testing.T) {
		repository := product.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Delete(ctx, expectedProductResult.ID)

		assert.Error(t, err)
	})
}
func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
