package twitch

import (
	"fmt"
	"github.com/nicklaw5/helix"
)

const (
	// ExtAnalyticsScope extension analytics user permissions
	ExtAnalyticsScope string = "analytics:read:extensions"

	// GameAnalyticsScope game analytics user permissions
	GameAnalyticsScope string = "analytics:read:games"
)

// SetUserAccessToken set a new user access token for use
// in following API calls to the Twitch api
func (t *Twitch) SetUserAccessToken(code string) {
	t.helix.SetUserAccessToken(code)
}

// GetAuthURL returns url to begin oAUTH token exchange
func (t *Twitch) GetAuthURL(scopes []string) (url string) {
	t.helix.SetScopes(scopes)
	url = t.helix.GetAuthorizationURL("", false)

	return
}

// GetUserAccessToken post request to get user access token
// from returned oAUTH request exchange
func (t *Twitch) GetUserAccessToken(code string, scopes []string) (data helix.UserAccessCredentials, err error) {
	t.helix.SetScopes(scopes)

	resp, err := t.helix.GetUserAccessToken(code)
	if err != nil {
		return
	}

	if resp.Error != "" {
		err = fmt.Errorf(
			"err:%s msg:%s status:%d",
			resp.Error,
			resp.ErrorMessage,
			resp.StatusCode,
		)
		return
	}
	data = resp.Data

	return
}

// RenewUserAccessToken will refresh a user access token
// using a provided refresh code
func (t *Twitch) RenewUserAccessToken(refreshToken string) (data helix.UserAccessCredentials, err error) {
	resp, err := t.helix.RefreshUserAccessToken(refreshToken)
	if err != nil {
		return
	}

	if resp.Error != "" {
		err = fmt.Errorf(
			"err:%s msg:%s status:%d",
			resp.Error,
			resp.ErrorMessage,
			resp.StatusCode,
		)
		return
	}
	data = resp.Data

	return
}
