package sbb_api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCertificateHash(t *testing.T) {
	hash, err := getCertificateHash()
	assert.Nil(t, err)
	assert.Equal(t, "WdfnzdQugRFUF5b812hZl3lAahM=", hash)
}

func TestUrl1(t *testing.T){
	date := "2020-06-28"
	auth := "hL89gUidDebOUNUCP/+5vbj+0Iw="
	path := "/unauth/fahrplanservice/v1/verbindungen/s/Z%25C3%25BCrich%2520HB/s/Bern/ab/2019-09-20/21-14/"

	resAuth := GetAuthorization(path, date)
	assert.Equal(t, auth, resAuth)
}