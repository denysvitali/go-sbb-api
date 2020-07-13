package sbb_api

import (
	"fmt"
	"net/http"
	"time"
)

type Connection struct {

}

const ApiBasePath = "/unauth/fahrplanservice/v1";

type SbbApi struct {
	client *http.Client
}

func New() *SbbApi {
	return &SbbApi{
		client: http.DefaultClient,
	}
}

func formatConnectionsUrl(from string, to string, at time.Time) string {
	return fmt.Sprintf("s/%s/s/%s/ab/%s/%s",
		from,
		to,
		at.Format("2006-01-02"),
		at.Format("15-04"),
	)
}

func (s *SbbApi) GetConnections(from string, to string, at time.Time) ([]Connection, error) {
	connectionsPath := fmt.Sprintf("%s%s", ApiBasePath, "/verbindungen")
	connectionsPartial := formatConnectionsUrl(from, to, at)
	_ = fmt.Sprintf("%s%s", connectionsPath, connectionsPartial)
	return nil, nil
	// s/Zurich/s/Bern/ab/2019-09-20/10-14/

}