package productrecord_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	productrecord "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_record"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

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

var db = InitDatabase()

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
