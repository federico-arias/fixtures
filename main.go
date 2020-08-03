package main

import (
	"database/sql"
	"flag"
	"github.com/go-testfixtures/testfixtures/v3"

	_ "github.com/lib/pq"
)

var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

var fixturesFolder = flag.String(
	"fixtures", "", "fixtures folder")
var conn = flag.String(
	"db", "", "db conn string")

func main() {
	flag.Parse()
	db, err := sql.Open("postgres", *conn)
	if err != nil {
		panic(err.Error())
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),        // You database connection
		testfixtures.Dialect("postgres"), // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.UseAlterConstraint(),
		testfixtures.Directory(*fixturesFolder), // the directory containing the YAML files
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.ResetSequencesTo(20000),
	)

	if err := fixtures.Load(); err != nil {
		panic(err.Error())
	}
}
