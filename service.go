package bungo

import (
	"errors"
	"net/http"
)

type Service struct {
	client   *http.Client
	basePath string // API endpoint base URL.

	// App      *AppService
	// User     *UserService
	// Content  *ContentService
	// Forum    *ForumService
	// Groupv2  *Groupv2Service
	Destiny2 *destiny2Service
	// CommunityContent *CommunityContentService
	// Trending *TrendingService
	// Fireteam *FireteamService
}

func NewService(client *http.Client) (*Service, error) {

	if client == nil {
		return nil, errors.New("client is nil")
	}

	// Create the service.
	s := &Service{
		client:   client,
		basePath: basePath,
	}

	// Create the sub services.
	s.Destiny2 = NewDestiny2Service(s)

	return s, nil
}
