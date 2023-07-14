package productbatch_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	productbatch "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product_batch"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	db                  = initDatabase()
	repoQuerysIncorrect = productbatch.Querys{
		SaveQuery: "INVALID QUERY",
	}
)

func TestCreateProductBatchRepository(t *testing.T) {
	productBatchExpected := domain.ProductBatch{
		ID:                 1,
		ProductID:          1,
		SectionID:          1,
		BatchNumber:        1,
		CurrentQuantity:    1,
		InitialQuantity:    1,
		ManufacturingDate:  "2021-01-01",
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		DueDate:            "2021-01-01",
		ManufacturingHour:  1,
	}
	t.Run("Should create a new product batch", func(t *testing.T) {
		repository := productbatch.NewRepository(db, productbatch.Querys{})
		_, err := repository.Save(productBatchExpected)
		assert.NoError(t, err)
	})
	t.Run("Should fail when query is invalid", func(t *testing.T) {
		repository := productbatch.NewRepository(db, repoQuerysIncorrect)
		_, err := repository.Save(productBatchExpected)
		assert.Error(t, err)
	})
}

func initDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", "txdb")
	return db
}
