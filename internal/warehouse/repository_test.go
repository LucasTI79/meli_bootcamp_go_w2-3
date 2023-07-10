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
	WarehouseCode:      "AEXA",
	MinimumCapacity:    10,
	MinimumTemperature: 10.0,
	LocalityId:         1,
}

var expectedWarehouses = []domain.Warehouse{
	{
		ID:                 1,
		Address:            "Rua Pedro Dias",
		Telephone:          "3712291281",
		WarehouseCode:      "DAE",
		MinimumCapacity:    10,
		MinimumTemperature: 10.0,
		LocalityId:         1,
	},
	{
		ID:                 2,
		Address:            "Rua Maria das Dores",
		Telephone:          "1722919394",
		WarehouseCode:      "EWQ",
		MinimumCapacity:    10,
		MinimumTemperature: 10.0,
		LocalityId:         1,
	},
}

func TestGetAllWarehousesRepository(t *testing.T) {
	t.Run("Should get all warehouses in database", func(t *testing.T) {
		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		warehouseResult, err := repository.GetAll(ctx)
		assert.NoError(t, err)
		assert.True(t, len(warehouseResult) > 1)
	})
}

func TestSaveWarehousesRepository(t *testing.T) {
	t.Run("should create a warehouse and test", func(t *testing.T) {
		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		warehouseExpected.WarehouseCode = "AS12"

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
		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		existsResult := repository.Exists(ctx, "AEXA")
		assert.True(t, existsResult)
	})
}

func TestGetWarehousesRepository(t *testing.T) {
	t.Run("Should get the warehouse when it exists in database", func(t *testing.T) {
		id := 5

		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		warehouseResult, err := repository.Get(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, "AEXA", warehouseResult.WarehouseCode)
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
		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		err := repository.Update(ctx, expectedWarehouses[0])
		assert.NoError(t, err)

		getResult, err := repository.Get(ctx, expectedWarehouses[0].ID)
		assert.NoError(t, err)
		assert.NotNil(t, getResult)
		assert.Equal(t, expectedWarehouses[0].WarehouseCode, getResult.WarehouseCode)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedMessage := warehouse.ErrNotFound.Error()

		repository := warehouse.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		expectedWarehouses[0].ID = 200000

		err := repository.Update(ctx, expectedWarehouses[0])
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestDeleteWarehousesRepository(t *testing.T) {
	t.Run("should delete a warehouse and test", func(t *testing.T) {
		expectedMessage := warehouse.ErrNotFound.Error()
		repository := warehouse.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		warehouseExpected.ID = 5

		err := repository.Delete(ctx, warehouseExpected.ID)
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

		err := repository.Update(ctx, expectedWarehouses[0])
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Delete database error", func(t *testing.T) {
		repository := warehouse.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Delete(ctx, expectedWarehouses[0].ID)
		assert.Error(t, err)
	})
}

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
