package sbb_api

import (
	"github.com/docker/distribution/uuid"
	"net/http"
	"time"
)

const ApiEndpoint = "https://p1.sbbmobile.ch"
const SBB_UA = "SBBmobile/flavorProdRelease-10.8.1 Android/10 (OnePlus;ONEPLUS A5010)"

func setHeaders(req *http.Request) {
	now := time.Now()
	req.Header.Set("X-App-Token", generateToken())
	req.Header.Set("X-API-AUTHORIZATION", GetAuthorization(req.URL.Path, now))
	req.Header.Set("X-API-DATE", now.Format("2006-01-02"))
	req.Header.Set("User-Agent", SBB_UA)
}

func generateToken() string {
	return uuid.Generate().String()
}