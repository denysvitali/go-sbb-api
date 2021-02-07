package sbb_api

import (
	"github.com/docker/distribution/uuid"
	"net/http"
	"time"
)

const ApiEndpoint = "https://p1.sbbmobile.ch"
const SBB_UA = "SBBmobile/flavorProdRelease-11.7.2.42.master Android/10 (OnePlus;A5010)"

func setHeaders(req *http.Request) {
	now := time.Now()
	req.Header.Set("X-APP-TOKEN", generateToken())
	req.Header.Set("X-API-AUTHORIZATION", GetAuthorization(req.URL.Path, now))
	req.Header.Set("X-API-DATE", now.Format("2006-01-02"))
	req.Header.Set("User-Agent", SBB_UA)
}

func generateToken() string {
	return uuid.Generate().String()
}