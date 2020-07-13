package sbb_api

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetCertificateHash(t *testing.T) {
	hash, err := getCertificateHash()
	assert.Nil(t, err)
	assert.Equal(t, "WdfnzdQugRFUF5b812hZl3lAahM=", hash)
}

func TestUrl1(t *testing.T){
	date := time.Date(2020, 6, 28, 0,0,0,0, time.UTC)
	auth := "hL89gUidDebOUNUCP/+5vbj+0Iw="
	path := "/unauth/fahrplanservice/v1/verbindungen/s/Z%25C3%25BCrich%2520HB/s/Bern/ab/2019-09-20/21-14/"

	resAuth := GetAuthorization(path, date)
	assert.Equal(t, auth, resAuth)
}

func TestConnections(t *testing.T){
	sbb_api := New()
	res, err := sbb_api.GetConnections("Zürich", "Bern", time.Now())
	assert.Nil(t, err)
	logrus.Printf("%v", res)
}

func TestParseErrorMessage(t *testing.T) {
	//connectionResultJson := "{\"title\":null,\"serviceErrorCode\":\"FIS-1015\",\"message\":null,\"developerMessage\":null,\"exception\":null,\"hasRetry\":false}"
	//parseConnectionResultJson(connectionResultJson)
}