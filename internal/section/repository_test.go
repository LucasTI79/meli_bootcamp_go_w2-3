package section_test

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/section"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	db              = initDatabase()
	sectionExpected = domain.Section{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseID:        1,
		ProductTypeID:      1,
	}
)

func TestCreateSectionRepository(t *testing.T) {

	t.Run("Should create a new section", func(t *testing.T) {
		repository := section.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		_, err := repository.Save(ctx, sectionExpected)
		assert.NoError(t, err)
	})
	t.Run("Should not create new section", func(t *testing.T) {
		repository := section.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		_, err := repository.Save(ctx, domain.Section{})
		assert.Error(t, err)
	})
}

func TestGetAllSectionRepository(t *testing.T) {

	t.Run("Should get all sections", func(t *testing.T) {
		repository := section.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		_, err := repository.Save(ctx, sectionExpected)

		assert.NoError(t, err)

		sections, err := repository.GetAll(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, sections)
	})
}

func TestGetSectionRepository(t *testing.T) {

	t.Run("Should get section by id", func(t *testing.T) {
		repository := section.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		newSectionId, err := repository.Save(ctx, sectionExpected)
		assert.NoError(t, err)

		section, err := repository.Get(ctx, newSectionId)
		assert.NoError(t, err)
		assert.NotEmpty(t, section)
	})
	t.Run("Should return error if section does not exists", func(t *testing.T) {
		repository := section.NewRepository(db)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		section, err := repository.Get(ctx, 0)
		assert.Error(t, err)
		assert.Equal(t, domain.Section{}, section)
	})
}

func TestExistsSectionRepository(t *testing.T) {

	// t.Run("Should return true if section exists", func(t *testing.T) {
	// 	repository := section.NewRepository(db)
	// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// 	defer cancel()
	// 	newSectionId, err := repository.Save(ctx, sectionExpected)
	// 	assert.NoError(t, err)

	// 	exists := repository.Exists(ctx, newSectionId)
	// 	assert.True(t, exists)
	// })
	t.Run("Should return false if section does not exists", func(t *testing.T) {
		repository := section.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		exists := repository.Exists(ctx, 0)
		assert.False(t, exists)
	})
}

func TestUpdateSectionRepository(t *testing.T) {

	t.Run("Should update section", func(t *testing.T) {
		repository := section.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		newSectionId, err := repository.Save(ctx, sectionExpected)
		assert.NoError(t, err)

		sectionExpected.SectionNumber = 2
		sectionExpected.CurrentTemperature = 2
		sectionExpected.MinimumTemperature = 2
		sectionExpected.CurrentCapacity = 2
		sectionExpected.MinimumCapacity = 2
		sectionExpected.MaximumCapacity = 2
		sectionExpected.WarehouseID = 2
		sectionExpected.ProductTypeID = 2

		sectionExpected.ID = newSectionId
		err = repository.Update(ctx, sectionExpected)
		assert.NoError(t, err)

		section, err := repository.Get(ctx, newSectionId)
		assert.NoError(t, err)
		assert.Equal(t, sectionExpected, section)
	})
	t.Run("Should return error if section does not exists", func(t *testing.T) {
		repository := section.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		err := repository.Update(ctx, domain.Section{})
		assert.Error(t, err)
	})
}
func TestDeleteSectionRepository(t *testing.T) {

	t.Run("Should delete section", func(t *testing.T) {
		repository := section.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		newSectionId, err := repository.Save(ctx, sectionExpected)
		assert.NoError(t, err)

		err = repository.Delete(ctx, newSectionId)
		assert.NoError(t, err)

		section, err := repository.Get(ctx, newSectionId)
		assert.Error(t, err)
		assert.Equal(t, domain.Section{}, section)
	})
	t.Run("Should return error if section does not exists", func(t *testing.T) {
		repository := section.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		err := repository.Delete(ctx, 0)
		assert.Error(t, err)
	})
}

func TestExistsByIdSectionRepository(t *testing.T) {
	t.Run("Should return true if section exists", func(t *testing.T) {
		repository := section.NewRepository(db)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		newSectionId, err := repository.Save(ctx, sectionExpected)
		assert.NoError(t, err)

		exists := repository.ExistsById(newSectionId)
		assert.True(t, exists)
	})
	t.Run("Should return false if section does not exists", func(t *testing.T) {
		repository := section.NewRepository(db)
		exists := repository.ExistsById(0)
		assert.False(t, exists)
	})
}

func TestSectionProductsReportsBySectionRepository(t *testing.T) {
	// t.Run("Should return products by section", func(t *testing.T) {
	// 	repository := section.NewRepository(db)
	// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// 	defer cancel()
	// 	newSectionId, err := repository.Save(ctx, sectionExpected)
	// 	assert.NoError(t, err)

	// 	section, err := repository.SectionProductsReportsBySection(newSectionId)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, domain.ProductBySection{}, section)
	// })
	// t.Run("Should return error if section does not exists", func(t *testing.T) {
	// 	repository := section.NewRepository(db)
	// 	section, err := repository.SectionProductsReportsBySection(0)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, domain.ProductBySection{}, section)
	// })

}

func initDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	id := fmt.Sprintf("txdb%f", rand.Float64()*5)
	db, _ := sql.Open("txdb", id)
	return db
}
