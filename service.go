package main

import (
	"errors"
	"net/http"
)

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL.

	// App      *AppService
	// User     *UserService
	// Content  *ContentService
	// Forum    *ForumService
	// Groupv2  *Groupv2Service
	Destiny2 *Destiny2Service
	// CommunityContent *CommunityContentService
	// Trending *TrendingService
	// Fireteam *FireteamService

}

func NewService(client *http.Client, config ServiceConfigInterface) (*Service, error) {

	if client == nil {
		return nil, errors.New("client is nil")
	}

	// Validate api key.
	if err := config.ValidateApiKey(); err != nil {
		return nil, errors.New("bad api key")
	}

	// Create the service.
	s := &Service{
		client:   client,
		BasePath: basePath,
	}

	// Create the sub services.
	s.Destiny2 = NewDestiny2Service(s)

	return s, nil
}
