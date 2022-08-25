package main

import (
	_ "github.com/renomarx/fizzbuzz/docs"
	"github.com/renomarx/fizzbuzz/pkg/controller"
	"github.com/sirupsen/logrus"
)

// @title Fizzbuzz API
// @version 1.0
func main() {
	logrus.Println("APP STARTING")

	api := controller.NewRestAPI()

	api.Serve()

}
