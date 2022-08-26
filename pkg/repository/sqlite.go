package repository

import (
	"crypto/md5"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	// SQLite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/renomarx/fizzbuzz/pkg/core/model"
	"github.com/sirupsen/logrus"
)

// SQLiteRepo repository using SQLite db
type SQLiteRepo struct {
	db *sqlx.DB
}

type requestsCounter struct {
	Hash    string `db:"hash"`
	Int1    int    `db:"int1"`
	Int2    int    `db:"int2"`
	Limit   int    `db:"lim"`
	Str1    string `db:"str1"`
	Str2    string `db:"str2"`
	Counter int    `db:"counter"`
}

func (c *requestsCounter) FromParams(params model.Params) {
	c.Int1 = params.Int1
	c.Int2 = params.Int2
	c.Limit = params.Limit
	c.Str1 = params.Str1
	c.Str2 = params.Str2
	str := fmt.Sprintf("%d_%d_%d_%s_%s", c.Int1, c.Int2, c.Limit, c.Str1, c.Str2)
	c.Hash = fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func (c *requestsCounter) ToStats() model.Stats {
	return model.Stats{
		Int1:    c.Int1,
		Int2:    c.Int2,
		Limit:   c.Limit,
		Str1:    c.Str1,
		Str2:    c.Str2,
		Counter: c.Counter,
	}
}

// NewSQLIteRepo SQLiteRepo constructor
func NewSQLIteRepo() *SQLiteRepo {
	// For in-memory: dsn = ":memory:"
	dsn := os.Getenv("SQLITE_DSN")
	db, err := sqlx.Open("sqlite3", dsn)
	if err != nil {
		logrus.Fatalf("Impossible to connect to sqlite on DSN %s: %s", dsn, err)
	}

	// force a connection and test that it worked
	err = db.Ping()
	if err != nil {
		logrus.Fatal(err)
	}
	return &SQLiteRepo{
		db: db,
	}
}

func (repo *SQLiteRepo) Inc(params model.Params, number int) error {
	requestCounter := requestsCounter{}
	requestCounter.FromParams(params)
	// Insert if not exists
	_, err := repo.db.NamedExec(`
		INSERT OR IGNORE INTO requests_counters
		(int1, int2, lim, str1, str2, counter)
		VALUES
		(:int1, :int2, :lim, :str1, :str2, 0)
		`, requestCounter)
	if err != nil {
		repo.error("error inserting requests_counter", err)
		return err
	}
	// Update
	requestCounter.Counter = number
	res, err := repo.db.NamedExec(`
		UPDATE requests_counters SET counter = counter + :counter
		WHERE int1=:int1
		AND int2=:int2
		AND int2=:int2
		AND lim=:lim
		AND str1=:str1
		AND str2=:str2
		`, requestCounter)
	if err != nil {
		repo.error("error incrementing requests_counter", err)
		return err
	}
	nbRows, err := res.RowsAffected()
	if err != nil {
		repo.error("error getting number of rows affected", err)
		return err
	}
	if nbRows != 1 {
		err = fmt.Errorf("no requests_counter match in database")
		repo.error("no requests_counter match in database", err)
		return err
	}
	return nil
}

func (repo *SQLiteRepo) GetMaxStats() (stats model.Stats, err error) {
	rc := requestsCounter{}
	err = repo.db.Get(&rc, "SELECT * FROM requests_counters ORDER BY counter DESC LIMIT 1")
	if err != nil {
		repo.error("error getting max counter of requests_counters", err)
		return
	}
	stats = rc.ToStats()
	return
}

func (repo *SQLiteRepo) error(msg string, err error) {
	model.AppMetrics.IncDatabaseErrors(msg)
	logrus.Error(err)
}
