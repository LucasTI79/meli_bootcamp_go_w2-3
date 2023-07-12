package product_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/product"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db = InitDatabase()

// var expectedProduct = domain.Product{
// 	// ID:             1,
// 	Description:    "milk",
// 	ExpirationRate: 1,
// 	FreezingRate:   2,
// 	Height:         6.4,
// 	Length:         4.5,
// 	Netweight:      3.4,
// 	ProductCode:    "PROD03",
// 	RecomFreezTemp: 1.3,
// 	Width:          1.2,
// 	ProductTypeID:  1,
// 	SellerID:       1,
// }

// expectedProducts := []domain.Product{
// 	{
// 		ID:             1,
// 		Description:    "milk",
// 		ExpirationRate: 1,
// 		FreezingRate:   2,
// 		Height:         6.4,
// 		Length:         4.5,
// 		Netweight:      3.4,
// 		ProductCode:    "PROD01",
// 		RecomFreezTemp: 1.3,
// 		Width:          1.2,
// 		ProductTypeID:  1,
// 		SellerID:       1,
// 	},
// 	{
// 		ID:             2,
// 		Description:    "milk",
// 		ExpirationRate: 1,
// 		FreezingRate:   2,
// 		Height:         6.4,
// 		Length:         4.5,
// 		Netweight:      3.4,
// 		ProductCode:    "PROD02",
// 		RecomFreezTemp: 1.3,
// 		Width:          1.2,
// 		ProductTypeID:  2,
// 		SellerID:       2,
// 	},
// }

func TestProductsGetAll(t *testing.T) {
	t.Run("Should get all products", func(t *testing.T) {

		repository := product.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		products, err := repository.GetAll(ctx)
		assert.NoError(t, err)
		assert.True(t, len(products) > 1)
	})
}

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
