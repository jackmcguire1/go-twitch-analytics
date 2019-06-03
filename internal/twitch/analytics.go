package twitch

import (
	"fmt"

	"github.com/nicklaw5/helix"
)

// GetExtensionAnalytics Requires user access token with
// 'analytics:read:extensions' permission
func (t *Twitch) GetExtensionAnalytics(
	extensionID string,
	startDate *helix.Time,
	endedDate *helix.Time,
	limit int,
	bookmark string,
) (
	data helix.ManyExtensionAnalytics,
	cursor string,
	err error,
) {
	payload := &helix.ExtensionAnalyticsParams{
		ExtensionID: extensionID,
		Type:        "overview_v2",
		First:       limit,
	}
	if bookmark != "" {
		payload.After = bookmark
	}

	if startDate != nil && endedDate != nil {
		payload.StartedAt = *startDate
		payload.EndedAt = *endedDate
	}

	resp, err := t.helix.GetExtensionAnalytics(payload)
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

	cursor = resp.Data.Pagination.Cursor
	data = resp.Data

	return
}

// GetGameAnalytics Requires user access token with
// 'analytics:read:games' permission
func (t *Twitch) GetGameAnalytics(
	gameID string,
	startDate *helix.Time,
	endedDate *helix.Time,
	limit int,
	bookmark string,
) (
	data helix.ManyGameAnalytics,
	cursor string,
	err error,
) {
	payload := &helix.GameAnalyticsParams{
		GameID: t.ClientId,
		First:  limit,
		Type:   "overview_v2",
	}

	if bookmark != "" {
		payload.After = bookmark
	}

	if startDate != nil && endedDate != nil {
		payload.StartedAt = *startDate
		payload.EndedAt = *endedDate
	}

	resp, err := t.helix.GetGameAnalytics(payload)
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

	cursor = resp.Data.Pagination.Cursor
	data = resp.Data

	return
}
