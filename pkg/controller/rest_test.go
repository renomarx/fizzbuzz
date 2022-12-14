package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/renomarx/fizzbuzz/pkg/core/model"
	"github.com/stretchr/testify/assert"
)

type FizzbuzzSVCMock struct {
	Params model.Params
	Result []string
}

func (svc *FizzbuzzSVCMock) Fizzbuzz(params model.Params) []string {
	svc.Params = params
	return svc.Result
}

type RequestsRepoMock struct {
	Params  model.Params
	Counter int
	Error   error
	Stats   model.Stats
}

func (repo *RequestsRepoMock) Inc(params model.Params, number int) error {
	repo.Params = params
	repo.Counter = number
	return repo.Error
}

func (repo *RequestsRepoMock) GetMaxStats() (model.Stats, error) {
	return repo.Stats, repo.Error
}

func NewMockedAPI() *RestAPI {
	return &RestAPI{
		MetricsController: NewMetricsController(),
	}
}

func TestRestAPIGenerateFizzbuzz(t *testing.T) {
	api := NewMockedAPI()
	svc := &FizzbuzzSVCMock{
		Result: []string{"1", "2", "fizz"},
	}
	api.fizzbuzzSVC = svc
	repo := &RequestsRepoMock{}
	api.requestsRepo = repo

	router := httprouter.New()
	api.Route(router)

	req, err := http.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=3&str1=fizz&str2=buzz", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	expectedParams := model.Params{
		Int1:  3,
		Int2:  5,
		Limit: 3,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	assert.Equal(t, expectedParams, svc.Params)
	assert.Equal(t, expectedParams, repo.Params)
	assert.Equal(t, 1, repo.Counter)
}

func TestRestAPIGenerateFizzbuzzBadRequest(t *testing.T) {
	api := NewMockedAPI()

	router := httprouter.New()
	api.Route(router)

	req, err := http.NewRequest("GET", "/fizzbuzz?int1=toto&int2=5&limit=3&str1=fizz&str2=buzz", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestRestAPIGenerateFizzbuzzValidationFailed(t *testing.T) {
	api := NewMockedAPI()

	router := httprouter.New()
	api.Route(router)

	req, err := http.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=0", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestRestAPIPing(t *testing.T) {
	api := NewMockedAPI()

	router := httprouter.New()
	api.Route(router)

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Bad http response code %d", rr.Code)
	}
}

func TestRestAPI404(t *testing.T) {
	api := NewMockedAPI()

	router := httprouter.New()
	api.Route(router)

	req, err := http.NewRequest("GET", "/inexistant-route-404-not-found", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Bad http response code %d", rr.Code)
	}
}
