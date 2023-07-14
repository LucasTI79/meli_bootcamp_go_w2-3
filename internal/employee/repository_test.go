package employee_test

import (
	"context"
	"database/sql"

	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/employee"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db = InitDatabase()

func TestGetAllEmployeesRepository(t *testing.T) {
	t.Run("Should get all employees in database", func(t *testing.T) {
		var employeeExpected = domain.Employee{
			ID:           01,
			CardNumberID: "001",
			FirstName:    "Joana",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		repository := employee.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Save(ctx, employeeExpected)
		assert.NoError(t, err)

		employeeResult, err := repository.GetAll(ctx)
		assert.NoError(t, err)
		assert.True(t, len(employeeResult) > 1)
	})
}

func TestSaveEmployeesRepository(t *testing.T) {
	t.Run("should create a employees and test", func(t *testing.T) {
		var employeeExpected = domain.Employee{
			ID:           01,
			CardNumberID: "001",
			FirstName:    "Joana",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		repository := employee.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		result, err := repository.Save(ctx, employeeExpected)
		assert.NoError(t, err)

		getResult, err := repository.Get(ctx, result)
		assert.NoError(t, err)
		assert.NotNil(t, getResult)
		assert.Equal(t, employeeExpected.ID, getResult.ID)
	})
}

func TestExistsRepository(t *testing.T) {
	t.Run("should test if exists a specific card number ID", func(t *testing.T) {
		var employeeExpected = domain.Employee{
			ID:           01,
			CardNumberID: "001",
			FirstName:    "Joana",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		repository := employee.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		_, err := repository.Save(ctx, employeeExpected)
		assert.NoError(t, err)

		existsResult := repository.Exists(ctx, "001")
		assert.True(t, existsResult)
	})
}

func TestGetEmployeesRepository(t *testing.T) {
	t.Run("Should get the employee when it exists in database", func(t *testing.T) {
		var employeeExpected = domain.Employee{
			ID:           01,
			CardNumberID: "001",
			FirstName:    "Joana",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		repository := employee.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		resultId, err := repository.Save(ctx, employeeExpected)
		assert.NoError(t, err)

		employeeResult, err := repository.Get(ctx, resultId)
		assert.NoError(t, err)
		assert.Equal(t, 01, employeeResult.ID)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedMessage := employee.ErrNotFound.Error()

		repository := employee.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Get(ctx, 200000)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestUpdateEmployeesRepository(t *testing.T) {
	t.Run("should update a employee and test", func(t *testing.T) {
		var employeeExpected = domain.Employee{
			ID:           01,
			CardNumberID: "001",
			FirstName:    "Joana",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		repository := employee.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		result, err := repository.Save(ctx, employeeExpected)
		assert.NoError(t, err)

		employeeExpected.ID = result
		employeeExpected.FirstName = "Luciana"

		err = repository.Update(ctx, employeeExpected)
		assert.NoError(t, err)

		getResult, err := repository.Get(ctx, employeeExpected.ID)
		assert.NoError(t, err)
		assert.NotNil(t, getResult)
		assert.Equal(t, employeeExpected.ID, getResult.ID)
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedMessage := employee.ErrNotFound.Error()

		var employeeExpected = domain.Employee{
			ID:           01,
			CardNumberID: "001",
			FirstName:    "Joana",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		repository := employee.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		employeeExpected.ID = 200000

		err := repository.Update(ctx, employeeExpected)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestDeleteEmployeesRepository(t *testing.T) {
	t.Run("should delete a employee and test", func(t *testing.T) {
		var employeeExpected = domain.Employee{
			ID:           01,
			CardNumberID: "001",
			FirstName:    "Joana",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		expectedMessage := employee.ErrNotFound.Error()
		repository := employee.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		resultId, err := repository.Save(ctx, employeeExpected)
		assert.NoError(t, err)

		err = repository.Delete(ctx, resultId)
		assert.NoError(t, err)

		_, err = repository.Get(ctx, employeeExpected.ID)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
	t.Run("Should return error when there is not exists in database", func(t *testing.T) {
		expectedMessage := employee.ErrNotFound.Error()

		repository := employee.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Delete(ctx, 2000000)
		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})
}

func TestAllEndpointsRepositoryWithErrorDatabaseClosed(t *testing.T) {
	db.Close()
	var employeeExpected = domain.Employee{
		ID:           01,
		CardNumberID: "001",
		FirstName:    "Joana",
		LastName:     "Silva",
		WarehouseID:  1,
	}
	t.Run("Should return error when there is an GetAll database error", func(t *testing.T) {
		repository := employee.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.GetAll(ctx)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Get database error", func(t *testing.T) {
		repository := employee.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Get(ctx, 200)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Save database error", func(t *testing.T) {
		repository := employee.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.Save(ctx, employeeExpected)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Update database error", func(t *testing.T) {
		repository := employee.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Update(ctx, employeeExpected)
		assert.Error(t, err)
	})
	t.Run("Should return error when there is an Delete database error", func(t *testing.T) {
		repository := employee.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		err := repository.Delete(ctx, employeeExpected.ID)
		assert.Error(t, err)
	})
}

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
