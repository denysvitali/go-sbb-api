package sbb_api

import (
	"encoding/json"
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