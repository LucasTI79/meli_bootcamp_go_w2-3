package locality_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/seller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db = InitDatabase()

var localityExpected = domain.Locality{
    ID: 1,
	LocalityName: "Florianopolis",
	ProvinceName: "Santa Catarina",
}

var localityInputExpected = domain.LocalityInput{
	LocalityName: "Florianopolis",
	IdProvince:   1,
}

func TestCreateLocality(t *testing.T) {
	t.Run("should create a locality and check if it exists", func(t *testing.T) {
		repository := locality.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		id, err := repository.Save(ctx, localityInputExpected)
		assert.NoError(t, err)

		exists := repository.ExistsById(ctx, id)
		assert.True(t, exists)
	})
}

func TestGetProvinceByName(t *testing.T) {
	t.Run("should search for a province by name", func(t *testing.T) {
		repository := locality.NewRepository(db)
        IdProvince :=   1

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
        
		id, err := repository.GetProvinceByName(ctx, localityExpected.ProvinceName)
		assert.NoError(t, err)
		assert.Equal(t, IdProvince, id)
	})
	t.Run("Should return error when not finding the province", func(t *testing.T) {
		repository := locality.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
        
		_, err := repository.GetProvinceByName(ctx, "aaaa")
		assert.Error(t, err)
	})
}

func TestExistsByIdRepository(t *testing.T) {
	t.Run("should return a true value if the id passed exists", func(t *testing.T) {
        repository := locality.NewRepository(db)
        ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		IdLocality, err := repository.Save(ctx, localityInputExpected)
		assert.NoError(t, err)

        exists := repository.ExistsById(ctx, IdLocality)
        assert.True(t, exists)
    })
}

func TestReportLocalityId(t *testing.T) {
	t.Run("should return a location id report with the sum of sellers", func(t *testing.T) {
        repository := locality.NewRepository(db)
        repositorySeller := seller.NewRepository(db)
		sellerExpected := domain.Seller{
			CID: 11,
			CompanyName: "Mercado Livre",
			Address: "Rua Feliz",
			Telephone: "123456",
			LocalityId: 1,
		}

        ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
        
        IdLocality, err := repository.Save(ctx, localityInputExpected)
		assert.NoError(t, err)

		sellerExpected.LocalityId = IdLocality

        _, errs := repositorySeller.Save(ctx, sellerExpected)
		assert.NoError(t, errs)

        report, err := repository.ReportLocalityId(ctx, IdLocality)
        assert.NoError(t, err)
		assert.Equal(t, IdLocality, report.IdLocality)
		assert.True(t, report.SellersCount == 1)
    })
	t.Run("Should return error when locality id is not exists in database", func(t *testing.T) {
		repository := locality.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		_, err := repository.ReportLocalityId(ctx, 20000000)
		assert.Error(t, err)
	})
}

func TestReportLocality(t *testing.T) {
	t.Run("should return a report of all locations with the number of sellers", func(t *testing.T) {
		repository := locality.NewRepository(db)
        repositorySeller := seller.NewRepository(db)

		sellerExpected := domain.Seller{
			CID: 12,
			CompanyName: "Mercado Livre",
			Address: "Rua Feliz",
			Telephone: "123456",
			LocalityId: 1,
		}

        ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
        
        IdLocality, err := repository.Save(ctx, localityInputExpected)
		assert.NoError(t, err)

		localityInputExpected.ID = IdLocality

        _, errs := repositorySeller.Save(ctx, sellerExpected)
		assert.NoError(t, errs)

        sellerExpected.LocalityId = IdLocality

        report, err := repository.ReportLocality(ctx)

        assert.NoError(t, err)
		assert.True(t, len(report) > 0)
    })
}

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
