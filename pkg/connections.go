package sbb_api

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type POI struct {

}

type LegendEntry struct {

}

type ConnectionsResult struct {
	Abfahrt POI `json:"abfahrt"`
	Ankunft POI `json:"ankunft"`
	EarlierUrl string `json:"earlierUrl"`
	LaterUrl string `json:"laterUrl"`
	LegendBfrItems []LegendEntry `json:"legendBfrItems"`
	Legend []LegendEntry `json:"legend"`
}

const ApiBasePath = "/unauth/fahrplanservice/v1";

type SbbApi struct {
	client *http.Client
}

func New() *SbbApi {

	customClient := http.DefaultClient

	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	caCert, err := GetCACert()
	if err != nil {
		logrus.Panic(err)
	}
	rootCAs.AddCert(caCert)

	config := &tls.Config{RootCAs: rootCAs}
	tr := &http.Transport{TLSClientConfig: config}
	customClient.Transport = tr

	return &SbbApi{
		client: http.DefaultClient,
	}
}

func formatConnectionsUrl(from string, to string, at time.Time) string {
	return fmt.Sprintf("/s/%s/s/%s/ab/%s/%s",
		url.QueryEscape(from),
		url.QueryEscape(to),
		at.Format("2006-01-02"),
		at.Format("15-04"),
	)
}

func (s *SbbApi) GetConnections(from string, to string, at time.Time) (*ConnectionsResult, error) {
	connectionsPath := fmt.Sprintf("%s%s", ApiBasePath, "/verbindungen")
	connectionsPartial := formatConnectionsUrl(from, to, at)
	url := fmt.Sprintf("%s%s%s/", ApiEndpoint, connectionsPath, connectionsPartial)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var connectionResult ConnectionsResult
	err = json.Unmarshal(data, &connectionResult)

	if err != nil {
		return nil, err
	}

	return &connectionResult, nil
}