package productrecord

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
)

expectedReport := domain.ProductRecordReport{
	ProductID:    1,
	Description:  "Product 1",
	RecordsCount: 3,
}

expectedReports := []domain.ProductRecordReport{
	{
		ProductID:    1,
		Description:  "Product 1",
		RecordsCount: 3,
	},
	{
		ProductID:    2,
		Description:  "Product 2",
		RecordsCount: 3,
	},
}

func InitDatabase() *sql.DB {
	txdb.Register("txdb", "mysql", "root:@/melisprint")
	db, _ := sql.Open("txdb", uuid.New().String())
	return db
}
