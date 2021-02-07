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

type Accessibility string

const (
	Independant Accessibility = "SELBSTAENDIG"
)

type PoiType string

const (
	Station PoiType = "STATION"
)

type POI struct {
	DisplayName string  `json:"displayName"`
	ExternalId  string  `json:"externalId"`
	Type        PoiType `json:"type"`
	Coordinates
	Accessibility Accessibility `json:"barriereFreiheit"`
}

type LegendEntry struct {
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Actions     []string `json:"actions"`
}

type Coordinates struct {
	Latitude  int64 `json:"latitude"`
	Longitude int64 `json:"longitude"`
}

type RealtimeInfo struct {
	DepartureTime string `json:"abfahrtIstZeit"`
	DepartureDate string `json:"abfahrtIstDatum"`

	ArrivalTime string `json:"ankunftIstZeit"`
	ArrivalDate string `json:"ankunftIstDatum"`

	DeparturePlatformChange bool `json:"abfahrtPlatformChange"`
	ArrivalPlatformChange   bool `json:"ankunftPlatformChange"`

	DepartureCancellation bool `json:"abfahrtCancellation"`
	ArrivalCancellation   bool `json:"ankunftCancellation"`

	DepartureDelayUndefined bool `json:"abfahrtDelayUndefined"`
	ArrivalDelayUndefined   bool `json:"ankunftDelayUndefined"`

	Icon                        string `json:"icon"`
	DetailMessage               string `json:"detailMsg"`
	CancellationMessage         string `json:"cancellationMsg"`
	PlatformChange              string `json:"platformChange"`
	NextAlternative             string `json:"nextAlternative"`
	AlternativeMessage          string `json:"alternativeMsg"`
	DetailsMessageAccessibility string `json:"detailsMsgAccessibility"`
	IsAlternative               bool   `json:"isAlternative"`
}

type OevIcon = string

const (
	Zug OevIcon = "ZUG"
)

type TransportIcon string

const (
	Re TransportIcon = "RE"
	Ic TransportIcon = "IC"
)

type TransportDesignation struct {
	OevIcon                 OevIcon       `json:"oevIcon"`
	TransportIcon           TransportIcon `json:"transportIcon"`
	TransportIconSuffix     string        `json:"transportIconSuffix"`
	TransportLabel          string        `json:"transportLabel"`
	TransportText           string        `json:"transportText"`
	TransportName           string        `json:"transportName"`
	TransportDirection      string        `json:"transportDirection"`
	TransportLabelBgColor   string        `json:"transportLabelBgColor"`
	TransportLabelTextColor string        `json:"transportLabelTextColor"`
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
	SectionConnection

	// Departure
	DepartureCancellation   bool        `json:"abfahrtCancellation"`
	DepartureDate           string      `json:"abfahrtDatum"`
	DepartureTrack          string      `json:"abfahrtGleis"`
	DepartureCoordinates    Coordinates `json:"abfahrtKoordinaten"`
	DepartureName           string      `json:"abfahrtName"`
	DeparturePlatformChange bool        `json:"abfahrtPlatformChange"`

	ActionUrl string `json:"actionUrl"`

	// Arrival
	ArrivalCancellation            bool        `json:"ankunftCancellation"`
	ArrivalDate                    string      `json:"ankunftDatum"`
	ArrivalTrack                   string      `json:"ankunftGleis"`
	ArrivalCoordinates             Coordinates `json:"ankunftKoordinaten"`
	ArrivalName                    string      `json:"ankunftName"`
	ArrivalPlatformChange          bool        `json:"ankunftPlatformChange"`
	ArrivalTrackLabel              string      `json:"arrivalTrackLabel"`
	ArrivalTrackLabelAccessibility string      `json:"arrivalTrackLabelAccessibility"`

	DepartureTrackLabel              string `json:"departureTrackLabel"`
	DepartureTrackLabelAccessibility string `json:"departureTrackLabelAccessibility"`

	DuspUrl                string `json:"duspUrl"`
	DuspPreviewUrl         string `json:"duspPreviewUrl"`
	DuspNativeStyleUrl     string `json:"duspNativeStyleUrl"`
	DuspNativeDarkStyleUrl string `json:"duspNativeDarkStyleUrl"`
	DuspNativeUrl          string `json:"duspNativeUrl"`

	DurationPercentage string `json:"durationProzent"`

	FormationUrl string       `json:"formationUrl"`
	PreviewType  PreviewType  `json:"previewType"`
	RealtimeInfo RealtimeInfo `json:"realtimeInfo"`

	TransportDesignation       TransportDesignation        `json:"transportBezeichnung"`
	TransportSuggestion        string                      `json:"transportHinweis"`
	TransportServiceAttributes []TransportServiceAttribute `json:"transportServiceAttributes"`

	Type string `json:"type"`

	WalkDescription              string   `json:"walkBezeichnung"`
	WalkDescriptionAccessibility string   `json:"walkBezeichnungAccessibility"`
	WalkIcon                     WalkIcon `json:"walkIcon"`
}

type SectionConnection struct {
	DepartureTime string `json:"abfahrtTime"`
	ArrivalTime   string `json:"ankunftTime"`

	DepartureCancellation bool `json:"abfahrtCancellation"`
	ArrivalCancellation   bool `json:"ankunftCancellation"`

	DepartureTrack string `json:"abfahrtGleis"`

	OccupancyFirst  Occupancy `json:"belegungErste"`
	OccupancySecond string    `json:"belegungZweite"`

	RealtimeInfo RealtimeInfo `json:"realtimeInfo"`
}

type TicketingInfo struct {
	ButtonText    string `json:"buttonText"`
	DialogMessage string `json:"dialogMessage"`
	DialogTitle   string `json:"dialogTitle"`
	IsAvailable   bool   `json:"isAvailable"`
}

type Connection struct {
	SectionConnection

	DepartureDate string `json:"abfahrtDate"`
	ArrivalDate   string `json:"ankunftDate"` // Can't merge w/ Section because this field is name .*Datum

	DayDifference              string `json:"dayDifference"`
	DayDifferenceAccessibility string `json:"dayDifferenceAccessibility"`

	DepartureTrackLabel              string `json:"departureTrackLabel"`
	DepartureTrackLabelAccessibility string `json:"departureTrackLabelAccessibility"`

	Departure   string `json:"abfahrt"`
	Destination string `json:"ankunft"`

	Duration              string `json:"duration"`
	DurationAccessibility string `json:"durationAccessibility"`

	IsInternationalConnection bool `json:"isInternationalVerbindung"`

	LegendBfrItems       []LegendEntry `json:"legendBfrItems"`
	LegendItems          []LegendEntry `json:"legendItems"`
	LegendOccupancyItems []LegendEntry `json:"legendOccupancyItems"`

	ReconstructionContext string   `json:"reconstructionContext"`
	ServiceAttributes     []string `json:"serviceAttributes"`

	TicketingInfo        TicketingInfo        `json:"ticketingInfo"`
	Transfers            int32                `json:"transfers"`
	TransportDesignation TransportDesignation `json:"transportBezeichnung"`

	ConnectionDiscountContext string    `json:"verbindungAbpreisContext"`
	ConnectionId              string    `json:"verbindungId"`
	Sections                  []Section `json:"verbindungSections"`
	TrafficStage              []string  `json:"verkehrstage"`
	Via                       []string  `json:"vias"`
	SurchargeObligation       bool      `json:"zuschlagspflicht"`
	OfferUrl                  string    `json:"angeboteUrl"`
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

const FahrplanV1 = "/unauth/fahrplanservice/v1"
const FahrplanV2 = "/unauth/fahrplanservice/v2"

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
	connectionsPath := fmt.Sprintf("%s%s", FahrplanV1, "/verbindungen")
	connectionsPartial := formatConnectionsUrl(url.PathEscape(from), url.PathEscape(to), at)
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

type TrainRoute struct {
	TransportDesignation TransportDesignation `json:"transportBezeichnung"`
	ServiceAttributes    []string             `json:"serviceAttributes"`
	Stations             []POI                `json:"stations"`
	TrafficStage         []string             `json:"verkehrstage"`
	LegendItems          []LegendEntry        `json:"legendItems"`
	LegendOccupancyItems []LegendEntry        `json:"legendOccupancyItems"`
	RefreshUrl           string               `json:"refreshUrl"`
	RealtimeAnnounces    string               `json:"realtimeMeldungen"`
}

type Departure struct {
	FromStation            POI        `json:"vonHaltestelle"`
	DepartureAt            string     `json:"abfahrt"`
	DepartureDate          string     `json:"abfahrtDatum"`
	EffectiveDeparture     string     `json:"abfahrtAktuell"`
	EffectiveDepartureDate string     `json:"abfahrtAktuellDatum"`
	Direction              string     `json:"richtung"`
	Platform               string     `json:"gleis"`
	PlatformChange         bool       `json:"platformChange"`
	Cancellation           bool       `json:"cancellation"`
	Delayed                bool       `json:"delayed"`
	TransportNewStops      bool       `json:"transportNewStops"`
	TransportPassageStops  bool       `json:"transportPassageStops"`
	HasAlternativeStop     bool       `json:"hasAlternativeStop"`
	TrainFormation         *string    `json:"zugformation"`
	TrainRoute             TrainRoute `json:"zuglauf"`
	AlternativeStop        bool       `json:"alternativeStop"`
}

type DepartureTable struct {
	FromStation POI         `json:"abHaltestelle"`
	Departures  []Departure `json:"abfahrts"`
	TrackLabel  string      `json:"trackLabel"`
}