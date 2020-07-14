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

type Coordinates struct {
	Latitude  int64 `json:"latitude"`
	Longitude int64 `json:"longitude"`
}

type RealtimeInfo struct {
	DepartureTime string `json:"abfahrtIstZeit"`
	DepartureDate string `json:"abfahrtIstDatum"`

	ArrivalTime string `json:"ankunkftIstZeit"`
	ArrivalDate string `json:"ankunkftIstDatum"`

	DeparturePlatformChange bool `json:"abfahrtPlatformChange"`
	ArrivalPlatformChange   bool `json:"ankunftPlatformChange"`

	DepartureCancellation bool `json:"abfahrtCancellation"`
	ArrivalCancellation   bool `json:"ankunftCancellation"`

	DepartureDelayUndefined bool `json:"abfahrtDelayUndefined"`
	ArrivalDelayUndefined   bool `json:"ankunftDelayUndefined"`
}

type OevIcon = string

const (
	Zug OevIcon = "ZUG"
)

// TODO: Fill struct
type TransportDesignation struct{
	OevIcon OevIcon `json:"oevIcon"`
}

type TransportServiceAttribute = string
type Occupancy string

const (
	OM TransportServiceAttribute = "OM"
)

const (
	Low     Occupancy = "LOW"
	Unknown Occupancy = "UNKOWN"
)

type PreviewType string

const (
	Dusp PreviewType = "DUSP"
)

type WalkIcon string

const (
	Pedestrian WalkIcon = "pedestrian"
)

type Section struct {
	DepartureTime string `json:"abfahrtTime"`
	ArrivalTime   string `json:"ankunftTime"`

	DepartureDate string `json:"abfahrtDatum"`
	ArrivalDate   string `json:"ankunftDatum"`

	DepartureName string `json:"abfahrtName"`
	ArrivalName   string `json:"ankunftName"`

	DepartureTrack string `json:"abfahrtGleis"`
	ArrivalTrack   string `json:"ankunftGleis"`

	DepartureTrackLabel string `json:"departureTrackLabel"`
	ArrivalTrackLabel   string `json:"arrivalTrackLabel"`

	DepartureTrackLabelAccessibility string `json:"departureTrackLabelAccessibility"`
	ArrivalTrackLabelAccessibility   string `json:"arrivalTrackLabelAccessibility"`

	DepartureCoordinates Coordinates `json:"abfahrtKoordinaten"`
	ArrivalCoordinates   Coordinates `json:"ankunftKoordinaten"`

	DeparturePlatformChange string `json:"abfahrtPlatformChange"`
	ArrivalPlatformChange   string `json:"ankunftPlatformChange"`

	DepartureCancellation bool `json:"abfahrtCancellation"`
	ArrivalCancellation   bool `json:"ankunftCancellation"`

	RealtimeInfo RealtimeInfo `json:"realtimeInfo"`

	TransportDesignation       TransportDesignation        `json:"transportBezeichnung"`
	PreviewType                PreviewType                 `json:"previewType"`
	TransportServiceAttributes []TransportServiceAttribute `json:"transportServiceAttributes"`
	TransportSuggestion        string                      `json:"transportHinweis"`

	DuspUrl                string `json:"duspUrl"`
	DuspPreviewUrl         string `json:"duspPreviewUrl"`
	DuspNativeStyleUrl     string `json:"duspNativeStyleUrl"`
	DuspNativeDarkStyleUrl string `json:"duspNativeDarkStyleUrl"`
	DuspNativeUrl          string `json:"duspNativeUrl"`

	OccupancyFirst  Occupancy `json:"belegungErste"`
	OccupancySecond string    `json:"belegungZweite"`

	WalkDescription              string   `json:"walkBezeichnung"`
	WalkDescriptionAccessibility string   `json:"walkBezeichnungAccessibility"`
	WalkIcon                     WalkIcon `json:"walkIcon"`

	ActionUrl string `json:"actionUrl"`
}

type Connection struct {
	Sections              []Section `json:"verbindungSections"`
	Destination           string    `json:"ankunft"`
	Via                   []string  `json:"vias"`
	Transfers             int32     `json:"transfers"`
	Duration              string    `json:"duration"`
	DurationAccessibility string    `json:"durationAccessibility"`
	DepartureTime         string    `json:"abfahrtTime"`
	DepartureDate         string    `json:"ankunftDate"`
	ArrivalTime           string    `json:"ankunftTime"`
	ArrivalDate           string    `json:"ankunftDate"`
}

type ConnectionsResult struct {
	Abfahrt              POI           `json:"abfahrt"`
	Ankunft              POI           `json:"ankunft"`
	EarlierUrl           string        `json:"earlierUrl"`
	LaterUrl             string        `json:"laterUrl"`
	LegendBfrItems       []LegendEntry `json:"legendBfrItems"`
	Legend               []LegendEntry `json:"legendItems"`
	LegendOccupancyItems []LegendEntry `json:"legendOccupancyItems"`
	ConnectionPriceUrl   string        `json:"verbindungPreisUrl"`
	Connections          []Connection  `json:"verbindungen"`
}

const ApiBasePath = "/unauth/fahrplanservice/v1"

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
	reqUrl := fmt.Sprintf("%s%s%s/", ApiEndpoint, connectionsPath, connectionsPartial)

	req, err := http.NewRequest("GET", reqUrl, nil)
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
