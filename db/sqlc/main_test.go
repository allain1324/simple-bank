package db

import (
	"database/sql"

	"log"
	"os"
	"testing"

	"github.com/allain1324/simplebank/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Can't load config", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Failed to open database", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
