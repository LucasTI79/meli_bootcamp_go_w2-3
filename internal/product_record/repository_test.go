package productrecord_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	productrecord "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_record"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var expectedProductRecordResult = domain.ProductRecord{
	ID:             1,
	LastUpdateDate: "2021-04-04",
	PurchasePrice:  10,
	SalePrice:      15,
	ProductID:      1,
}

func TestGetAllProductReports(t *testing.T) {
	t.Run("Should get product reports of all products", func(t *testing.T) {

		repository := productrecord.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		productRecords, err := repository.RecordsByAllProductsReport(ctx)
		assert.NoError(t, err)
		assert.True(t, len(productRecords) > 1)
	})
}

func TestProductGet(t *testing.T) {
	t.Run("It should get the report of a product record by the product ID.", func(t *testing.T) {
		id := 1
		repository := productrecord.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		productReport, err := repository.RecordsByOneProductReport(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, "Product 1", productReport.Description)
	})
	t.Run("It should return an error when the record of a product is not found.", func(t *testing.T) {
		repository := productrecord.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.RecordsByOneProductReport(ctx, 50000000)
		assert.Error(t, err)
		expectedErrorMessage := productrecord.ErrNotFound.Error()
		assert.Equal(t, expectedErrorMessage, err.Error())
	})
}

func TestProductSave(t *testing.T) {
	t.Run("should create a product", func(t *testing.T) {
		repository := productrecord.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		// Salva um produto e espera não obter erro
		_, err := repository.Save(ctx, expectedProductRecordResult)
		assert.NoError(t, err)

	})
}

func TestEndpointsWithDatabaseClosed(t *testing.T) {
	db.Close()
	t.Run("Should return error when there is an get all product report database error", func(t *testing.T) {
		repository := productrecord.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.RecordsByAllProductsReport(ctx)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Get database error", func(t *testing.T) {
		repository := productrecord.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.RecordsByOneProductReport(ctx, 50000000)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Save database error", func(t *testing.T) {
		repository := productrecord.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Save(ctx, expectedProductRecordResult)
		assert.Error(t, err)
	})

}

var db = InitDatabase()

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
