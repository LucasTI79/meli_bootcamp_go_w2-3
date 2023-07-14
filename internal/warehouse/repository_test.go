package warehouse_test

import (
	"context"
	"database/sql"

	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/warehouse"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db = InitDatabase()

var warehouseExpected = domain.Warehouse{
	Address:            "Rua Pedro Dias",
	Telephone:          "3712291281",
	MinimumCapacity:    10,
	MinimumTemperature: 10.0,
	LocalityId:         1,
}

func TestGetAllWarehousesRepository(t *testing.T) {
	t.Run("Should get all warehouses in database", func(t *testing.T) {
		warehouseExpected.WarehouseCode = "AAWA"

		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Save(ctx, warehouseExpected)
		assert.NoError(t, err)

		warehouseResult, err := repository.GetAll(ctx)
		assert.NoError(t, err)
		assert.True(t, len(warehouseResult) > 1)
	})
}

func TestSaveWarehousesRepository(t *testing.T) {
	t.Run("should create a warehouse and test", func(t *testing.T) {
		warehouseExpected.WarehouseCode = "AEXA"

		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		result, err := repository.Save(ctx, warehouseExpected)
		assert.NoError(t, err)

		getResult, err := repository.Get(ctx, result)
		assert.NoError(t, err)
		assert.NotNil(t, getResult)
		assert.Equal(t, warehouseExpected.WarehouseCode, getResult.WarehouseCode)
	})
}

func TestExistsRepository(t *testing.T) {
	t.Run("should test if exists a specific warehouseCode", func(t *testing.T) {
		warehouseExpected.WarehouseCode = "AEXD"


		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		_, err := repository.Save(ctx, warehouseExpected)
		assert.NoError(t, err)

		existsResult := repository.Exists(ctx, "AEXD")
		assert.True(t, existsResult)
	})
}

func TestGetWarehousesRepository(t *testing.T) {
	t.Run("Should get the warehouse when it exists in database", func(t *testing.T) {
		warehouseExpected.WarehouseCode = "AAAA"


		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		resultId, err := repository.Save(ctx, warehouseExpected)
		assert.NoError(t, err)

		warehouseResult, err := repository.Get(ctx, resultId)
		assert.NoError(t, err)
		assert.Equal(t, "AAAA", warehouseResult.WarehouseCode)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedMessage := warehouse.ErrNotFound.Error()

		repository := warehouse.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Get(ctx, 200000)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestUpdateWarehousesRepository(t *testing.T) {
	t.Run("should update a warehouse and test", func(t *testing.T) {
		warehouseExpected.WarehouseCode = "AADA"

		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		result, err := repository.Save(ctx, warehouseExpected)
		assert.NoError(t, err)

		warehouseExpected.ID = result
		warehouseExpected.Address = "Rua Maria"

		err = repository.Update(ctx, warehouseExpected)
		assert.NoError(t, err)

		getResult, err := repository.Get(ctx, warehouseExpected.ID)
		assert.NoError(t, err)
		assert.NotNil(t, getResult)
		assert.Equal(t, warehouseExpected.WarehouseCode, getResult.WarehouseCode)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedMessage := warehouse.ErrNotFound.Error()

		warehouseExpected.WarehouseCode = "BASA"

		repository := warehouse.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		warehouseExpected.ID = 200000

		err := repository.Update(ctx, warehouseExpected)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestDeleteWarehousesRepository(t *testing.T) {
	t.Run("should delete a warehouse and test", func(t *testing.T) {
		var warehouseExpectedDelete = domain.Warehouse{
			Address:            "Rua Pedro Dias",
			Telephone:          "3712291281",
			WarehouseCode:      "BASD",
			MinimumCapacity:    10,
			MinimumTemperature: 10.0,
			LocalityId:         1,
		}

		expectedMessage := warehouse.ErrNotFound.Error()
		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		resultId, err := repository.Save(ctx, warehouseExpectedDelete)
		assert.NoError(t, err)

		err = repository.Delete(ctx, resultId)
		assert.NoError(t, err)

		_, err = repository.Get(ctx, warehouseExpected.ID)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedMessage := warehouse.ErrNotFound.Error()

		repository := warehouse.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Delete(ctx, 2000000)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestAllEndpointsRepositoryWithErrorDatabaseClosed(t *testing.T) {
	db.Close()
	warehouseExpected.WarehouseCode = "FAAS"
	t.Run("Should return error when there is an GetAll database error", func(t *testing.T) {
		repository := warehouse.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.GetAll(ctx)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Get database error", func(t *testing.T) {
		repository := warehouse.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Get(ctx, 200)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Save database error", func(t *testing.T) {
		repository := warehouse.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Save(ctx, warehouseExpected)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Update database error", func(t *testing.T) {
		repository := warehouse.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Update(ctx, warehouseExpected)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Delete database error", func(t *testing.T) {
		repository := warehouse.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Delete(ctx, warehouseExpected.ID)
		assert.Error(t, err)
	})
}

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
