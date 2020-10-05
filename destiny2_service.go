package destiny2lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"root/responses"
)

type Destiny2Service struct {
	s *Service
}

func NewDestiny2Service(s *Service) *Destiny2Service {
	return &Destiny2Service{s: s}
}

func (s *Destiny2Service) GetDestinyManifest() *GetDestinyManifestCall {
	return &GetDestinyManifestCall{s: s.s}
}

type GetDestinyManifestCall struct {
	s *Service
}

func (c *GetDestinyManifestCall) Do() (*responses.GetDestinyManifestResponse, error) {

	// Make the request.
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response.
	var ret = &responses.GetDestinyManifestResponse{}
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	// Check the error code.
	if ret.ErrorCode != 1 {
		return nil, fmt.Errorf("%s: %s", ret.ErrorStatus, ret.Message)
	}

	return ret, nil
}

func (c *GetDestinyManifestCall) doRequest() (*http.Response, error) {

	// Create the url.
	url := c.s.BasePath + "Destiny2/Manifest/"

	// Create the request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Execute the request.
	return c.s.client.Do(req)
}

func (s *Destiny2Service) GetDestinyEntityDefinition(entityType, hashIdentifier string) *GetDestinyEntityDefinitionCall {
	return &GetDestinyEntityDefinitionCall{
		s:              s.s,
		entityType:     entityType,
		hashIdentifier: hashIdentifier,
	}
}

type GetDestinyEntityDefinitionCall struct {
	s              *Service
	entityType     string
	hashIdentifier string
	header         *http.Header
}

func (c *GetDestinyEntityDefinitionCall) Do() (*responses.GetEntityDefinitionResponse, error) {

	// Make the request.
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response.
	var ret = &responses.GetEntityDefinitionResponse{}
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	// Check the error code.
	if ret.ErrorCode != 1 {
		return nil, fmt.Errorf("%s: %s", ret.ErrorStatus, ret.Message)
	}

	return ret, nil
}

func (c *GetDestinyEntityDefinitionCall) doRequest() (*http.Response, error) {

	// Create url.
	url := fmt.Sprintf("%sDestiny2/Manifest/%s/%s",
		c.s.BasePath,
		c.entityType,
		c.hashIdentifier)

	// Create request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-KEY", c.s.Config.ApiKey)

	// Execute the request.
	return c.s.client.Do(req)
}

func (s *Destiny2Service) SearchDestinyPlayer(membershipType, displayName string) *SearchDestinyPlayerCall {
	return &SearchDestinyPlayerCall{
		s:              s.s,
		membershipType: membershipType,
		displayName:    displayName,
		queryParams:    map[string]string{},
	}
}

type SearchDestinyPlayerCall struct {
	s              *Service
	membershipType string
	displayName    string
	queryParams    map[string]string
	header         *http.Header
}

func (c *SearchDestinyPlayerCall) ReturnOriginalProfile(arg string) *SearchDestinyPlayerCall {
	c.queryParams["returnOriginalProfile"] = arg
	return c
}

func (c *SearchDestinyPlayerCall) Do() (*responses.SearchDestinyPlayerResponse, error) {

	// Make the request.
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response.
	var ret = &responses.SearchDestinyPlayerResponse{}
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	// Check the error code.
	if ret.ErrorCode != 1 {
		return nil, fmt.Errorf("%s: %s", ret.ErrorStatus, ret.Message)
	}

	return ret, nil
}

func (c *SearchDestinyPlayerCall) doRequest() (*http.Response, error) {

	// Setup url.
	url := fmt.Sprintf("%sDestiny2/SearchDestinyPlayer/%s/%s",
		c.s.BasePath,
		c.membershipType,
		c.displayName)

	// Atatch params params to the url.
	params := "?"
	for k, v := range c.queryParams {
		params += k + "=" + v + "&"
	}
	url += params

	// Create request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-KEY", c.s.Config.ApiKey)

	// Execute the request.
	return c.s.client.Do(req)
}
