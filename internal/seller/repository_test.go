package seller_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/seller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var sellerExpected = domain.Seller{
	CID:         11,
	CompanyName: "Mercado Livre",
	Address:     "Rua Feliz",
	Telephone:   "123456",
	LocalityId:  1,
}

var db = InitDatabase()

func TestGetAll(t *testing.T) {
	t.Run("should look for all sellers", func(t *testing.T) {
		repository := seller.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		SellersResult, err := repository.GetAll(ctx)
		assert.NoError(t, err)
		assert.True(t, len(SellersResult) > 1)
	})
}

func TestGet(t *testing.T) {
	t.Run("should look for the seller corresponding to the last id", func(t *testing.T) {
		id := 6
		repository := seller.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		sellersResult, err := repository.Get(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, 6, sellersResult.ID)
	})
}

func TestExists(t *testing.T) {
	t.Run("should return value true if the CID exists", func(t *testing.T) {
		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		cid := 1

		exists := repository.Exists(ctx, cid)
		assert.True(t, exists)
	})
}

func TestSave(t *testing.T) {
	t.Run("should create a seller and check if it exists", func(t *testing.T) {
		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		idSeller, err := repository.Save(ctx, sellerExpected)
		assert.NoError(t, err)

		sellersResult, err := repository.Get(ctx, idSeller)
		assert.NoError(t, err)
		assert.NotNil(t, sellersResult)
		assert.Equal(t, idSeller, sellersResult.ID)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should update a seller and verify", func(t *testing.T) {
		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		sellerUpdate := domain.Seller{
			ID:          6,
			CID:         22,
			CompanyName: "Mercado Livre",
			Address:     "Rua Feliz",
			Telephone:   "123456",
			LocalityId:  1,
		}

		err := repository.Update(ctx, sellerUpdate)
		assert.NoError(t, err)

		sellersResult, err := repository.Get(ctx, sellerUpdate.ID)
		fmt.Print(sellersResult)

		assert.NoError(t, err)
		assert.NotNil(t, sellersResult)
		assert.Equal(t, sellerUpdate.ID, sellersResult.ID)

	})
}

func TestDelete(t *testing.T) {
	t.Run("should delete a seller and verify", func(t *testing.T) {
		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		expectedMessage := seller.ErrNotFound.Error()
		sellerExpected.ID = 6

		err := repository.Delete(ctx, sellerExpected.ID)
		assert.NoError(t, err)

		_, err = repository.Get(ctx, sellerExpected.ID)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedMessage := seller.ErrNotFound.Error()

		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Delete(ctx, 0001)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestAllEndpointsRepositoryWithErrorDatabaseClosed(t *testing.T) {
	db.Close()
	t.Run("Should return error when there is an GetAll database error", func(t *testing.T) {
		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.GetAll(ctx)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Get database error", func(t *testing.T) {
		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Get(ctx, 200)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Save database error", func(t *testing.T) {
		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Save(ctx, sellerExpected)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Update database error", func(t *testing.T) {
		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Update(ctx, sellerExpected)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Delete database error", func(t *testing.T) {
		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Delete(ctx, sellerExpected.ID)
		assert.Error(t, err)
	})
}

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
