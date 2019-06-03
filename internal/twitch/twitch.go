package twitch

import (
	"github.com/nicklaw5/helix"
)

// Twitch package struct to retain details
type Twitch struct {
	AppAccessToken  string
	UserAccessToken string
	ClientId        string
	ClientSecret    string
	helix           *helix.Client
}

// New Init Twitch API Client
func New(
	clientID string,
	clientSecret string,
	appAccessToken string,
	redirectURI string,
) (
	t *Twitch,
	err error,
) {

	client, err := helix.NewClient(&helix.Options{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		AppAccessToken: appAccessToken,
		RedirectURI:    redirectURI,
	})
	if err != nil {
		return
	}

	t = &Twitch{
		ClientId:       clientID,
		ClientSecret:   clientSecret,
		AppAccessToken: appAccessToken,
		helix:          client,
	}
	return
}
