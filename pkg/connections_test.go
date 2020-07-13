package sbb_api

import (
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, "s/Zürich/s/Bern/ab/2020-07-13/21-00", res)
}
