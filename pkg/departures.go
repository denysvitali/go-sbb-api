package sbb_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (s *SbbApi) GetDepartureTable(from string, fromType PoiType) (*DepartureTable, error) {

	fromTypeConv := "s"
	switch fromType {
	case Station:
		fromTypeConv = "s"
	}

	departureTablePath := fmt.Sprintf("%s/%s/%s/%s", FahrplanV2, "abfahrtstabelle",fromTypeConv, url.PathEscape(from))
	reqUrl := fmt.Sprintf("%s%s", ApiEndpoint, departureTablePath)

	requestUrl, err := url.Parse(reqUrl)
	if err != nil {
		return nil, err
	}

	/*
		ZUG, BAHN_ICE_TGV_RJ, BAHN_EC_IC, BAHN_IR, BAHN_RE_D, BAHN_S_SN_R, BAHN_ARZ_EXT, BUS, SCHIFF, SEILBAHN, TRAM_METRO
	*/

	query := map[string][]string {
		//"nachType": {string(Station)},
		//"verkehrsmittel[]": {"ZUG"},
		"abAn": {"ab"},
		//"nach": {"Altstetten"},
	}

	requestUrl.RawQuery = url.Values(query).Encode()

	req, err := http.NewRequest("GET", requestUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: 200 expected but %d received", resp.StatusCode)
	}

	var departureTable DepartureTable
	dec := json.NewDecoder(resp.Body)

	dec.DisallowUnknownFields()
	err = dec.Decode(&departureTable)
	if err != nil {
		return nil, err
	}

	return &departureTable, nil
}
