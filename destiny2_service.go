package bungo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type destiny2Service struct {
	s *Service
}

func NewDestiny2Service(s *Service) *destiny2Service {
	return &destiny2Service{s: s}
}

func (s *destiny2Service) GetDestinyManifest() *GetDestinyManifestCall {
	return &GetDestinyManifestCall{s: s.s}
}

type GetDestinyManifestCall struct {
	s      *Service
	header http.Header
}

func (c *GetDestinyManifestCall) Hander() http.Header {

	if c.header == nil {
		c.header = make(http.Header)
	}

	return c.header
}

func (c *GetDestinyManifestCall) Do() (*GetDestinyManifestResponse, error) {

	// Make the request.
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response.
	var ret = &GetDestinyManifestResponse{}
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
	url := c.s.basePath + "Destiny2/Manifest/"

	// Create the request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}

func (s *destiny2Service) GetDestinyEntityDefinition(entityType, hashIdentifier string) *GetDestinyEntityDefinitionCall {
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
	header         http.Header
}

func (c *GetDestinyEntityDefinitionCall) Hander() http.Header {

	if c.header == nil {
		c.header = make(http.Header)
	}

	return c.header
}

func (c *GetDestinyEntityDefinitionCall) Do() (*GetEntityDefinitionResponse, error) {

	// Make the request.
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response.
	var ret = &GetEntityDefinitionResponse{}
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
		c.s.basePath,
		c.entityType,
		c.hashIdentifier)

	// Create request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}

func (s *destiny2Service) SearchDestinyPlayer(membershipType, displayName string) *SearchDestinyPlayerCall {
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
	header         http.Header
}

func (c *SearchDestinyPlayerCall) Hander() http.Header {

	if c.header == nil {
		c.header = make(http.Header)
	}

	return c.header
}

func (c *SearchDestinyPlayerCall) ReturnOriginalProfile(arg string) *SearchDestinyPlayerCall {
	c.queryParams["returnOriginalProfile"] = arg
	return c
}

func (c *SearchDestinyPlayerCall) Do() (*SearchDestinyPlayerResponse, error) {

	// Make the request.
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response.
	var ret = &SearchDestinyPlayerResponse{}
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
	url := fmt.Sprintf("%sDestiny2/SearchDestinyPlayer/%s/%s?",
		c.s.basePath,
		c.membershipType,
		c.displayName)

	// Atatch params params to the url.
	for k, v := range c.queryParams {
		url += k + "=" + v + "&"
	}

	// Create request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}

func (s *destiny2Service) EquipItem(body *ItemActionRequest) *EquipItemCall {
	return &EquipItemCall{s: s.s, requestBody: body}
}

type EquipItemCall struct {
	s           *Service
	requestBody *ItemActionRequest
	header      http.Header
}

func (c *EquipItemCall) Hander() http.Header {

	if c.header == nil {
		c.header = make(http.Header)
	}

	return c.header
}

func (c *EquipItemCall) Do() (*EquipItemResponse, error) {

	// Make the request.
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response.
	var ret = &EquipItemResponse{}
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	// Check the error code.
	if ret.ErrorCode != 1 {
		return nil, fmt.Errorf("%s: %s", ret.ErrorStatus, ret.Message)
	}

	return ret, nil
}

func (c *EquipItemCall) doRequest() (*http.Response, error) {

	// Setup url.
	url := fmt.Sprintf("%sDestiny2/Actions/Items/EquipItem", c.s.basePath)

	// Marshal the request body.
	jsonValue, err := json.Marshal(c.requestBody)
	if err != nil {
		return nil, err
	}

	// Create request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}

func (s *destiny2Service) GetLinkedProfiles(membershipId, membershipType string) *GetLinkedProfilesCall {
	return &GetLinkedProfilesCall{
		s:              s.s,
		membershipId:   membershipId,
		membershipType: membershipType,
	}
}

type GetLinkedProfilesCall struct {
	s              *Service
	membershipId   string
	membershipType string
	queryParams    map[string]string
	header         http.Header
}

func (c *GetLinkedProfilesCall) Hander() http.Header {

	if c.header == nil {
		c.header = make(http.Header)
	}

	return c.header
}

func (c *GetLinkedProfilesCall) GetAllMemeberships(arg string) *GetLinkedProfilesCall {
	c.queryParams["getAllMemberships"] = arg
	return c
}

func (c *GetLinkedProfilesCall) Do() (*LinkedProfilesResponse, error) {

	// Make the request.
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response.
	var ret = &LinkedProfilesResponse{}
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	// Check the error code.
	if ret.ErrorCode != 1 {
		return nil, fmt.Errorf("%s: %s", ret.ErrorStatus, ret.Message)
	}

	return ret, nil
}

func (c *GetLinkedProfilesCall) doRequest() (*http.Response, error) {

	// Setup url.
	url := fmt.Sprintf("%sDestiny2/%s/Profile/%s/LinkedProfiles/?", c.s.basePath, c.membershipType, c.membershipId)

	// Atatch params params to the url.
	for k, v := range c.queryParams {
		url += k + "=" + v + "&"
	}

	// Create request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}
