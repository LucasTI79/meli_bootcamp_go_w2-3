package carry_test

import (
	"context"
	"database/sql"

	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/carry"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/locality"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db = InitDatabase()

var carryExpected = domain.Carry{
	Cid:         "111111",
	CompanyName: "Carrier 1",
	Address:     "Carrier Address 1",
	Telephone:   "111111111",
	LocalityId:  1,
}

var localityExpected = domain.LocalityInput{
	LocalityName: "Belo Horizonte",
	IdProvince:   1,
}

func TestCreateCarriersRepository(t *testing.T) {
	t.Run("should create a carry and test", func(t *testing.T) {
		repository := carry.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		result, err := repository.Create(ctx, carryExpected)
		assert.NoError(t, err)

		getResult, err := repository.Get(ctx, result)
		assert.NoError(t, err)
		assert.NotNil(t, getResult)
		assert.Equal(t, carryExpected.Cid, getResult.Cid)
	})
}

func TestExistsByCidCarryRepository(t *testing.T) {
	t.Run("should test if exists a specific cid", func(t *testing.T) {
		repository := carry.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		result, err := repository.Create(ctx, carryExpected)
		assert.NoError(t, err)

		getResult, _ := repository.Get(ctx, result)

		existsResult := repository.ExistsByCidCarry(ctx, getResult.Cid)
		assert.True(t, existsResult)
	})
}

func TestGetCarriersRepository(t *testing.T) {
	t.Run("Should get the carry when it exists in database", func(t *testing.T) {
		id := 1

		repository := carry.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		carryResult, err := repository.Get(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, carryExpected.Cid, carryResult.Cid)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedMessage := carry.ErrNotFound.Error()

		repository := carry.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Get(ctx, 200000)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestReadCarriersWithLocalityIdRepository(t *testing.T) {
	t.Run("Should get the locality when it exists in database when id is passed", func(t *testing.T) {
		repository := carry.NewRepository(db)
		repositoryLocality := locality.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		resultCreateLocality, err := repositoryLocality.Save(ctx, localityExpected)
		assert.NoError(t, err)

		carryExpected.LocalityId = resultCreateLocality

		_, err = repository.Create(ctx, carryExpected)
		assert.NoError(t, err)

		result, err := repository.ReadCarriersWithLocalityId(ctx, resultCreateLocality)
		assert.NoError(t, err)
		assert.Equal(t, resultCreateLocality, result.LocalityID)
		assert.True(t, result.CarriersCount == 1)
	})
	t.Run("Should return error when locality id is not exists in database", func(t *testing.T) {
		expectedMessage := carry.ErrNotFound.Error()

		repository := carry.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.ReadCarriersWithLocalityId(ctx, 20000000)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestReadAllCarriersRepository(t *testing.T) {
	t.Run("Should get the locality when it exists in database when id is not passed", func(t *testing.T) {
		repository := carry.NewRepository(db)
		repositoryLocality := locality.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		resultCreateLocality, err := repositoryLocality.Save(ctx, localityExpected)
		assert.NoError(t, err)

		carryExpected.LocalityId = resultCreateLocality

		_, err = repository.Create(ctx, carryExpected)
		assert.NoError(t, err)

		result, err := repository.ReadAllCarriers(ctx)
		assert.NoError(t, err)
		assert.True(t, len(result) > 0)
	})
}

func TestAllEndpointsRepositoryWithErrorDatabaseClosed(t *testing.T) {
	db.Close()
	t.Run("Should return error when there is an ReadAllCarriers database error", func(t *testing.T) {
		repository := carry.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.ReadAllCarriers(ctx)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Get database error", func(t *testing.T) {
		repository := carry.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Get(ctx, 200)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an ReadCarriersWithLocalityId database error", func(t *testing.T) {
		repository := carry.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.ReadCarriersWithLocalityId(ctx, 20000000)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Create database error", func(t *testing.T) {
		repository := carry.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Create(ctx, carryExpected)
		assert.Error(t, err)
	})
}

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
