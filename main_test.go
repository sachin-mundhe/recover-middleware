package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestM(t *testing.T) {
	templistenAndServe := listenAndServeFunc
	defer func() {
		listenAndServeFunc = templistenAndServe
	}()
	listenAndServeFunc = func(port string, hanle http.Handler) error {
		panic("testing")
	}
	assert.PanicsWithValuef(t, "testing", main, "they should be equal")
}

func TestMain(m *testing.M) {
	m.Run()
}
