package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)

	// r.Post("/", controllers.CreateNewPet)

	// http.ListenAndServe(":3000", r)

	db, _ := sql.Open("mysql", "maria:123@tcp(localhost:3306)/petdex?multiStatements=true")
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	m.Up()
}
