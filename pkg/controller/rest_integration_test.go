package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/renomarx/fizzbuzz/pkg/core/model"
	"github.com/renomarx/fizzbuzz/pkg/repository"
	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"
)

func initTestDB(t *testing.T, db *sqlx.DB) {
	db.MustExec(`
    CREATE TABLE IF NOT EXISTS requests_counters (
      int1 INTEGER,
      int2 INTEGER,
      lim INTEGER,
      str1 TEXT,
      str2 TEXT,
      counter INTEGER,
      PRIMARY KEY(int1,int2,lim,str1,str2)
    );

    CREATE INDEX requests_counters_encoded_params_idx ON requests_counters (counter DESC);
    `)
}

func TestRestAPIWholeFlow(t *testing.T) {
	os.Setenv("SQLITE_DSN", ":memory:")

	// Instanciating REST API
	api := NewRestAPI()

	// Initialization of test database
	requestsRepo, ok := api.requestsRepo.(*repository.SQLiteRepo)
	if !ok {
		t.Fatal("requestsRepo is not SQLiteRepo - implementation has changed ?")
	}
	initTestDB(t, requestsRepo.DB)

	router := httprouter.New()
	api.Route(router)

	// Getting first stats: empty
	var stats model.Stats
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	response := rr.Body.Bytes()
	json.Unmarshal(response, &stats)
	assert.Equal(t, 0, stats.Counter)

	// First fizzbuzz request
	req, err = http.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=16&str1=fizz&str2=buzz", nil)
	if err != nil {
		t.Error(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Getting stats: should be fizz buzz params
	req, err = http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Error(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	response = rr.Body.Bytes()
	json.Unmarshal(response, &stats)
	assert.Equal(t, model.Stats{
		Int1:    3,
		Int2:    5,
		Limit:   16,
		Str1:    "fizz",
		Str2:    "buzz",
		Counter: 1,
	}, stats)

	// 2 fizzbuzz requests with gin tonic params
	req, err = http.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=16&str1=gin&str2=tonic", nil)
	if err != nil {
		t.Error(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	req, err = http.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=16&str1=gin&str2=tonic", nil)
	if err != nil {
		t.Error(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Getting stats again: should be gin tonic params, with counter = 2
	req, err = http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Error(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	response = rr.Body.Bytes()
	json.Unmarshal(response, &stats)
	assert.Equal(t, model.Stats{
		Int1:    3,
		Int2:    5,
		Limit:   16,
		Str1:    "gin",
		Str2:    "tonic",
		Counter: 2,
	}, stats)
}
