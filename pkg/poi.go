package sbb_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PoiNearby struct {
	Poi []POI `json:"standorte"`
}

func (s *SbbApi) GetPoiNearby(latitude int, longitude int) (*PoiNearby, error) {
	poiNearbyPath := fmt.Sprintf("%s/%s/%d/%d/", FahrplanV2, "standortenearby", latitude, longitude)
	reqUrl := fmt.Sprintf("%s%s", ApiEndpoint, poiNearbyPath)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var poiNearby PoiNearby
	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&poiNearby)
	if err != nil {
		return nil, err
	}

	return &poiNearby, nil
}
