package money_test

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/radicalcompany/money"
	"github.com/stretchr/testify/assert"
)

func TestMoneyDriver(t *testing.T) {
	f := tempFilename(t)

	db := initDB(f, t)
	// close the database at the end of the function
	defer dropDB(f, t, db)

	costAlwaysInEurStoredAsInt64 := money.MustForge(1000, "EUR")
	payedByCustomerCanBeAnyCurrenciesStoredAsString := money.MustForge(1132, "USD")
	payedAlwaysInGBPButStoredAsFloat := money.MustForge(123, "GBP")

	table := "tshirt_catalog_" + t.Name()
	n := "radical-tshirt-" + uuid.New().String()
	//new Customer
	c := &tshirt{
		n,
		costAlwaysInEurStoredAsInt64,
		payedByCustomerCanBeAnyCurrenciesStoredAsString,
		payedAlwaysInGBPButStoredAsFloat,
	}
	// insert customer!
	_, err := db.Exec(
		"insert into "+table+" "+
			"(name, costAlwaysInEurStoredAsInt64, payedByCustomerCanBeAnyCurrenciesStoredAsString, payedAlwaysInGBPButStoredAsFloat) "+
			"values (?, ?, ?, ?);",
		c.name,
		c.costAlwaysInEurStoredAsInt64.Int64(),
		c.payedByCustomerCanBeAnyCurrenciesStoredAsString.String(),
		c.payedAlwaysInGBPButStoredAsFloat.Float(),
	)
	assert.Nil(t, err, err)

	// , costAlwaysInEur, payedByCustomerCanBeAnyCurrencies
	rows, err := db.Query("SELECT name, costAlwaysInEurStoredAsInt64, payedByCustomerCanBeAnyCurrenciesStoredAsString, payedAlwaysInGBPButStoredAsFloat FROM "+table+" WHERE name like ?", n)
	assert.Nil(t, err)
	defer rows.Close()
	tss := make([]tshirt, 0)

	for rows.Next() {
		ts := tshirt{}
		ts.costAlwaysInEurStoredAsInt64 = money.EUR(0)
		ts.payedByCustomerCanBeAnyCurrenciesStoredAsString = money.JPY(0)
		ts.payedAlwaysInGBPButStoredAsFloat = money.GBP(0)

		e := rows.Scan(&ts.name,
			&ts.costAlwaysInEurStoredAsInt64,
			&ts.payedByCustomerCanBeAnyCurrenciesStoredAsString,
			&ts.payedAlwaysInGBPButStoredAsFloat)
		assert.Nil(t, e)
		tss = append(tss, ts)
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	assert.Nil(t, rerr)

	// Rows.Err will report the last error encountered by Rows.Scan.
	if e := rows.Err(); e != nil {
		assert.Nil(t, e)
	}

	_, err = db.Exec(
		"delete from "+table+" where name like (?);",
		n,
	)
	assert.Nil(t, err, err)

	assert.Equal(t, tss[0].costAlwaysInEurStoredAsInt64, costAlwaysInEurStoredAsInt64)

	assert.Equal(t, tss[0].payedByCustomerCanBeAnyCurrenciesStoredAsString, payedByCustomerCanBeAnyCurrenciesStoredAsString)

	assert.Equal(t, tss[0].payedAlwaysInGBPButStoredAsFloat, payedAlwaysInGBPButStoredAsFloat)
}

type tshirt struct {
	name                                            string
	costAlwaysInEurStoredAsInt64                    money.Money
	payedByCustomerCanBeAnyCurrenciesStoredAsString money.Money
	payedAlwaysInGBPButStoredAsFloat                money.Money
}

func tempFilename(t *testing.T) string {
	f, err := ioutil.TempFile("", fmt.Sprintf("go_sqlite_test_%s.db", t.Name()))
	if err != nil {
		t.Fatal(err)
	}
	_ = f.Close()
	return f.Name()
}

func initDB(dbname string, t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", dbname)
	assert.Nil(t, err)

	q := "tshirt_catalog_" + t.Name()

	qr := "CREATE TABLE if not exists `" + q + "` (" +
		"`name` varchar(255) NOT NULL," +
		"`costAlwaysInEurStoredAsInt64` int(20) DEFAULT NULL," +
		"`payedByCustomerCanBeAnyCurrenciesStoredAsString` varchar(255) NOT NULL," +
		"`payedAlwaysInGBPButStoredAsFloat` decimal(13,4) NOT NULL," +
		"PRIMARY KEY (`name`)" +
		");"

	_, err = db.Exec(qr, nil)
	assert.Nil(t, err)
	return db
}

func dropDB(tempFilename string, t *testing.T, db *sql.DB) {
	q := "tshirt_catalog_" + t.Name()
	_, _ = db.Exec("DROP TABLE "+q, nil)
	_ = db.Close()

	defer func() {
		err := os.Remove(tempFilename)
		if err != nil {
			t.Error("temp file remove error:", err)
			return
		}
	}()
}
