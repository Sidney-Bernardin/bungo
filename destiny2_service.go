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

type GetDestinyManifestResponse struct {
	Response        responses.DestinyManifest `json:"Response"`
	ErrorCode       int                       `json:"ErrorCode"`
	ThrottleSeconds int                       `json:"ThrottleSeconds"`
	ErrorStatus     string                    `json:"ErrorStatus"`
	Message         string                    `json:"Message"`
	MessageData     map[string]string         `json:"MessageData"`
}

type GetDestinyManifestCall struct {
	s *Service
}

func (c *GetDestinyManifestCall) Do() (*GetDestinyManifestResponse, error) {

	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret = &GetDestinyManifestResponse{}

	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	if ret.ErrorCode != 1 {
		return nil, fmt.Errorf("%s: %s", ret.ErrorStatus, ret.Message)
	}

	return ret, nil
}

func (c *GetDestinyManifestCall) doRequest() (*http.Response, error) {

	url := c.s.BasePath + "Destiny2/Manifest/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.s.client.Do(req)
}

func (s *Destiny2Service) GetDestinyEntityDefinition() *GetDestinyEntityDefinitionCall {
	return &GetDestinyEntityDefinitionCall{s: s.s, pathParams: map[string]string{}}
}

type GetDestinyEntityDefinitionResponse struct {
	Response        responses.DestinyDefinition `json:"Response"`
	ErrorCode       int                         `json:"ErrorCode"`
	ThrottleSeconds int                         `json:"ThrottleSeconds"`
	ErrorStatus     string                      `json:"ErrorStatus"`
	Message         string                      `json:"Message"`
	MessageData     map[string]string           `json:"MessageData"`
}

type GetDestinyEntityDefinitionCall struct {
	s          *Service
	pathParams map[string]string
	header     *http.Header
}

func (c *GetDestinyEntityDefinitionCall) EntityType(entityType string) *GetDestinyEntityDefinitionCall {
	c.pathParams["entityType"] = entityType
	return c
}

func (c *GetDestinyEntityDefinitionCall) HashIdentifier(hashIdentifier string) *GetDestinyEntityDefinitionCall {
	c.pathParams["hashIdentifier"] = hashIdentifier
	return c
}

func (c *GetDestinyEntityDefinitionCall) Do() (*GetDestinyEntityDefinitionResponse, error) {

	res, err := c.doRequest()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret = &GetDestinyEntityDefinitionResponse{}

	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	if ret.ErrorCode != 1 {
		return nil, fmt.Errorf("%s: %s", ret.ErrorStatus, ret.Message)
	}

	return ret, nil
}

func (c *GetDestinyEntityDefinitionCall) doRequest() (*http.Response, error) {

	url := fmt.Sprintf("%sDestiny2/Manifest/%s/%s",
		c.s.BasePath,
		c.pathParams["entityType"],
		c.pathParams["hashIdentifier"])

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", c.s.Config.ApiKey)

	return c.s.client.Do(req)
}
