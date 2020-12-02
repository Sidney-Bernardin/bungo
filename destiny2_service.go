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

func (s *destiny2Service) GetDestinyManifest() *GetDestinyManifestCall {
	return &GetDestinyManifestCall{s: s.s}
}

type GetDestinyManifestCall struct {
	s *Service

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
		return nil, &bungoError{
			ErrorCode:       ret.ErrorCode,
			ThrottleSeconds: ret.ThrottleSeconds,
			ErrorStatus:     ret.ErrorStatus,
			Message:         ret.Message,
		}
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
		header:         http.Header{},
		entityType:     entityType,
		hashIdentifier: hashIdentifier,
	}
}

type GetDestinyEntityDefinitionCall struct {
	s *Service

	header http.Header

	entityType     string
	hashIdentifier string
}

func (c *GetDestinyEntityDefinitionCall) Header() http.Header {
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
		return nil, &bungoError{
			ErrorCode:       ret.ErrorCode,
			ThrottleSeconds: ret.ThrottleSeconds,
			ErrorStatus:     ret.ErrorStatus,
			Message:         ret.Message,
		}
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

	// If there is a preset api key, attach it to the request header.
	if c.s.apiKey != "" && c.s.apiKey != " " {
		c.header.Set("X-API-KEY", c.s.apiKey)
	}

	// Synchronize headers.
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}

func (s *destiny2Service) SearchDestinyPlayer(membershipType, displayName string) *SearchDestinyPlayerCall {
	return &SearchDestinyPlayerCall{
		s:              s.s,
		queryParams:    map[string]string{},
		header:         http.Header{},
		membershipType: membershipType,
		displayName:    displayName,
	}
}

type SearchDestinyPlayerCall struct {
	s *Service

	queryParams map[string]string
	header      http.Header

	membershipType string
	displayName    string
}

func (c *SearchDestinyPlayerCall) Header() http.Header {
	return c.header
}

func (c *SearchDestinyPlayerCall) ReturnOriginalProfile(value string) *SearchDestinyPlayerCall {
	c.queryParams["returnOriginalProfile"] = value
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
		return nil, &bungoError{
			ErrorCode:       ret.ErrorCode,
			ThrottleSeconds: ret.ThrottleSeconds,
			ErrorStatus:     ret.ErrorStatus,
			Message:         ret.Message,
		}
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

	// If there is a preset api key, attach it to the request header.
	if c.s.apiKey != "" && c.s.apiKey != " " {
		c.header.Set("X-API-KEY", c.s.apiKey)
	}

	// Synchronize headers.
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}

func (s *destiny2Service) GetLinkedProfiles(membershipType, membershipID string) *GetLinkedProfilesCall {
	return &GetLinkedProfilesCall{
		s:              s.s,
		queryParams:    map[string]string{},
		header:         http.Header{},
		membershipID:   membershipID,
		membershipType: membershipType,
	}
}

type GetLinkedProfilesCall struct {
	s *Service

	queryParams map[string]string
	header      http.Header

	membershipID   string
	membershipType string
}

func (c *GetLinkedProfilesCall) Header() http.Header {
	return c.header
}

func (c *GetLinkedProfilesCall) GetAllMemeberships(value string) *GetLinkedProfilesCall {
	c.queryParams["getAllMemberships"] = value
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
		return nil, &bungoError{
			ErrorCode:       ret.ErrorCode,
			ThrottleSeconds: ret.ThrottleSeconds,
			ErrorStatus:     ret.ErrorStatus,
			Message:         ret.Message,
		}
	}

	return ret, nil
}

func (c *GetLinkedProfilesCall) doRequest() (*http.Response, error) {

	// Setup url.
	url := fmt.Sprintf("%sDestiny2/%s/Profile/%s/LinkedProfiles?",
		c.s.basePath,
		c.membershipType,
		c.membershipID)

	// Atatch params params to the url.
	for k, v := range c.queryParams {
		url += k + "=" + v + "&"
	}

	// Create request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// If there is a preset api key, attach it to the request header.
	if c.s.apiKey != "" && c.s.apiKey != " " {
		c.header.Set("X-API-KEY", c.s.apiKey)
	}

	// Synchronize headers.
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}

func (s *destiny2Service) GetProfile(membershipType, membershipID string) *GetProfileCall {
	return &GetProfileCall{
		s:              s.s,
		queryParams:    map[string]string{},
		header:         http.Header{},
		membershipType: membershipType,
		membershipID:   membershipID,
	}
}

type GetProfileCall struct {
	s *Service

	queryParams map[string]string
	header      http.Header

	membershipType string
	membershipID   string
}

func (c *GetProfileCall) Header() http.Header {
	return c.header
}

func (c *GetProfileCall) Components(value string) *GetProfileCall {
	c.queryParams["components"] = value
	return c
}

func (c *GetProfileCall) Do() (*GetProifleResponse, error) {

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
	if ret.Profiles.ErrorCode.isError() {
		return nil, &bungoError{
			ErrorCode:       ret.Profiles.ErrorCode,
			ThrottleSeconds: ret.Profiles.ThrottleSeconds,
			ErrorStatus:     ret.Profiles.ErrorStatus,
			Message:         ret.Profiles.Message,
		}
	}

	return ret, nil
}

func (c *GetProfileCall) doRequest() (*http.Response, error) {

	// Setup url.
	url := fmt.Sprintf("%sDestiny2/%s/Profile/%s?",
		c.s.basePath,
		c.membershipType,
		c.membershipID)

	// Atatch params params to the url.
	for k, v := range c.queryParams {
		url += k + "=" + v + "&"
	}

	// Create request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// If there is a preset api key, attach it to the request header.
	if c.s.apiKey != "" && c.s.apiKey != " " {
		c.header.Set("X-API-KEY", c.s.apiKey)
	}

	// Synchronize headers.
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}

func (c *GetProfileCall) decodeResponse(body io.ReadCloser) (*GetProifleResponse, error) {

	var ret = &GetProifleResponse{}
	decoder := json.NewDecoder(body)

	// Decode the response into the right type based on the components parameter.
	switch c.queryParams["components"] {

	case "100", "Profiles":

		err := decoder.Decode(&ret.Profiles)
		return ret, err
	}

	return nil, fmt.Errorf("component %s is not yet supported", c.queryParams["components"])
}

func (s *destiny2Service) GetCharacter(membershipType, membershipID, characterID string) *GetCharacterCall {
	return &GetCharacterCall{
		s:              s.s,
		queryParams:    map[string]string{},
		header:         http.Header{},
		membershipType: membershipType,
		membershipID:   membershipID,
		characterID:    characterID,
	}
}

type GetCharacterCall struct {
	s *Service

	queryParams map[string]string
	header      http.Header

	membershipType string
	membershipID   string
	characterID    string
}

func (c *GetCharacterCall) Header() http.Header {
	return c.header
}

func (c *GetCharacterCall) Components(value string) *GetCharacterCall {
	c.queryParams["components"] = value
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

		// If the status code is 401 and we get a decode error,
		// it's likely do to a bad oauth token.
		if res.StatusCode == http.StatusUnauthorized {

			return nil, ErrUnauthorized
		}

		return nil, err
	}

	// Check for errors.
	if ret.CharacterEquipment.ErrorCode.isError() {
		return nil, &bungoError{
			ErrorCode:       ret.CharacterEquipment.ErrorCode,
			ThrottleSeconds: ret.CharacterEquipment.ThrottleSeconds,
			ErrorStatus:     ret.CharacterEquipment.ErrorStatus,
			Message:         ret.CharacterEquipment.Message,
		}
	}

	// Check for errors.
	if ret.Characters.ErrorCode.isError() {
		return nil, &bungoError{
			ErrorCode:       ret.Characters.ErrorCode,
			ThrottleSeconds: ret.Characters.ThrottleSeconds,
			ErrorStatus:     ret.Characters.ErrorStatus,
			Message:         ret.Characters.Message,
		}
	}

	// Check for errors.
	if ret.CharacterInventories.ErrorCode.isError() {
		return nil, &bungoError{
			ErrorCode:       ret.CharacterInventories.ErrorCode,
			ThrottleSeconds: ret.CharacterInventories.ThrottleSeconds,
			ErrorStatus:     ret.CharacterInventories.ErrorStatus,
			Message:         ret.CharacterInventories.Message,
		}
	}

	return ret, nil
}

func (c *GetCharacterCall) doRequest() (*http.Response, error) {

	// Setup url.
	url := fmt.Sprintf("%sDestiny2/%s/Profile/%s/Character/%s?",
		c.s.basePath,
		c.membershipType,
		c.membershipID,
		c.characterID)

	// Atatch params params to the url.
	for k, v := range c.queryParams {
		url += k + "=" + v + "&"
	}

	// Create request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// If there is a preset api key, attach it to the request header.
	if c.s.apiKey != "" && c.s.apiKey != " " {
		c.header.Set("X-API-KEY", c.s.apiKey)
	}

	// Synchronize headers.
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}

func (c *GetCharacterCall) decodeResponse(body io.ReadCloser) (*DestinyCharacterResponse, error) {

	var ret = &DestinyCharacterResponse{}
	decoder := json.NewDecoder(body)

	// Decode the response into the right type based on the components parameter.
	switch c.queryParams["components"] {

	case "200", "Characters":

		err := decoder.Decode(&ret.Characters)
		return ret, err

	case "201", "CharacterInventories":

		err := decoder.Decode(&ret.CharacterInventories)
		return ret, err

	case "205", "CharacterEquipment":

		err := decoder.Decode(&ret.CharacterEquipment)
		return ret, err
	}

	return nil, fmt.Errorf("component %s is not yet supported", c.queryParams["components"])
}

func (s *destiny2Service) EquipItem(body *ItemActionRequest) *EquipItemCall {
	return &EquipItemCall{
		s:           s.s,
		header:      http.Header{},
		requestBody: body,
	}
}

type EquipItemCall struct {
	s *Service

	header http.Header

	requestBody *ItemActionRequest
}

func (c *EquipItemCall) Header() http.Header {
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
		return nil, &bungoError{
			ErrorCode:       ret.ErrorCode,
			ThrottleSeconds: ret.ThrottleSeconds,
			ErrorStatus:     ret.ErrorStatus,
			Message:         ret.Message,
		}
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

	// If there is a preset api key, attach it to the request header.
	if c.s.apiKey != "" && c.s.apiKey != " " {
		c.header.Set("X-API-KEY", c.s.apiKey)
	}

	// Synchronize headers.
	req.Header = c.header

	// Execute the request.
	return c.s.client.Do(req)
}
