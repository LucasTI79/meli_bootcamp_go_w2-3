package purchase_orders_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	purchaseOrders "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/purchase_orders"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db = initDatabase()

func TestSaveOrders(t *testing.T) {
	t.Run("Should create orders", func(t *testing.T) {
		repositoryPurchaseOrders := purchaseOrders.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		expectedOrder := domain.PurchaseOrders{
			OrderNumber:     "9423i",
			OrderDate:       "2021-04-04",
			TrackingCode:    "afijaehn",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		err := repositoryPurchaseOrders.Save(ctx, expectedOrder)
		assert.NoError(t, err)
	})
}

func TestExistsOrder(t *testing.T) {

	t.Run("Should return true if order exists", func(t *testing.T) {
		repository := purchaseOrders.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		expectedOrder := domain.PurchaseOrders{
			OrderNumber:     "9423i",
			OrderDate:       "2021-04-04",
			TrackingCode:    "afijaehn",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		err := repository.Save(ctx, expectedOrder)
		assert.NoError(t, err)

		exists := repository.ExistsOrder(ctx, expectedOrder.OrderNumber)
		assert.True(t, exists)
	})
	t.Run("Should return false if buyer not exists", func(t *testing.T) {
		repository := purchaseOrders.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		exists := repository.ExistsOrder(ctx, "0248945")
		assert.False(t, exists)
	})
}

func initDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
