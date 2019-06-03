package env

import "os"

var (
	// Twitch extension settings

	//ClientID the extension or game clientID
	ClientID = os.Getenv("CLIENT_ID")

	//ClientSecret for extension, NOT EXTENSION SECRET
	ClientSecret = os.Getenv("CLIENT_SECRET")

	// OwnerID user's twitch account ID - NUMERIC
	OwnerID = os.Getenv("OWNER_ID")

	// TwitchAuthRedirect twitch auth redirect uri i.e. localhost:8080
	TwitchAuthRedirect = os.Getenv("AUTH_REDIRECT_URL")
)

var (
	// AWS Resources

	// Bucket the S3 bucket name
	Bucket = os.Getenv("BUCKET")

	// UsersTable the user dynamoDB table name
	UsersTable = os.Getenv("USERS_TBL")
)
