package bungo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type destiny2Service struct {
	s *Service
}

func NewDestiny2Service(s *Service) *destiny2Service {
	return &destiny2Service{s: s}
}

// ============================================================================
// Get Destiny Manifest
// ============================================================================

func (s *destiny2Service) GetDestinyManifest() *GetDestinyManifestCall {
	return &GetDestinyManifestCall{s: s.s}
}

type GetDestinyManifestCall struct {
	s      *Service
	header http.Header
}

func (c *GetDestinyManifestCall) Header() http.Header {

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

// ============================================================================
// Get Destiny Entity Definition
// ============================================================================

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

func (c *GetDestinyEntityDefinitionCall) Header() http.Header {

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

// ============================================================================
// Search Destiny Player
// ============================================================================

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

func (c *SearchDestinyPlayerCall) Header() http.Header {

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

// ============================================================================
// Get Linked Profiles
// ============================================================================

func (s *destiny2Service) GetLinkedProfiles(membershipType, membershipId string) *GetLinkedProfilesCall {
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

func (c *GetLinkedProfilesCall) Header() http.Header {

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
	url := fmt.Sprintf("%sDestiny2/%s/Profile/%s/LinkedProfiles?", c.s.basePath, c.membershipType, c.membershipId)

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

// ============================================================================
// Get Character
// ============================================================================

func (s *destiny2Service) GetCharacter(memType, memID, charID string) *GetCharacterCall {
	return &GetCharacterCall{
		s:           s.s,
		queryParams: map[string]string{},
		memType:     memType,
		memID:       memID,
		charID:      charID,
	}
}

type GetCharacterCall struct {
	s *Service

	queryParams map[string]string
	header      http.Header

	memType string
	memID   string
	charID  string
}

func (c *GetCharacterCall) Header() http.Header {

	if c.header == nil {
		c.header = make(http.Header)
	}

	return c.header
}

func (c *GetCharacterCall) Components(arg string) *GetCharacterCall {
	c.queryParams["components"] = arg
	return c
}

func (c *GetCharacterCall) Do() (*DestinyCharacterResponse, error) {

	// Make the request.
	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response.
	ret, err := c.decodeResponse(res.Body)
	if err != nil {
		return nil, err
	}

	// Check the error code.
	if ret.Characters.ErrorCode != 1 {
		if ret.CharacterInventories.ErrorCode != 1 {
			return nil, fmt.Errorf("%s: %s", ret.CharacterInventories.ErrorStatus, ret.CharacterInventories.Message)
		}
		return ret, nil
	}

	return ret, nil
}

func (c *GetCharacterCall) doRequest() (*http.Response, error) {

	// Setup url.
	url := fmt.Sprintf("%sDestiny2/%s/Profile/%s/Character/%s?",
		c.s.basePath,
		c.memType,
		c.memID,
		c.charID)

	// Create request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.header

	// Atatch params params to the url.
	for k, v := range c.queryParams {
		url += k + "=" + v + "&"
	}

	// Execute the request.
	return c.s.client.Do(req)
}

func (c *GetCharacterCall) decodeResponse(body io.ReadCloser) (*DestinyCharacterResponse, error) {

	ret := &DestinyCharacterResponse{}
	decoder := json.NewDecoder(body)

	switch c.queryParams["components"] {

	case "200":

		err := decoder.Decode(&ret.Characters)
		return ret, err

	case "205":

		err := decoder.Decode(&ret.CharacterInventories)
		return ret, err
	}

	return nil, fmt.Errorf("component %s is not yet supported", c.queryParams["components"])
}

// ============================================================================
// Equip Item
// ============================================================================

func (s *destiny2Service) EquipItem(body *ItemActionRequest) *EquipItemCall {
	return &EquipItemCall{s: s.s, requestBody: body}
}

type EquipItemCall struct {
	s           *Service
	requestBody *ItemActionRequest
	header      http.Header
}

func (c *EquipItemCall) Header() http.Header {

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
