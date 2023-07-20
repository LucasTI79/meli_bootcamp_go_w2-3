package inbound_order_test

import (
	"context"
	"database/sql"

	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/inbound_order"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db = InitDatabase()

var InboundOrdersExpected = domain.InboundOrders{
	OrderDate:      "01/01/01",
	OrderNumber:    "001",
	EmployeeID:     1,
	ProductBatchID: 1,
	WarehouseID:    1,
}

func TestCreateInboundOrdersRepository(t *testing.T) {
	t.Run("should create a inbound order and test", func(t *testing.T) {
		repository := inbound_order.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		result, err := repository.Create(ctx, InboundOrdersExpected)
		assert.NoError(t, err)

		InboundOrdersExpected.ID = result

		getResult, err := repository.Get(ctx, result)
		assert.NoError(t, err)
		assert.NotNil(t, getResult)
		assert.Equal(t, InboundOrdersExpected.ID, getResult.ID)
	})
}

func TestExistsByIdInboundOrderRepository(t *testing.T) {
	t.Run("should test if exists a specific id", func(t *testing.T) {
		repository := inbound_order.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		result, err := repository.Create(ctx, InboundOrdersExpected)
		assert.NoError(t, err)

		getResult, _ := repository.Get(ctx, result)

		existsResult := repository.Exists(ctx, getResult.OrderNumber)
		assert.True(t, existsResult)
	})
}

func TestGetInboundOrderRepository(t *testing.T) {
	t.Run("Should get the inbound order when it exists in database", func(t *testing.T) {
		repository := inbound_order.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		result, err := repository.Create(ctx, InboundOrdersExpected)
		assert.NoError(t, err)

		InboundOrdersExpected.ID = result

		inboundOrderResult, err := repository.Get(ctx, result)

		assert.NoError(t, err)
		assert.Equal(t, InboundOrdersExpected.ID, inboundOrderResult.ID)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {

		repository := inbound_order.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Get(ctx, 200000)
		assert.Error(t, err)
	})
}

func TestGetAllInboundOrders(t *testing.T) {
	t.Run("Should get inbound order reports of all", func(t *testing.T) {

		repository := inbound_order.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		report, err := repository.ReportByAll(ctx)
		assert.NoError(t, err)
		assert.True(t, len(report) > 1)
	})
}

func TestGetOneInboundOrders(t *testing.T) {
	t.Run("It should get the report of a employee record by the inbound order id.", func(t *testing.T) {
		id := 1
		repository := inbound_order.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		report, err := repository.ReportByOne(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, "card 1", report.CardNumberID)
	})
	t.Run("It should return an error when the record of a inbound order is not found.", func(t *testing.T) {
		repository := inbound_order.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.ReportByOne(ctx, 50000000)
		assert.Error(t, err)
		expectedErrorMessage := inbound_order.ErrNotFound.Error()
		assert.Equal(t, expectedErrorMessage, err.Error())
	})
}

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
