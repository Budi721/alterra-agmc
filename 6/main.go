package main

import (
	"flag"
	"github.com/Budi721/alterra-agmc/v6/database"
	"github.com/Budi721/alterra-agmc/v6/database/migration"
	"github.com/Budi721/alterra-agmc/v6/database/seeder"
	"github.com/Budi721/alterra-agmc/v6/internal/factory"
	"github.com/Budi721/alterra-agmc/v6/internal/http"
	"github.com/Budi721/alterra-agmc/v6/internal/middleware"
	"github.com/labstack/echo/v4"
)

func init() {
	//if err := godotenv.Load(); err != nil {
	//	panic(err)
	//}
	database.GetConnection()
}

func main() {
	database.CreateConnection()

	var m string // for check migration
	var s string // for check seeder

	flag.StringVar(
		&m,
		"m",
		"none",
		`this argument for check if user want to migrate table, rollback table, or status migration
to use this flag:
	use -m=migrate for migrate table
	use -m=rollback for rollback table
	use -m=status for get status migration`,
	)

	flag.StringVar(
		&s,
		"s",
		"none",
		`this argument for check if user want to seed table
to use this flag:
	use -s=all to seed all table`,
	)

	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
	} else if m == "rollback" {
		migration.Rollback()
	} else if m == "status" {
		migration.Status()
	}

	if s == "all" {
		seeder.NewSeeder().DeleteAll()
		seeder.NewSeeder().SeedAll()
	}

	f := factory.NewFactory()
	e := echo.New()

	middleware.LogMiddlewares(e)

	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":8080"))
}
