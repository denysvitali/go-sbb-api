package sbb_api

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseDepartureTableResult(t *testing.T) {
	file, err := os.Open("../resources/test/abfahrt.json")
	assert.Nil(t, err)

	var departureTable DepartureTable
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&departureTable)
	assert.Nil(t, err)
	spew.Dump(departureTable)
}