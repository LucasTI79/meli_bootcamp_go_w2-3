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

var db = InitDatabase()

var sellerExpected = domain.Seller{
	CID:         99,
	CompanyName: "Mercado Livre",
	Address:     "Rua Feliz",
	Telephone:   "123456",
	LocalityId:  1,
}


func TestGetAll(t *testing.T) {
	t.Run("should look for all sellers", func(t *testing.T) {
		repository := seller.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		sellersCreated := []domain.Seller{
			{
				CID:         11,
				CompanyName: "Mercado Livre",
				Address:     "Rua Feliz",
				Telephone:   "123456",
				LocalityId:  1,
			},
			{
				CID:         12,
				CompanyName: "Mercado Livre",
				Address:     "Rua Feliz",
				Telephone:   "123456",
				LocalityId:  1,
			},
		}

		_, err1 := repository.Save(ctx, sellersCreated[0])
		assert.NoError(t, err1)

		_, err2 := repository.Save(ctx, sellersCreated[1])
		assert.NoError(t, err2)

		SellersResult, err := repository.GetAll(ctx)
		assert.NoError(t, err)
		assert.True(t, len(SellersResult) > 1)
	})
}

func TestGet(t *testing.T) {
	t.Run("should look for the seller corresponding to the last id", func(t *testing.T) {
		repository := seller.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		sellerCreated := domain.Seller{
			CID:         13,
			CompanyName: "Mercado Livre",
			Address:     "Rua Feliz",
			Telephone:   "123456",
			LocalityId:  1,
		}

		sellerId, err1 := repository.Save(ctx, sellerCreated)
		assert.NoError(t, err1)

		sellersResult, err := repository.Get(ctx, sellerId)
		assert.NoError(t, err)
		assert.Equal(t, sellerId, sellersResult.ID)
	})
}

func TestExists(t *testing.T) {
	t.Run("should return value true if the CID exists", func(t *testing.T) {
		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		sellerCreated := domain.Seller{
			CID:         14,
			CompanyName: "Mercado Livre",
			Address:     "Rua Feliz",
			Telephone:   "123456",
			LocalityId:  1,
		}
		_, err1 := repository.Save(ctx, sellerCreated)
		assert.NoError(t, err1)

		exists := repository.Exists(ctx, sellerCreated.CID)
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

		sellerCreated := domain.Seller{
			CID:         15,
			CompanyName: "Mercado Livre",
			Address:     "Rua Feliz",
			Telephone:   "123456",
			LocalityId:  1,
		}

		sellerId, err1 := repository.Save(ctx, sellerCreated)
		assert.NoError(t, err1)

		sellerUpdate := domain.Seller{
			ID:          sellerId,
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

		sellerCreated := domain.Seller{
			CID:         33,
			CompanyName: "Mercado Livre",
			Address:     "Rua Feliz",
			Telephone:   "123456",
			LocalityId:  1,
		}

		sellerId, err1 := repository.Save(ctx, sellerCreated)
		assert.NoError(t, err1)

		err := repository.Delete(ctx, sellerId)
		assert.NoError(t, err)

		_, err = repository.Get(ctx, sellerId)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedMessage := seller.ErrNotFound.Error()

		repository := seller.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Delete(ctx, 12357654432)
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
