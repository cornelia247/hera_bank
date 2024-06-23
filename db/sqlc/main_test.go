package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecret@localhost:5432/hera_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}

// func TestCreateAccount(t *testing.T) {
//     arg := CreateAccountParams{
//         Owner:    "Alice",
//         Balance:  1000,
//         Currency: "USD",
//     }

//     account, err := testQueries.CreateAccount(context.Background(), arg)
//     if err != nil {
//         t.Fatal(err)
//     }

//     if account.Owner != arg.Owner {
//         t.Errorf("expected owner %v; got %v", arg.Owner, account.Owner)
//     }
//     // Add more assertions as needed
// }
