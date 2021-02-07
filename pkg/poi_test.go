package sbb_api

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPoiNearbyParse(t *testing.T) {
	file, err := os.Open("../resources/test/standorte-nearby.json")
	assert.Nil(t, err)

	var poiNearby PoiNearby
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&poiNearby)
	assert.Nil(t, err)
	spew.Dump(poiNearby)
}