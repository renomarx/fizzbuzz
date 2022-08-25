package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/renomarx/fizzbuzz/pkg/core/model"
	"github.com/renomarx/fizzbuzz/pkg/core/ports"
	"github.com/renomarx/fizzbuzz/pkg/core/service"

	"github.com/julienschmidt/httprouter"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"

	"github.com/swaggo/http-swagger"
)

func init() {
	logrus.AddHook(filename.NewHook())
}

// RestAPI Main service handling http requests
type RestAPI struct {
	MetricsController *metricsController
	fizzbuzzSVC       ports.FizzbuzzService
}

type RestAPIError struct {
	Error string `json:"error"`
}

// NewRestAPI RestAPI constructor with dependencies injected
func NewRestAPI() *RestAPI {
	fizzbuzzSVC := service.NewFizzbuzzSVC()
	return &RestAPI{
		MetricsController: NewMetricsController(),
		fizzbuzzSVC:       fizzbuzzSVC,
	}
}

// Serve Http listen to REST_PORT
func (api *RestAPI) Serve() {
	router := httprouter.New()
	api.Route(router)
	port := os.Getenv("REST_PORT")
	if port == "" {
		logrus.Fatalf("No REST_PORT env variable found")
	}

	// API doc
	router.GET("/docs/:any", swaggerHandler)

	logrus.Infof("API listening on %s", port)
	logrus.Fatal(http.ListenAndServe(port, router))
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}

// Route configure http router
func (api *RestAPI) Route(r *httprouter.Router) {
	// Handlers
	r.GET("/ping", api.Ping)

	// Because I like to have some metrics on every service, and I like prometheus
	r.Handler("GET", "/metrics", api.MetricsController.HTTPHandler())

	logrus.Infof("Serving metrics on /metrics")
	r.NotFound = http.HandlerFunc(api.NotFound)

	r.GET("/fizzbuzz", api.GenerateFizzbuzz)
	logrus.Infof("Serving /fizzbuzz")
}

// Ping handle /ping http requests (for health checks)
// Always useful for monitoring tools or orchestrators like K8s or Nomad
func (api *RestAPI) Ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}

// NotFound handle http routes not found by router - only to trace bad 404 http calls
func (api *RestAPI) NotFound(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("HTTP Not found: %s", r.URL.Path)
	metrics := api.MetricsController.GetMetrics()
	metrics.RouterHTTPNotFound.Inc()
	w.WriteHeader(http.StatusNotFound)
}

// GenerateFizzbuzz generate a fizzbuzz string from params
// @Router       /fizzbuzz [get]
// @Summary      fizzbuzz
// @Description  generate a fizzbuzz string from params
// @Produce      json
// @Success      200  {array} string
// @Failure      500
// @Failure      400
func (api *RestAPI) GenerateFizzbuzz(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := model.Params{}
	query := r.URL.Query()
	err := params.Parse(
		query.Get("int1"),
		query.Get("int2"),
		query.Get("limit"),
		query.Get("str1"),
		query.Get("str2"))
	if err != nil {
		api.sendJSONError(w, http.StatusBadRequest, err)
		return
	}
	err = params.Validate()
	if err != nil {
		api.sendJSONError(w, http.StatusBadRequest, err)
		return
	}
	result := api.fizzbuzzSVC.Fizzbuzz(params)
	api.sendSuccessJSONObject(w, result)
}

func (api *RestAPI) sendSuccessJSONObject(w http.ResponseWriter, o interface{}) {
	b, err := json.Marshal(o)
	if err != nil {
		api.sendJSONError(w, http.StatusInternalServerError, err)
		return
	}
	api.sendJSONResponse(w, http.StatusOK, b)
}

func (api *RestAPI) sendJSONResponse(w http.ResponseWriter, statusCode int, msg []byte) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}

func (api *RestAPI) sendJSONError(w http.ResponseWriter, statusCode int, err error) {
	logrus.Errorf("API error %d: %s", statusCode, err)
	responseError := RestAPIError{
		Error: err.Error(),
	}
	jsonError, _ := json.Marshal(responseError)
	api.sendJSONResponse(w, statusCode, jsonError)
}
