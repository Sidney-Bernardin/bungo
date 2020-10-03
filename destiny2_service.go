package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
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
	Response        Manifest          `json:"Response"`
	ErrorCode       int               `json:"ErrorCode"`
	ThrottleSeconds int               `json:"ThrottleSeconds"`
	ErrorStatus     string            `json:"ErrorStatus"`
	Message         string            `json:"Message"`
	MessageData     map[string]string `json:"MessageData"`
}

type GetDestinyManifestCall struct {
	s      *Service
	header http.Header
}

func (c *GetDestinyManifestCall) Do() (*GetDestinyManifestResponse, error) {

	res, err := c.doRequest()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't do request")
	}
	defer res.Body.Close()

	var resBody = &GetDestinyManifestResponse{}

	if err := json.NewDecoder(res.Body).Decode(resBody); err != nil {
		return nil, errors.Wrap(err, "couldn't decode response")
	}

	if resBody.ErrorCode != 1 {
		return nil, fmt.Errorf("%s: %s", resBody.ErrorStatus, resBody.Message)
	}

	return resBody, nil
}

func (c *GetDestinyManifestCall) doRequest() (*http.Response, error) {

	url := c.s.BasePath + "Destiny2/Manifest/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't create request")
	}

	return c.s.client.Do(req)
}
