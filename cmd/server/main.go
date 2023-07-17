package main

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// NO MODIFICAR
	db, err := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/melisprint")
	if err != nil {
		panic(err)
	}

	eng := gin.Default()

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}
