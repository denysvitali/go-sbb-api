package sbb_api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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

	fileContent, err := ioutil.ReadAll(file)
	assert.Nil(t, err)

	var connections ConnectionsResult
	err = json.Unmarshal(fileContent, &connections)
	assert.Nil(t, err)
	logrus.Printf("%v", connections)
}