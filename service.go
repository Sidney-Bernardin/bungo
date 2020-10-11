package bungo

import (
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

type Service struct {
	client   *http.Client
	basePath string // API endpoint base URL.
	config   *Config

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

type Config struct {
	ApiKey      string
	OAuthConfig *oauth2.Config
}

func NewService(client *http.Client, config *Config) (*Service, error) {

	if client == nil {
		return nil, errors.New("client is nil")
	}

	// Create the service.
	s := &Service{
		client:   client,
		basePath: basePath,
		config:   config,
	}

	// Create the sub services.
	s.Destiny2 = NewDestiny2Service(s)

	return s, nil
}
