package buyer_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	buyers "github.com/extmatperez/meli_bootcamp_go_w2-3/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/purchase_orders"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db = initDatabase()

func TestGetBuyersOrdersRepository(t *testing.T) {
	t.Run("Should return orders by buyers", func(t *testing.T) {
		repositoryBuyer := buyers.NewRepository(db)
		repositoryPurchaseOrders := purchase_orders.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		expectedBuyer1 := domain.Buyer{
			CardNumberID: "2556",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}
		expectedBuyer2 := domain.Buyer{
			CardNumberID: "98424",
			FirstName:    "Giuli",
			LastName:     "Oli",
		}

		_, err := repositoryBuyer.Save(ctx, expectedBuyer1)
		assert.NoError(t, err)
		_, err = repositoryBuyer.Save(ctx, expectedBuyer2)
		assert.NoError(t, err)

		expectedOrder := domain.PurchaseOrders{
			OrderNumber:     "9423i",
			OrderDate:       "2021-04-04",
			TrackingCode:    "afijaehn",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		err = repositoryPurchaseOrders.Save(ctx, expectedOrder)
		assert.NoError(t, err)

		buyer, err := repositoryBuyer.GetBuyersOrders(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, buyer)
	})
}

func TestGetBuyerOrdersRepository(t *testing.T) {
	t.Run("Should return orders by buyer", func(t *testing.T) {
		repositoryBuyer := buyers.NewRepository(db)
		repositoryPurchaseOrders := purchase_orders.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		expectedBuyer := domain.Buyer{
			CardNumberID: "138935",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		buyerID, err := repositoryBuyer.Save(ctx, expectedBuyer)
		assert.NoError(t, err)

		expectedOrder := domain.PurchaseOrders{
			OrderNumber:     "9423i",
			OrderDate:       "2021-04-04",
			TrackingCode:    "afijaehn",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		err = repositoryPurchaseOrders.Save(ctx, expectedOrder)
		assert.NoError(t, err)

		buyer, err := repositoryBuyer.GetBuyerOrders(ctx, buyerID)
		assert.NoError(t, err)
		assert.NotNil(t, buyer)
	})
	t.Run("Should return error if buyer not exists", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		buyer, err := repository.GetBuyerOrders(ctx, 50)
		assert.Error(t, err)
		assert.Equal(t, domain.BuyerOrders{}, buyer)
	})

}

func TestDeleteBuyer(t *testing.T) {

	t.Run("Should delete buyer", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		expectedBuyer := domain.Buyer{
			CardNumberID: "138935",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		buyerID, err := repository.Save(ctx, expectedBuyer)
		assert.NoError(t, err)

		err = repository.Delete(ctx, buyerID)
		assert.NoError(t, err)

		buyer, err := repository.Get(ctx, buyerID)
		assert.Error(t, err)
		assert.Equal(t, domain.Buyer{}, buyer)
	})
	t.Run("Should return error if buyer not exists", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		err := repository.Delete(ctx, 50)
		assert.Error(t, err)
	})
}

func TestUpdateBuyer(t *testing.T) {

	t.Run("Should update buyer", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		oldBuyer := domain.Buyer{
			CardNumberID: "4198437",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		oldBuyerID, err := repository.Save(ctx, oldBuyer)
		assert.NoError(t, err)

		oldBuyer.FirstName = "Giu"
		oldBuyer.LastName = "Oli"
		oldBuyer.ID = oldBuyerID

		err = repository.Update(ctx, oldBuyer)
		assert.NoError(t, err)

		buyer, err := repository.Get(ctx, oldBuyerID)
		assert.NoError(t, err)
		assert.Equal(t, oldBuyer, buyer)
	})
	t.Run("Should return error if buyer not exists", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		_, err := repository.Get(ctx, 50)
		assert.Error(t, err)
	})
}

func TestExistsBuyer(t *testing.T) {

	t.Run("Should return true if buyer exists", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		expectedBuyer := domain.Buyer{
			CardNumberID: "138935",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		_, err := repository.Save(ctx, expectedBuyer)
		assert.NoError(t, err)

		exists := repository.ExistsBuyer(ctx, expectedBuyer.CardNumberID)
		assert.True(t, exists)
	})
	t.Run("Should return false if buyer not exists", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		exists := repository.ExistsBuyer(ctx, "0248945")
		assert.False(t, exists)
	})
}

func TestExistsBuyerID(t *testing.T) {

	t.Run("Should return true if buyer id exists", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		expectedBuyer := domain.Buyer{
			CardNumberID: "138935",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		buyerID, err := repository.Save(ctx, expectedBuyer)
		assert.NoError(t, err)

		exists := repository.ExistsID(ctx, buyerID)
		assert.True(t, exists)
	})
	t.Run("Should return false if buyer id not exists", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		exists := repository.ExistsID(ctx, 50)
		assert.False(t, exists)
	})
}

func TestSaveBuyer(t *testing.T) {

	t.Run("Should create a new buyer", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		expectedBuyer := domain.Buyer{
			CardNumberID: "138935",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		_, err := repository.Save(ctx, expectedBuyer)
		assert.NoError(t, err)
	})
}

func TestGetBuyer(t *testing.T) {

	t.Run("Should get buyer by id", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		expectedBuyer := domain.Buyer{
			ID:           9,
			CardNumberID: "2556",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		buyerID, err := repository.Save(ctx, expectedBuyer)
		assert.NoError(t, err)

		buyer, err := repository.Get(ctx, buyerID)

		assert.NoError(t, err)
		assert.NotEmpty(t, buyer)
	})
	t.Run("Should return error if buyer not exists", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		buyer, err := repository.Get(ctx, 0)

		assert.Error(t, err)
		assert.Equal(t, domain.Buyer{}, buyer)
	})
}

func TestGetAllBuyers(t *testing.T) {

	t.Run("Should get all buyers", func(t *testing.T) {
		repository := buyers.NewRepository(db)

		expectedBuyers := domain.Buyer{
			ID:           9,
			CardNumberID: "2556",
			FirstName:    "Giulianna",
			LastName:     "Oliveira",
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		_, err := repository.Save(ctx, expectedBuyers)

		assert.NoError(t, err)

		buyers, err := repository.GetAll(ctx)

		assert.NoError(t, err)
		assert.NotEmpty(t, buyers)
	})
}

func initDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
