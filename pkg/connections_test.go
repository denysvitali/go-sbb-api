package sbb_api

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestConnectionURL(t *testing.T){
	res := formatConnectionsUrl("Zürich", "Bern", time.Date(
		2020,
		7,
		13,
		21,
		0,
		0,
		0,
		time.UTC),
	)

	assert.Equal(t, "/s/Z%C3%BCrich/s/Bern/ab/2020-07-13/21-00", res)
}

func TestGetConnections(t *testing.T) {
	s := New()
	connections, err := s.GetConnections("Zürich HB", "Altstetten", time.Now())
	assert.Nil(t, err)
	assert.NotNil(t, connections)
	fmt.Printf("connections=%v", connections)
}


func TestParseResult(t *testing.T){
	file, err := os.Open("../resources/test/result-1.json")
	assert.Nil(t, err)

	var connections ConnectionsResult
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&connections)
	assert.Nil(t, err)
	//logrus.Printf("%+v\n", connections)
	spew.Dump(connections)
}

func TestGetDepartureTable(t *testing.T) {
	s := New()
	depTable, err := s.GetDepartureTable("Zürich HB", Station)
	assert.Nil(t, err)

	fmt.Printf("departureTable=%v", depTable)
}

func TestPoiNearby(t *testing.T) {
	s := New()
	poiNearby, err := s.GetPoiNearby(47378177, 8540193)
	assert.Nil(t, err)

	fmt.Printf("poiNearby=%v", poiNearby)
}