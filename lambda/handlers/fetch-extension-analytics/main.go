package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackmcguire1/go-twitch-analytics/configuration"
	"github.com/jackmcguire1/go-twitch-analytics/internal/twitch"
	"github.com/jackmcguire1/go-twitch-analytics/internal/users"
	"github.com/jackmcguire1/go-twitch-analytics/internal/utils"
	"github.com/nicklaw5/helix"
	"github.com/satori/go.uuid"
)

var (
	twitchInternal *twitch.Twitch
	uService       *users.Users
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	twitchInternal, _ = twitch.New(env.ClientID, env.ClientSecret, "", env.TwitchAuthRedirect)
	uService = users.Client(env.UsersTable)
}

func handler(
	ctx context.Context,
	event events.CloudWatchEvent,
) (
	err error,
) {
	log.Println(utils.ToJSON(event))

	// renew user access token
	user, err := uService.Get(ctx, env.OwnerID)
	if err != nil {
		log.Println(err)
		return
	}

	credential, err := twitchInternal.RenewUserAccessToken(user.UserAccessToken.RefreshToken)
	if err != nil {
		log.Println(err)
		return
	}

	user.UserAccessToken = &credential

	err = uService.Put(ctx, user)
	if err != nil {
		log.Println(err)
		return
	}

	twitchInternal.SetUserAccessToken(credential.AccessToken)

	analytics := []helix.ExtensionAnalytic{}
	var bookmark string

	// Get all possible extension analytics retrievable
	// by the user
	for {
		responses, cursor, err := twitchInternal.GetExtensionAnalytics(
			"",
			nil,
			nil,
			10,
			bookmark,
		)
		if err != nil {
			break
		}
		log.Printf("found %d extension analytics", len(responses.ExtensionAnalytics))

		analytics = append(analytics, responses.ExtensionAnalytics...)

		if cursor == "" {
			break
		}
		bookmark = cursor
	}

	ch := make(chan error, len(analytics))

	for _, analytic := range analytics {

		log.Printf(
			"got new extension analytic file type:%q url:%q end-date:%q",
			analytic.Type,
			analytic.URL,
			analytic.DateRange.EndedAt.Format(time.RFC3339),
		)

		uuid := uuid.NewV4()
		key := fmt.Sprintf(
			"extension-analytics/%s/%s//%s/%s.csv",
			analytic.Type,
			analytic.ExtensionID,
			analytic.DateRange.EndedAt.Format("2006-01-02"),
			uuid.String(),
		)

		go func(key, url string) {
			ch <- utils.DownloadFileToS3(env.Bucket, key, url)
		}(key, analytic.URL)
	}

	counter := len(analytics)
	for resp := range ch {
		if resp != nil {
			log.Print(resp)
			err = resp
		}

		counter--
		if counter < 1 {
			break
		}
	}

	log.Println("done")
	return
}

func main() {
	lambda.Start(handler)
}
